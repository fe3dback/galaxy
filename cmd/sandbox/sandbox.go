package main

import (
	"fmt"
	"math"
)

func main() {
	for t := 0.0; t < 1; t += 0.01 {
		z := 1 - easeOutSine(t)
		fmt.Printf("%.2f \n", z)
	}

	for t := 0.0; t < 1; t += 0.2 {
		for y := 0.0; y < 1; y += 0.05 {
			z := easeOutSine(t)

			if fmt.Sprintf("%.1f", z) == fmt.Sprintf("%.1f", y) {
				fmt.Printf(".")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func easeOutSine(time float64) float64 {
	return 1 - math.Sin((time*math.Pi)/2)
}
