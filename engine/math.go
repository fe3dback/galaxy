package engine

import (
	"math"
	"math/rand"
)

const magic32 = 0x5F375A86

func FastInvSqrt64(n float32) float32 {
	// If n is negative return NaN
	if n < 0 {
		return float32(math.NaN())
	}
	// n2 and th are for one iteration of Newton's method later
	n2, th := n*0.5, float32(1.5)
	// Use math.Float32bits to represent the float32, n, as
	// an uint32 without modification.
	b := math.Float32bits(n)
	// Use the new uint32 view of the float32 to shift the bits
	// of the float32 1 to the right, chopping off 1 bit from
	// the fraction part of the float32.
	b = magic32 - (b >> 1)
	// Use math.Float32frombits to convert the uint32 bits back
	// into their float32 representation, again no actual change
	// in the bits, just a change in how we treat them in memory.
	// f is now our answer of 1 / sqrt(n)
	f := math.Float32frombits(b)
	// Perform one iteration of Newton's method on f to improve
	// accuracy
	f *= th - (n2 * f * f)

	// And return our fast inverse square root result
	return f
}

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
