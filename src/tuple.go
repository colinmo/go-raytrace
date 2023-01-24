package main

type Tuple struct {
	X, Y, Z, W float64
}

func (v Tuple) isPoint() bool {
	return EpsilonEquals(v.W, 1.0)
}

func (v Tuple) isVector() bool {
	return EpsilonEquals(v.W, 0.0)
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

func (v Tuple) EqualsTuple(t Tuple) bool {
	return EpsilonEquals(v.X, t.X) &&
		EpsilonEquals(v.Y, t.Y) &&
		EpsilonEquals(v.Z, t.Z) &&
		EpsilonEquals(v.W, t.W)
}
