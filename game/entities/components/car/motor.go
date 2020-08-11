package car

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/game/units"

	"github.com/fe3dback/galaxy/engine"
)

const (
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
)

func engineTorque(gearInd gearInd, wheelsRadius float64, speed units.SpeedKmH, gasPedal float64) float64 {
	rpm := clcEngineRpm(gearInd, wheelsRadius, speed)
	maxTorque := clcEngineTorque(rpm)
	return maxTorque * gasPedal
}

func clcEngineRpm(gearIndex gearInd, wheelsRadius float64, speed units.SpeedKmH) float64 {
	// - 20 km/h = 20,000 m / 3600 s = 5.6 m/s.
	speedMetersPerSecond := speed * 1000 / 3600
	wheelRotationRate := speedMetersPerSecond / wheelsRadius
	gearRatio := clcGearRation(gearIndex)

	rpm := wheelRotationRate * gearRatio * cDifferentialRatio * 60 / cRpmConversionRate

	if rpm < rpmMin {
		rpm = rpmMin
	}

	if rpm > rpmMax {
		rpm = rpmMax
	}

	return rpm
}

func clcEngineTorque(rpm float64) float64 {
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

func clcGearRation(index gearInd) float64 {
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
