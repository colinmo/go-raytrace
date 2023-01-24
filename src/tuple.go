package main

type Tuple struct {
	X, Y, Z, W float64
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
