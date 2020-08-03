package engine

import "math"

func roundTo(n float64) float64 {
	if n == 0 {
		return 0
	}

	return math.Round(floatRoundPow*n) / floatRoundPow
}

func floatPrecision(num float64) float64 {
	// todo: remove
	output := math.Pow(10, float64(5))
	return floatRound(num*output) / output
}

func floatRound(x float64) float64 {
	// todo: remove
	if math.IsNaN(x) {
		return x
	}
	if x == 0.0 {
		return x
	}
	roundFn := math.Ceil
	if math.Signbit(x) {
		roundFn = math.Floor
	}
	xOrig := x
	x -= math.Copysign(0.5, x)
	if x == 0 || math.Signbit(x) != math.Signbit(xOrig) {
		return math.Copysign(0.0, xOrig)
	}
	if x == xOrig-math.Copysign(1.0, x) {
		return xOrig
	}
	r := roundFn(x)
	if r != x {
		return r
	}
	return roundFn(x*0.5) * 2.0
}
