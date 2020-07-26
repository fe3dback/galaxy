package engine

import "math"

type Angle float64

func Anglef(f float64) Angle {
	return Angle(math.Mod(f, 360))
}
