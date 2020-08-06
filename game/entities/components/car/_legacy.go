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

// The engine generates torque, which when applied to the wheels causes them
// to rotate.
//
// F traction = T wheel / R wheel
//
// Friction between the tires and the ground resists this motion,
// resulting in a force applied to the tires in the direction opposite to the
// rotation of the tires.
func (c Calculator) traction(wheelTorque float64, wheelRadius float64) float64 {
	return wheelTorque / wheelRadius
}

// When the car is in motion, an aerodynamic dragResistance will develop that will resist
// the motion of the car. Drag force is proportional to the square of the velocity
// and the formula to calculate the force is as follows.
//
// dragResistance = -cDrag * V * |V|
//
// Where V is the velocity vector and C-dragResistance is a constant, which is proportional to
// the frontal area of the car.
func (c Calculator) dragResistance(velocity Vec, cDrag Vec) Vec {
	abs := Vec{
		X: math.Abs(velocity.X),
		Y: math.Abs(velocity.Y),
	}

	return velocity.Mul(cDrag.Scale(-1)).Mul(abs)
}

// Rolling resistance is caused by friction between the rubber and road surfaces
// as the wheels roll along and is proportional to the velocity
//
// rollingResistance = -roadSurface * V
//
// roadSurface = 0 .. 1
func (c Calculator) rollingResistance(velocity Vec, roadSurface float64) Vec {
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

// The acceleration of the car is determined by the net force on the car and the
// car’s mass via Newton’s second law.
//
// a = F / M
//
func (c Calculator) acceleration(netForce Vec, mass float64) Vec {
	return netForce.Decrease(mass)
}

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

// ========================================================================
// C O M P U T E
// ========================================================================

type computeResults struct {
	// primary (only this should by used by physics):
	acceleration        Vec
	angularAcceleration float64

	// for debug/info purpose only
	infoWheelsSlipFront    float64
	infoWheelsSlipRear     float64
	infoForceLatFront      float64
	infoForceLatRear       float64
	infoEngineTurnoverRate float64
	infoEngineRPM          float64
	infoEngineTorque       float64
	infoWheelsTorque       float64
	infoTraction           float64
	infoRollingResistance  Vec
	infoResistance         Vec
	infoDragResistance     Vec
	infoForce              Vec
	infoBodyTorque         float64
}

func (c Calculator) compute(
	// from prev step:
	velocity Vec,
	angularSpeed float64,

	// input:
	steeringAngle engine.Angle,
	gearRatio float64,
	isBreaking bool,

	// env:
	roadSurface float64,

	// const:
	mass float64,
	distanceToFrontAxle float64,
	distanceToRearAxle float64,
	corneringStiffness float64,
	driveRatio float64,
	wheelsRadius float64,
	airResistance Vec,
) computeResults {

	// 2. Compute the slip angles for front and rear wheels (equation 5.2)
	wheelsSlipFront := c.wheelsSlipAngle(
		velocity,
		angularSpeed,
		steeringAngle,
		distanceToFrontAxle,
		true,
	)

	wheelsSlipRear := c.wheelsSlipAngle(
		velocity,
		angularSpeed,
		steeringAngle,
		distanceToRearAxle,
		false,
	)

	// 3. Compute Flat = Ca * slip angle (do for both rear and front wheels)
	forceLatFront := corneringStiffness * wheelsSlipFront
	forceLatRear := corneringStiffness * wheelsSlipRear

	// 4. Cap Flat to maximum normalized frictional force (do for both rear and
	// front wheels)
	const forceLat = 1.0 // todo?

	if forceLatFront > forceLat {
		forceLatFront = forceLat
	}
	if forceLatRear > forceLat {
		forceLatRear = forceLat
	}

	//5. Multiply Flat by the load (do for both rear and front wheels) to obtain
	//the cornering forces.
	const load = 0.01 // todo?

	forceLatFront *= load
	forceLatRear *= load

	//6. Compute the engine turn over rate Ωe = Vx 60*gk*G / (2π * rw)
	engineTurnoverRate := c.engineTurnoverRate(
		velocity.Y,
		gearRatio,
		driveRatio,
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

	engineRPM := c.engineAngularVelocity(
		engineTurnoverRate,
	)

	rpmMin := 500.0   // todo
	rpmPeek := 4200.0 // todo
	rpmMax := 6500.0  // todo

	//10. From 9, compute the maximum engine torque, Te
	var engineTorque float64

	if engineRPM <= rpmMin {
		engineTorque = rpmMin
	} else if engineRPM <= rpmPeek {
		// torque up min .. max
		engineTorque = engine.Lerp(rpmMin, rpmPeek, (engineRPM-rpmMin)/(rpmPeek-rpmMin))
	} else if engineRPM < rpmMax {
		// torque down max .. min, after some max torque
		engineTorque = engine.Lerp(rpmPeek, rpmMax, (engineRPM-rpmPeek)/(rpmMax-rpmPeek))
	}

	//11. Compute the maximum torque applied to the wheel Tw = Te * gk * G

	wheelsTorque := c.wheelsTorque(
		engineTorque,
		gearRatio,
		driveRatio,
	)

	//12. Multiply the maximum torque with the fraction of the throttle
	//position to get the actual torque applied to the wheel (F traction - The
	//traction force)

	traction := c.traction(
		wheelsTorque,
		wheelsRadius,
	)

	//13. If the player is braking replace the traction force from 12 to a defined
	//braking force

	if isBreaking {
		// todo (zero vector for zero traction)
		traction = 0.0
	}

	//14. If the car is in reverse gear replace the traction force from 12 to a
	//defined reverse force

	// todo

	//15. Compute rolling resistance Frr,
	// Frr,x = - Crr * Vx
	// Frr,z = - Crr * Vz
	rollingResistance := c.rollingResistance(velocity, roadSurface)

	//16. Compute drag resistance
	//  Fdrag, x = - Cdrag * Vx * |Vx|
	//  Fdrag, z = - Cdrag * Vz * |Vz|
	dragResistance := c.dragResistance(velocity, airResistance)

	//17. Compute total resistance (Fresistance) = rolling resistance + dragResistance resistance
	resistance := engine.VectorSum(rollingResistance, dragResistance)

	//18. Sum the force on the car body
	//F(x) =  F(traction) + F(lat, front) * sin (σ) * F(resistance, x)
	//F(z) = F(lat, rear) + F(lat, front) * cos (σ) * F(resistance, z)
	force := Vec{
		Y: traction + forceLatFront*math.Sin(steeringAngle.Radians())*resistance.Y,
		X: forceLatRear + forceLatFront*math.Cos(steeringAngle.Radians())*resistance.X,
	}

	//19. Compute the torque on the car body
	//Torque = cos (σ) * F(lat, front) * b – F(lat, rear) * c

	a := math.Cos(steeringAngle.Radians())
	bodyTorque := a*forceLatFront*distanceToFrontAxle - forceLatRear*distanceToRearAxle

	//20. Compute the acceleration
	//a = F / M

	acceleration := c.acceleration(force, mass)

	//21. Compute the angular acceleration
	//α = Torque/Inertia

	inertia := 10.0 // todo
	angularAcceleration := bodyTorque / inertia

	return computeResults{
		acceleration:        acceleration,
		angularAcceleration: angularAcceleration,

		// info data
		infoWheelsSlipFront:    wheelsSlipFront,
		infoWheelsSlipRear:     wheelsSlipRear,
		infoForceLatFront:      forceLatFront,
		infoForceLatRear:       forceLatRear,
		infoEngineTurnoverRate: engineTurnoverRate,
		infoEngineRPM:          engineRPM,
		infoEngineTorque:       engineTorque,
		infoWheelsTorque:       wheelsTorque,
		infoTraction:           traction,
		infoRollingResistance:  rollingResistance,
		infoDragResistance:     dragResistance,
		infoResistance:         resistance,
		infoForce:              force,
		infoBodyTorque:         bodyTorque,
	}
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
	velocity Vec,
	angularSpeed float64,
	steeringAngle engine.Angle,
	distanceToAxle float64,
	isFrontAxle bool,
) float64 {
	lateral := velocity.X + angularSpeed*distanceToAxle
	wheels := math.Atan(lateral / velocity.Y)

	if !isFrontAxle {
		return wheels
	}

	// todo:  sign function that extracts the sign of a real number.
	frontWheelsSteering := steeringAngle.Radians() * math.Abs(velocity.Y)
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
