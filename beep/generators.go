package beep

import "math"

// Square is a square generator function
func Square(period float64) Generator {
	return func(x float64) float64 {
		val := math.Sin(period * x)
		switch {
		case val == 0:
			return 0
		case val < 0:
			return -1
		case val > 0:
			return 1
		default:
			return 1
		}
	}
}

// Sawtooth is a sawtooth generator function
func Sawtooth(period float64) Generator {
	return func(x float64) float64 {
		return x/period - math.Floor(x/period)
	}
}

// Triangle is a triangle generator function
func Triangle(period float64) Generator {
	return func(x float64) float64 {
		return 2 * math.Abs(x/period-math.Floor(x/period+0.5))
	}
}
