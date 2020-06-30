package gusic

import "math"

// Square is a square generator function
func Square(x float64) float64 {
	val := math.Sin(x)

	switch {
	case val == 0:
		return 0
	case val < 0:
		return -1
	case val > 0:
		return 1
	}

	return 1
}

// Sawtooth is a sawtooth generator function
func Sawtooth(x float64) float64 {
	return x - math.Floor(x)
}

// Triangle is a triangle generator function
func Triangle(x float64) float64 {
	return 2 * math.Abs(x/1-math.Floor(x/1+0.5))
}
