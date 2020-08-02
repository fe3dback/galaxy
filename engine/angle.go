package engine

import "math"

type Angle float64

const max = 360.0

func Anglef(f float64) Angle {
	return Angle(math.Mod(max+math.Mod(f, max), max))
}

func (a Angle) ToFloat() float64 {
	return float64(a)
}

func (a Angle) Add(t Angle) Angle {
	return Anglef(a.ToFloat() + t.ToFloat())
}
