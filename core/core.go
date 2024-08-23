package core

import "math"

const EPSILON = 0.00001

func FpEquals(a float64, b float64) bool {
	diff := math.Abs(a - b)
	return diff < EPSILON
}
