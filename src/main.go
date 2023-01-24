package main

import (
	"fmt"
	"math"
)

const epsilon = 0.00001

func epsilonEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func main() {
	fmt.Printf("hi")
}
