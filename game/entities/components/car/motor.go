package car

import (
	"fmt"
	"math"

	"github.com/fe3dback/galaxy/engine"
)

const turnoverRateMin = 5.0
const turnoverRateMax = 36000.0

const rpmMin = 250.0
const rpmPeek = 3600.0
const rpmMax = 4200.0

const powerMax = 100000.0
const powerMin = 0.0

const (
	motorGearReverse gearId = -1
	motorGearNeutral gearId = 0
	motorGearFirst   gearId = 1
	motorGearSecond  gearId = 2
	motorGearThird   gearId = 3
)

type (
	motor struct {
		engineOn bool
		power    float64
		gear     gearId
	}

	motorResult struct {
		torque float64

		infoEngineOn bool
		infoPower    float64
		infoGear     int8

		infoGearRatio    float64
		infoTurnoverRate float64
		infoRPM          float64
	}

	gearId int8
)

func newMotor() *motor {
	return &motor{ // todo not defaults
		engineOn: true,
		power:    0,
		gear:     motorGearFirst,
	}
}

func (m *motor) UpdateMotor() motorResult {
	gearRatio := m.gearRatio()
	turnoverRate := m.turnoverRate(gearRatio)
	rpm := m.angularVelocityRPM(turnoverRate)
	torque := m.torque(rpm)

	return motorResult{
		torque: torque,

		infoEngineOn:     m.engineOn,
		infoPower:        m.power,
		infoGear:         int8(m.gear),
		infoGearRatio:    gearRatio,
		infoTurnoverRate: turnoverRate,
		infoRPM:          rpm,
	}
}

func (m *motor) Start() {
	m.engineOn = true
}

func (m *motor) Stop() {
	m.engineOn = true
}

func (m *motor) SwitchGear(id gearId) {
	m.gear = id
}

func (m *motor) IncreasePower(forward float64) {
	m.power += forward
	m.power = engine.Clamp(m.power, powerMin, powerMax)
}

func (m *motor) torque(rpm float64) float64 {
	if rpm <= rpmMin {
		return rpmMin
	}

	if rpm <= rpmPeek {
		return engine.Lerp(rpmMin, rpmPeek, (rpm-rpmMin)/(rpmPeek-rpmMin))
	}

	if rpm < rpmMax {
		return engine.Lerp(rpmPeek, rpmMax, (rpm-rpmPeek)/(rpmMax-rpmPeek))
	}

	return rpmMax
}

func (m *motor) angularVelocityRPM(turnoverRate float64) float64 {
	return math.Pi * 2 * turnoverRate / 60
}

func (m *motor) turnoverRate(gearRatio float64) float64 {
	if !m.engineOn {
		return 0.0
	}

	rate := m.power * gearRatio
	rate = engine.Clamp(rate, turnoverRateMin, turnoverRateMax)

	return rate
}

func (m *motor) gearRatio() float64 {
	switch m.gear {
	case motorGearReverse:
		return -0.5
	case motorGearNeutral:
		return 0
	case motorGearFirst:
		return 0.5
	case motorGearSecond:
		return 1
	case motorGearThird:
		return 1.5
	default:
		panic(fmt.Sprintf("unknown gear %d", m.gear))
	}
}
