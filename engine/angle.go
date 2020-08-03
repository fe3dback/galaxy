package engine

import "math"

const circleFull = 360.0
const circleHalf = 180.0
const convDeg2Rad = math.Pi / circleHalf
const convRad2Deg = circleHalf / math.Pi
const floatRoundPow = 10000

type Angle float64

func NewAngle(deg float64) Angle {
	return Angle(deg2rad(
		clampDeg(deg),
	))
}

func (a Angle) Degrees() float64 {
	return roundTo(clampDeg(
		rad2deg(float64(a)),
	))
}

func (a Angle) Radians() float64 {
	return float64(a)
}

func (a Angle) Add(t Angle) Angle {
	return a + t
}

func clampDeg(deg float64) float64 {
	return math.Mod(circleFull+math.Mod(deg, circleFull), circleFull)
}

func deg2rad(deg float64) float64 {
	return deg * convDeg2Rad
}

func rad2deg(rad float64) float64 {
	return rad * convRad2Deg
}
