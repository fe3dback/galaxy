package car

import (
	"math"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/units"
)

type Calculator struct {
}

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
	return velocity.Scale(roadSurface)
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
