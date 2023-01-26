package main

import (
	"math"
)

type Tuple struct {
	X, Y, Z, W float64
}

type Color struct {
	Red, Green, Blue float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

func NewColor(r, g, b float64) Color {
	return Color{
		Red:   r,
		Green: g,
		Blue:  b,
	}
}
func (v Tuple) isPoint() bool {
	return epsilonEquals(v.W, 1.0)
}

func (v Tuple) isVector() bool {
	return epsilonEquals(v.W, 0.0)
}

func (v Tuple) EqualsTuple(t Tuple) bool {
	return epsilonEquals(v.X, t.X) &&
		epsilonEquals(v.Y, t.Y) &&
		epsilonEquals(v.Z, t.Z) &&
		epsilonEquals(v.W, t.W)
}

func (v Tuple) Add(t Tuple) Tuple {
	return Tuple{
		v.X + t.X,
		v.Y + t.Y,
		v.Z + t.Z,
		v.W + t.W,
	}
}

func (v Tuple) Subtract(t Tuple) Tuple {
	return Tuple{
		v.X - t.X,
		v.Y - t.Y,
		v.Z - t.Z,
		v.W - t.W,
	}
}

func (v Tuple) Negative() Tuple {
	return Tuple{
		-1 * v.X,
		-1 * v.Y,
		-1 * v.Z,
		-1 * v.W,
	}
}

func (v Tuple) MultiplyScalar(scalar float64) Tuple {
	return Tuple{
		v.X * scalar,
		v.Y * scalar,
		v.Z * scalar,
		v.W * scalar,
	}
}

func (v Tuple) DivideScalar(scalar float64) Tuple {
	return Tuple{
		v.X / scalar,
		v.Y / scalar,
		v.Z / scalar,
		v.W / scalar,
	}
}

func (v Tuple) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

func (v Tuple) Normalize() Tuple {
	magnitude := v.Magnitude()
	return Tuple{
		v.X / magnitude,
		v.Y / magnitude,
		v.Z / magnitude,
		v.W / magnitude,
	}
}

func (v Tuple) DotProduct(b Tuple) float64 {
	return v.X*b.X +
		v.Y*b.Y +
		v.Z*b.Z +
		v.W*b.W
}

func (v Tuple) CrossProduct(b Tuple) Tuple {
	return NewVector(v.Y*b.Z-v.Z*b.Y,
		v.Z*b.X-v.X*b.Z,
		v.X*b.Y-v.Y*b.X)
}

func (v Tuple) ToMatrix() Matrix {
	M := NewMatrix(4, 1)
	M.Cells[0][0] = v.X
	M.Cells[1][0] = v.Y
	M.Cells[2][0] = v.Z
	M.Cells[3][0] = v.W
	return M
}

func (c1 Color) Equals(c2 Color) bool {
	return epsilonEquals(c1.Red, c2.Red) &&
		epsilonEquals(c1.Green, c2.Green) &&
		epsilonEquals(c1.Blue, c2.Blue)
}

func (c1 Color) Add(c2 Color) Color {
	return NewColor(
		c1.Red+c2.Red,
		c1.Green+c2.Green,
		c1.Blue+c2.Blue,
	)
}

func (c1 Color) Subtract(c2 Color) Color {
	return NewColor(
		c1.Red-c2.Red,
		c1.Green-c2.Green,
		c1.Blue-c2.Blue,
	)
}

func (c1 Color) MultiplyScalar(scalar float64) Color {
	return NewColor(
		c1.Red*scalar,
		c1.Green*scalar,
		c1.Blue*scalar,
	)
}

func (c1 Color) MultiplyColor(c2 Color) Color {
	return NewColor(
		c1.Red*c2.Red,
		c1.Green*c2.Green,
		c1.Blue*c2.Blue,
	)
}
