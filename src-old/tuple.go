package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Tuple is a 4 element structure representing a Point or Vector in 3d space
type Tuple struct {
	X, Y, Z, W float64
}

func (v Tuple) isPoint() bool {
	return epsilonEquals(v.W, 1.0)
}

func (v Tuple) isVector() bool {
	return epsilonEquals(v.W, 0.0)
}

func (v Tuple) equals(t2 Tuple) bool {
	return epsilonEquals(v.X, t2.X) &&
		epsilonEquals(v.Y, t2.Y) &&
		epsilonEquals(v.Z, t2.Z) &&
		epsilonEquals(v.W, t2.W)
}

func (v Tuple) add(t2 Tuple) Tuple {
	return Tuple{v.X + t2.X, v.Y + t2.Y, v.Z + t2.Z, v.W + t2.W}
}

func (v Tuple) sub(t2 Tuple) Tuple {
	return Tuple{v.X - t2.X, v.Y - t2.Y, v.Z - t2.Z, v.W - t2.W}
}

func (v Tuple) neg() Tuple {
	return Tuple{v.X * -1, v.Y * -1, v.Z * -1, v.W * -1}
}

func (v Tuple) mult(m float64) Tuple {
	return Tuple{v.X * m, v.Y * m, v.Z * m, v.W * m}
}

func (v Tuple) matrixMult(m *mat.Dense) Tuple {
	x := Tuple{0, 0, 0, 0}
	for row := 0; row < 4; row++ {
		y := float64(m.At(row, 0))*v.X +
			float64(m.At(row, 1))*v.Y +
			float64(m.At(row, 2))*v.Z +
			float64(m.At(row, 3))*v.W
		switch row {
		case 0:
			x.X = y
		case 1:
			x.Y = y
		case 2:
			x.Z = y
		case 3:
			x.W = y
		}
	}
	return x
}

func (v Tuple) div(m float64) Tuple {
	return Tuple{v.X / m, v.Y / m, v.Z / m, v.W / m}
}

func (v Tuple) magnitude() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2) + math.Pow(v.W, 2))
}

func (v Tuple) normalize() Tuple {
	m := v.magnitude()
	return Tuple{
		v.X / m,
		v.Y / m,
		v.Z / m,
		v.W / m}
}

func (v Tuple) dot(b Tuple) float64 {
	return v.X*b.X +
		v.Y*b.Y +
		v.Z*b.Z +
		v.W*b.W
}

func (v Tuple) cross(b Tuple) Tuple {
	if v.W == 0 && b.W == 0 {
		return vector(v.Y*b.Z-v.Z*b.Y,
			v.Z*b.X-v.X*b.Z,
			v.X*b.Y-v.Y*b.X)
	}
	return vector(0, 0, 0)
}

func point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}
