package car

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/game/units"

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

	// Hardcoded braking force (in reverse direction)
	cBraking = 15000

	// https://en.wikipedia.org/wiki/Differential_(mechanical_device)
	cDifferentialRatio = 3.42

	// 2 pi radians per revolution
	cRpmConversionRate = 2 * math.Pi
)

const (
	rpmMin  = 1000
	rpmMax  = 6000
	rpmPeek = 4400

	torqueMin     = 290
	torqueMax     = 350
	torqueRedLine = 280
)

const (
	gearReverse gearInd = -1
	gearNeutral gearInd = 0
	gear1       gearInd = 1
	gear2       gearInd = 2
	gear3       gearInd = 3
	gear4       gearInd = 4
	gear5       gearInd = 5
	gear6       gearInd = 6
)

type (
	gearInd int8

	motor struct {
		// variables
		mass                   float64 // mass of car in KG
		transmissionEfficiency float64 // 0..1 (1 - 100%) - how many torque will stay on engine torque
		wheelsRadius           float64 // in meters (ex: 0.34 m)

		// state
		gearIndex        gearInd
		isBraking        bool
		throttlePosition float64 // 0..1 gas pedal push (0 to 100%), directed from user input
	}

	motorResult struct {
		acceleration        Vec
		angularAcceleration Angle

		infoForce    forceResult
		infoDrive    driveResult
		infoGasPedal float64
	}

	driveResult struct {
		driveForce Vec

		infoGear         int8
		infoGearRatio    float64
		infoRPM          float64
		infoMaxTorque    float64
		infoEngineTorque float64
	}

	forceResult struct {
		longitudinalForce Vec

		infoTraction          Vec
		infoDrag              Vec
		infoRollingResistance Vec
	}
)

func newMotor(
	mass float64,
	transmissionEfficiency float64,
	wheelsRadius float64,
) *motor {
	return &motor{
		mass:                   mass,
		transmissionEfficiency: transmissionEfficiency,
		wheelsRadius:           wheelsRadius,
		isBraking:              false,
		gearIndex:              gear1, // todo gearNeutral
	}
}

func (m *motor) Brake() {
	m.isBraking = true
}

// throttle 0 .. 1
func (m *motor) GasPedalPushPercent(throttlePosition float64) {
	m.throttlePosition = throttlePosition
}

func (m *motor) UpdateMotor(
	speed units.SpeedKmH,
	velocity Vec,
	direction Angle,
	steeringAngle Angle,
	angularVelocity Angle,
) motorResult {
	directionUnit := Vec{
		X: math.Cos(direction.Radians()),
		Y: math.Sin(direction.Radians()),
	}

	driveResult := m.calculateDriveForce(speed, directionUnit)

	force := calculateLongForce(velocity, driveResult.driveForce, directionUnit, m.isBraking)

	acceleration := force.longitudinalForce.Decrease(m.mass)

	angularAcceleration := calculateAngularAcceleration(
		velocity,
		force.infoTraction,
		steeringAngle,
		angularVelocity,
	)

	// return back braking model
	m.isBraking = false

	return motorResult{
		acceleration:        acceleration,
		angularAcceleration: angularAcceleration,

		infoForce:    force,
		infoDrive:    driveResult,
		infoGasPedal: m.throttlePosition,
	}
}

func (m *motor) calculateDriveForce(speed units.SpeedKmH, directionUnit Vec) driveResult {
	gearRatio := gearRation(m.gearIndex)
	rpm := engineRpm(m.gearIndex, m.wheelsRadius, speed)
	maxTorque := engineTorque(rpm)
	engineTorque := maxTorque * m.throttlePosition

	driveForce := directionUnit.
		Scale(engineTorque).
		Scale(gearRatio).
		Scale(cDifferentialRatio).
		Scale(m.transmissionEfficiency).
		Decrease(m.wheelsRadius)

	return driveResult{
		driveForce: driveForce,

		infoGear:         int8(m.gearIndex),
		infoGearRatio:    gearRatio,
		infoRPM:          rpm,
		infoMaxTorque:    maxTorque,
		infoEngineTorque: engineTorque,
	}
}

