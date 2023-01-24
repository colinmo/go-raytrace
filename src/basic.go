package main

import "math"

const epsilon = 0.00001

func EpsilonEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
