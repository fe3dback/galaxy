package engine

import "math"

func roundTo(n float64) float64 {
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

func Clamp(n, min, max float64) float64 {
	if n <= min {
		return min
	}

	if n >= max {
		return max
	}

	return n
}