func calculateAngularAcceleration(velocity Vec, traction Vec, steeringAngle Angle, angularVelocity Angle) Angle {
	return engine.Angle0
	const frontAxleToCG = 1.4
	const rearAxleToCG = 1.4
	const inertia = 1500

	yawSpeedFront := frontAxleToCG * angularVelocity.Radians()
	yawSpeedRear := -rearAxleToCG * angularVelocity.Radians()

	var steeringVelocity float64
	if velocity.X < 0 {
		steeringVelocity = -1 * steeringAngle.Radians()
	} else {
		steeringVelocity = steeringAngle.Radians()
	}

	var slipAngleFront = math.Atan2(velocity.Y+yawSpeedFront, math.Abs(velocity.X)) - steeringVelocity
	var slipAngleRear = math.Atan2(velocity.Y+yawSpeedRear, math.Abs(velocity.X))

	const maxGrip = 2
	frictionForceYFront := math.Min(maxGrip, math.Max(-maxGrip, -5.0*slipAngleFront))
	frictionForceYRear := math.Min(maxGrip, math.Max(-maxGrip, -5.2*slipAngleRear))

	angularTorque := (frictionForceYFront + traction.Y) - frictionForceYRear
	return Angle(angularTorque / inertia)
}

func calculateLongForce(velocity Vec, driveForce Vec, directionUnit Vec, isBraking bool) forceResult {
	tractionForce := directionUnit.Mul(driveForce)

	speed := velocity.Magnitude()
	dragForce := Vec{
		X: -cDrag * velocity.X * speed,
		Y: -cDrag * velocity.Y * speed,
	}

	rollingResistanceForce := Vec{
		X: -cAirResistance * velocity.X,
		Y: -cAirResistance * velocity.Y,
	}

	if isBraking {
		isForwardDirection := directionUnit.Dot(velocity) > 0
		if isForwardDirection {
			tractionForce = directionUnit.Scale(-cBraking)
		}
	}

	longForce := engine.VectorSum(
		tractionForce,
		dragForce,
		rollingResistanceForce,
	)

	return forceResult{
		longitudinalForce: longForce,

		infoTraction:          tractionForce,
		infoDrag:              dragForce,
		infoRollingResistance: rollingResistanceForce,
	}
}

func engineRpm(gearIndex gearInd, wheelsRadius float64, speed units.SpeedKmH) float64 {
	// - 20 km/h = 20,000 m / 3600 s = 5.6 m/s.
	speedMetersPerSecond := speed * 1000 / 3600
	wheelRotationRate := speedMetersPerSecond / wheelsRadius
	gearRatio := gearRation(gearIndex)

	rpm := wheelRotationRate * gearRatio * cDifferentialRatio * 60 / cRpmConversionRate

	if rpm < rpmMin {
		rpm = rpmMin
	}

	if rpm > rpmMax {
		rpm = rpmMax
	}

	return rpm
}

func engineTorque(rpm float64) float64 {
	if rpm < rpmMin {
		return torqueMin
	}

	if rpm > rpmMax {
		return torqueRedLine
	}

	if rpm < rpmPeek {
		return engine.Lerp(torqueMin, torqueMax, (rpm-rpmMin)/(rpmPeek-rpmMin))
	}

	return engine.Lerp(torqueMax, torqueRedLine, (rpm-rpmPeek)/(rpmMax-rpmPeek))
}

func gearRation(index gearInd) float64 {
	switch index {
	case gearReverse: // reverse
		return 2.90
	case gearNeutral: // Differential
		return cDifferentialRatio
	case gear1: // g1
		return 2.66
	case gear2: // g2
		return 1.78
	case gear3: // g3
		return 1.30
	case gear4: // g4
		return 1.0
	case gear5: // g5
		return 0.74
	case gear6: // g6
		return 0.5
	default:
		panic(fmt.Sprintf("unknown gear index %d", index))
	}
}
