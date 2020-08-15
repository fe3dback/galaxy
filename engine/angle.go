package engine

import "math"

const circleFullDeg = 360.0
const circleHalfDeg = 180.0
const circleFullRad = math.Pi * 2
const convDeg2Rad = math.Pi / circleHalfDeg
const convRad2Deg = circleHalfDeg / math.Pi
const floatRoundPow = 10000

type Angle float64

func NewAngle(deg float64) Angle {
	return Angle(deg2rad(
		clampDeg(deg),
	))
}

func (a Angle) Normalize() Angle {
	if a >= 0 {
		return a
	}

	return a + circleFullRad
}

func (a Angle) Flip() Angle {
	return circleFullRad - a
}

func (a Angle) Degrees() float64 {
	return RoundTo(clampDeg(
		rad2deg(float64(a)),
	))
}

func (a Angle) Radians() float64 {
	return float64(a)
}

func (a Angle) Unit() Vec {
	return Vec{
		X: math.Cos(a.Radians()),
		Y: -math.Sin(a.Radians()),
	}
}

func (a Angle) Add(t Angle) Angle {
	return a + t
}

func clampDeg(deg float64) float64 {
	return math.Mod(circleFullDeg+math.Mod(deg, circleFullDeg), circleFullDeg)
}

func deg2rad(deg float64) float64 {
	return deg * convDeg2Rad
}

func rad2deg(rad float64) float64 {
	return rad * convRad2Deg
}
