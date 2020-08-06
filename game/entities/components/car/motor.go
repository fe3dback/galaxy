package car

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/engine"
)

const (
	// Air density (rho) is 1.29 kg/m3 (0.0801 lb-mass/ft3),
	// frontal area is approx. 2.2 m2 (20 sq. feet),
	// Cd depends on the shape of the car and determined via wind tunnel tests.
	// Approximate value for a Corvette: 0.30. This gives us a value for Cdrag:
	// Cdrag = 0.5 * 0.30 * 2.2 * 1.29
	//       = 0.4257
	cDrag = 0.4257

	// We've already found that Crr should be approx. 30 times Cdrag.
	cAirResistance = cDrag * 30

	cBraking = 1000
)

type (
	motor struct {
		mass        float64
		engineForce float64
		isBraking   bool
	}

	motorResult struct {
		acceleration Vec

		infoForceLongitudinal Vec
		infoEngineForce       float64
	}
)

func newMotor(mass float64) *motor {
	return &motor{
		mass:        mass,
		engineForce: 0,
	}
}

func (m *motor) IncreaseForce(force float64) {
	m.engineForce += force
}

func (m *motor) Brake() {
	m.isBraking = true
}

func (m *motor) UpdateMotor(velocity Vec, direction Angle) motorResult {
	longitudinalForces := m.calculateLongForce(velocity, direction)

	acceleration := longitudinalForces.Decrease(m.mass)

	// return back braking model
	m.isBraking = false

	return motorResult{
		acceleration: acceleration,

		infoForceLongitudinal: longitudinalForces,
		infoEngineForce:       m.engineForce,
	}
}

func (m *motor) calculateLongForce(velocity Vec, direction Angle) engine.Vec {
	fmt.Println("--")
	directionUnit := Vec{
		X: math.Cos(direction.Radians()),
		Y: math.Sin(direction.Radians()),
	}
	fmt.Println(directionUnit)

	tractionForce := directionUnit.Scale(m.engineForce)
	fmt.Println(tractionForce)

	speed := velocity.Magnitude()
	fmt.Println(speed)
	dragForce := Vec{
		X: -cDrag * velocity.X * speed,
		Y: -cDrag * velocity.Y * speed,
	}
	fmt.Println(dragForce)

	rollingResistanceForce := Vec{
		X: -cAirResistance * velocity.X,
		Y: -cAirResistance * velocity.Y,
	}
	fmt.Println(rollingResistanceForce)

	if m.isBraking && speed > 0 {
		brakingForce := directionUnit.Scale(-cBraking)
		fmt.Println(brakingForce)

		return engine.VectorSum(
			brakingForce,
			dragForce,
			rollingResistanceForce,
		)
	}

	return engine.VectorSum(
		tractionForce,
		dragForce,
		rollingResistanceForce,
	)
}
