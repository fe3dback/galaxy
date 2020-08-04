package car

import (
	"math"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/units"
)

type Calculator struct {
}

// https://nccastaff.bournemouth.ac.uk/jmacey/MastersProjects/MSc12/Srisuchat/Thesis.pdf

// ========================================================================
// C A R  F O R C E S
// ========================================================================

// When the car is in motion, an aerodynamic drag will develop that will resist
// the motion of the car. Drag force is proportional to the square of the velocity
// and the formula to calculate the force is as follows.
//
// drag = -cDrag * V * |V|
//
// Where V is the velocity vector and C-drag is a constant, which is proportional to
// the frontal area of the car.
func (c Calculator) drag(velocity Vec, cDrag float64) Vec {
	abs := Vec{
		X: math.Abs(velocity.X),
		Y: math.Abs(velocity.Y),
	}

	return velocity.Scale(-cDrag).Mul(abs)
}

// Rolling resistance is caused by friction between the rubber and road surfaces
// as the wheels roll along and is proportional to the velocity
//
// rollingResistance = -roadSurface * V
//
// roadSurface = 0 .. 1
func (c Calculator) rollingResistance(roadSurface float64, velocity Vec) Vec {
	return velocity.Scale(-roadSurface)
}

// Gravity pulls the car towards the earth. The parallel component of
// gravitational force,
//
// Fg = m * g * sin(angle)
//
// can pull the car either forwards or backwards depending on whether
// the car is pointing uphill or downhill.
func (c Calculator) gravityForce(mass float64, slopeAngle float64) Vec {
	return engine.VectorForward(
		mass * units.Gravity * math.Sin(slopeAngle),
	)
}

// The engine generates torque, which when applied to the wheels causes them
// to rotate.
//
// F traction = T wheel / R wheel
//
// Friction between the tires and the ground resists this motion,
// resulting in a force applied to the tires in the direction opposite to the
// rotation of the tires.
func (c Calculator) traction(wheelTorque float64, wheelRadius float64) Vec {
	return engine.VectorForward(
		wheelTorque / wheelRadius,
	)
}

// To determine the position of the car first we must find the net force on the
// car. The total longitudinal force is the vector sum of these four forces.
//
// F long = F traction + F drag + F rr + Fg
//
// When braking, a braking force replaces the traction force, which is oriented
// in the opposite direction. A simple model of braking is as follows.
//
// F braking = - u * C braking
//
// Where u is a unit vector in the direction of the car’s heading and C braking is a
// constant.
func (c Calculator) netForce(
	traction Vec,
	drag Vec,
	rollingResistance Vec,
	gravityForce Vec,
	direction Vec,
	isBraking bool,
	breakingFactor float64,
) Vec {
	var wheelForce Vec

	if isBraking {
		wheelForce = direction.Scale(-breakingFactor)
	} else {
		wheelForce = traction
	}

	return engine.VectorSum(wheelForce, drag, rollingResistance, gravityForce)
}

// The acceleration of the car is determined by the net force on the car and the
// car’s mass via Newton’s second law.
//
// a = F / M
//
func (c Calculator) acceleration(netForce Vec, mass float64) Vec {
	return netForce.Decrease(mass)
}

// The car’s velocity is determined by integrating the acceleration over time
// using the numerical method for numerical integration
//
// V new = V + dt * a
//
// Where dt is the time increment in seconds between subsequent calls on the
// physics engine.
func (c Calculator) velocity(velocity Vec, acceleration Vec, dt float64) Vec {
	return velocity.Add(
		acceleration.Scale(dt),
	)
}

// The car’s position is in turn determined by integrating the velocity over time.
//
// P new = P + dt * v
//
// With these forces, we can simulate car acceleration fairly accurately.
// Together they also determined the top speed of the car since as the velocity of
// the car increases the resistance forces also increases. At some point the
// resistance forces and the engine force cancel each other out and the car has
// reached its top speed.
func (c Calculator) nextPosition(position Vec, velocity Vec, dt float64) Vec {
	return position.Add(
		velocity.Scale(dt),
	)
}

// ========================================================================
// W H E E L S
// ========================================================================

// The angular velocity of the engine in rad/s is obtained by multiplying
// the engine turnover rate by 2π and dividing by 60.
//
// ωe = 2π Ωe / 60
//
// When the engine runs, it generates a certain amount of torque. The torque
// that an engine delivers depend on the speed at which the engine is turning,
// usually expressed in terms of revolutions per minute, or rpm. The
// relationship torque versus rpm is usually provided as a curve known as a
// torque curve.
//
// RPM
func (c Calculator) engineAngularVelocity(turnoverRate float64) float64 {
	return math.Pi * 2 * turnoverRate / 60
}

// The relationship between the engine turnover rate and the wheel angular
// velocity is as follows.
// ωw = 2π Ωe / (60*gk*G)
//
func (c Calculator) wheelsAngularVelocity(turnoverRate float64, gearRatio float64, driveRatio float64) float64 {
	return math.Pi * 2 * turnoverRate / (60 * gearRatio * driveRatio)
}

// The torque applied to the wheels is not the same as the engine torque
// because the engine torque passes through a transmission before it is applied
// to the wheels. The gear ratio between two gears is the ratio of the gear
// diameters. Car transmission will typically have between three and six
// forward gears and one reverse gear. There is also an additional set of gears
// Development of a car physics engine for games 14
// between the transmission and the wheels known as the differential and the
// gear ratio of this gearset is called the final drive ratio. The wheel toque can
// be obtained using the following equation
//
// Tw = Te * gk * G
//
// Where Te is the engine torque, gk is the gear ratio of whatever gear the car is
// in and G is the final drive ratio.
func (c Calculator) wheelsTorque(engineTorque float64, gearRatio float64, driveRatio float64) float64 {
	return engineTorque * gearRatio * driveRatio
}

