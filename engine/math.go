package engine

import (
	"math"
	"math/rand"
)

func RoundTo(n float64) float64 {
	if n == 0 {
		return 0
	}

	return math.Round(floatRoundPow*n) / floatRoundPow
}

func Lerp(a, b, t float64) float64 {
	if t <= 0 {
		return a
	}

	if t >= 1 {
		return b
	}

	return a + t*(b-a)
}

func LerpInverse(a, b, t float64) float64 {
	if a == t {
		return 0
	}

	if b == t {
		return 1
	}

	return (t - a) / (b - a)
}

// Lerpf will remap values v=0:originA->targetA, v=1:originA->targetB
func Lerpf(oA, oB, tA, tB, o float64) float64 {
	return Lerp(tA, tB, LerpInverse(oA, oB, o))
}

func Clamp(n, min, max float64) float64 {
	if n <= min {
		return min
	}

	if n >= max {
		return max
	}

	return n
}

func ClampInt(n, min, max int) int {
	if n <= min {
		return min
	}

	if n >= max {
		return max
	}

	return n
}

func RandomRange(from, to float64) float64 {
	if from > to {
		from, to = to, from
	}

	return from + rand.Float64()*(to-from)
}
