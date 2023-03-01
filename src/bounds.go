package main

import (
	"math"
	"math/rand"
	"time"
)

type Bounds struct {
	Shaper
	ID      int
	Minimum Tuple
	Maximum Tuple
}

type BoundingCube map[string]Tuple

func NewBounds() *Bounds {
	rand.Seed(time.Now().UnixNano())
	return &Bounds{
		ID:      rand.Intn(100000),
		Maximum: NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
		Minimum: NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)),
	}
}

func (b *Bounds) GetTransform() Matrix {
	return IdentityMatrix()
}

func (b *Bounds) AsCube() BoundingCube {
	return BoundingCube{
		"miXmiYmiZ": NewPoint(b.Minimum.X, b.Minimum.Y, b.Minimum.Z),
		"miXmiYmxZ": NewPoint(b.Minimum.X, b.Minimum.Y, b.Maximum.Z),
		"miXmxYmiZ": NewPoint(b.Minimum.X, b.Maximum.Y, b.Minimum.Z),
		"miXmxYmxZ": NewPoint(b.Minimum.X, b.Maximum.Y, b.Maximum.Z),
		"mxXmiYmiZ": NewPoint(b.Maximum.X, b.Minimum.Y, b.Minimum.Z),
		"mxXmiYmxZ": NewPoint(b.Maximum.X, b.Minimum.Y, b.Maximum.Z),
		"mxXmxYmiZ": NewPoint(b.Maximum.X, b.Maximum.Y, b.Minimum.Z),
		"mxXmxYmxZ": NewPoint(b.Maximum.X, b.Maximum.Y, b.Maximum.Z),
	}
}
func (s *Bounds) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}
func (s *Bounds) LocalIntersects(r Ray) map[int]Intersection {
	xtmin, xtmax := BoundCheckAxis(r.Origin.X, r.Direction.X, s.Minimum.X, s.Maximum.X)
	ytmin, ytmax := BoundCheckAxis(r.Origin.Y, r.Direction.Y, s.Minimum.Y, s.Maximum.Y)
	ztmin, ztmax := BoundCheckAxis(r.Origin.Z, r.Direction.Z, s.Minimum.Z, s.Maximum.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return map[int]Intersection{}
	}

	return map[int]Intersection{
		0: NewIntersection(tmin, s),
		1: NewIntersection(tmax, s),
	}
}
func BoundCheckAxis(origin, direction, min, max float64) (float64, float64) {
	var tmin, tmax float64
	tmin_numerator := (min - origin)
	tmax_numerator := max - origin
	if math.Abs(direction) >= epsilon {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = tmin_numerator * math.Inf(1)
		tmax = tmax_numerator * math.Inf(1)
	}
	if tmin > tmax {
		bob := tmin
		tmin = tmax
		tmax = bob
	}
	return tmin, tmax
}