func (c Calculator) compute(
	lateralVelocity float64,
	forwardVelocity float64,
	angularSpeed float64,
	steeringAngle engine.Angle,
	distanceToAxle float64,
	corneringStiffness float64,
	gearRatio float64,
	driveRation float64,
) {
	// 2. Compute the slip angles for front and rear wheels (equation 5.2)
	wheelsSlipFront := c.wheelsSlipAngle(
		lateralVelocity,
		forwardVelocity,
		angularSpeed,
		steeringAngle,
		distanceToAxle,
		true,
	)

	wheelsSlipRear := c.wheelsSlipAngle(
		lateralVelocity,
		forwardVelocity,
		angularSpeed,
		steeringAngle,
		distanceToAxle,
		false,
	)

	// 3. Compute Flat = Ca * slip angle (do for both rear and front wheels)
	flatFront := corneringStiffness * wheelsSlipFront
	flatRear := corneringStiffness * wheelsSlipRear

	// 4. Cap Flat to maximum normalized frictional force (do for both rear and
	// front wheels)
	const flat = 1.0 // todo?

	if flatFront > flat {
		flatFront = flat
	}
	if flatRear > flat {
		flatRear = flat
	}

	//5. Multiply Flat by the load (do for both rear and front wheels) to obtain
	//the cornering forces.
	const load = 1.0 // todo?

	flatFront *= load
	flatRear *= load

	//6. Compute the engine turn over rate Ωe = Vx 60*gk*G / (2π * rw)
	engineTurnoverRate := c.engineTurnoverRate(
		forwardVelocity,
		gearRatio,
		driveRation,
		1.0, // todo?
	)

	//7. Clamp the engine turn over rate from 6 to the defined redline
	if engineTurnoverRate < 6 {
		engineTurnoverRate = 6 //todo
	}

	if engineTurnoverRate > 4000 {
		engineTurnoverRate = 4000 //todo
	}

	//8. If use automatic transmission call automaticTransmission() function
	//to shift the gear

	// todo

	//9. Compute the constant that define the torque curve line from the
	//engine turn over rate

	//10. From 9, compute the maximum engine torque, Te

	//11. Compute the maximum torque applied to the wheel Tw = Te * gk * G

	//12. Multiply the maximum torque with the fraction of the throttle
	//position to get the actual torque applied to the wheel (Ftraction - The
	//traction force)
	//Development of a car physics engine for games 20

	//13. If the player is braking replace the traction force from 12 to a defined
	//braking force

	//14. If the car is in reverse gear replace the traction force from 12 to a
	//defined reverse force

	//15. Compute rolling resistance Frr, x = - Crr * Vx and Frr,z = - Crr * Vz

	//16. Compute drag resistance Fdrag, x = - Cdrag * Vx * |Vx| and Fdrag, z = -
	//	Cdrag * Vz * |Vz|

	//17. Compute total resistance (Fresistance) = rolling resistance + drag
	//resistance

	//18. Sum the force on the car body
	//Fx = Ftraction + Flat, front * sin (σ) * Fresistance, x
	//Fz = Flat, rear + Flat, front * cos (σ) * Fresistance, z

	//19. Compute the torque on the car body
	//Torque = cos (σ) * Flat, front * b – Flat, rear * c

	//20. Compute the acceleration
	//a = F / M

	//21. Compute the angular acceleration
	//α = Torque/Inertia

	//22. Transform the acceleration from car reference frame to world
	//reference frame

	//23. Integrate the acceleration to get the velocity (in world reference
	//frame)
	//Vwc += dt * a

	//24. Integrate the velocity to get the new position in world coordinate
	//Pwc += dt * Vwc
}

// The slip angles for the front wheels and rear wheels are given by the
// following equations.
//
// front = arctan((Vlat + ω * b) / Vlong)) – σ * sgn(Vlong)
// rear  = arctan((Vlat - ω * c) / Vlong))
//
// Where ω is the angular speed of the car, σ is the steering angle, Vlat is the
// velocity in the lateral direction, Vlong is the velocity in the longitudinal
// direction, b is the distance from CG to the front axle, c is the distance from
// CG to the rear axle, and c is the sign function that extracts the sign of a real
// number.
func (c Calculator) wheelsSlipAngle(
	lateralVelocity float64,
	forwardVelocity float64,
	angularSpeed float64,
	steeringAngle engine.Angle,
	distanceToAxle float64,
	isFrontAxle bool,
) float64 {
	lateral := lateralVelocity + angularSpeed*distanceToAxle
	wheels := math.Atan(lateral / forwardVelocity)

	if !isFrontAxle {
		return wheels
	}

	// todo:  sign function that extracts the sign of a real number.
	frontWheelsSteering := steeringAngle.Radians() * math.Abs(forwardVelocity)
	return wheels - frontWheelsSteering
}

// Ωe = Vx 60*gk*G / (2π * rw)
func (c Calculator) engineTurnoverRate(
	forwardVelocity float64,
	gearRatio float64,
	driveRatio float64,
	rw float64, // todo
) float64 {
	return forwardVelocity * 60 * gearRatio * driveRatio / (math.Pi * rw)
}
