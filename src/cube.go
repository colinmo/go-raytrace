package main

import (
	"math"
	"math/rand"
	"time"
)

type Cube struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Parent    *Group
}

func NewCube() *Cube {
	rand.Seed(time.Now().UnixNano())
	return &Cube{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
		Parent:    nil,
	}
}
func (s *Cube) GetType() string { return "cube" }
func (s *Cube) GetID() int      { return s.ID }

func (s *Cube) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Cube) LocalIntersects(r Ray) map[int]Intersection {
	xtmin, xtmax := CheckAxis(r.Origin.X, r.Direction.X)
	ytmin, ytmax := CheckAxis(r.Origin.Y, r.Direction.Y)
	ztmin, ztmax := CheckAxis(r.Origin.Z, r.Direction.Z)

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

func CheckAxis(origin, direction float64) (float64, float64) {
	var tmin, tmax float64
	tmin_numerator := (-1 - origin)
	tmax_numerator := 1 - origin
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

func (s *Cube) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
}
func (s *Cube) NormalToWorld(p Tuple) Tuple {
	b := s.GetTransform()
	c := b.Inverse()
	d := c.Transpose()
	p = d.MultiplyTuple(p)
	p.W = 0
	p = p.Normalize()
	if s.Parent != nil {
		p = s.Parent.NormalToWorld(p)
	}
	return p
}
func (s *Cube) NormalAt(p Tuple) Tuple {
	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}

func (s *Cube) LocalNormalAt(p Tuple) Tuple {
	maxc := math.Max(math.Max(math.Abs(p.X), math.Abs(p.Y)), math.Abs(p.Z))
	if maxc == math.Abs(p.X) {
		return NewVector(p.X, 0, 0)
	}
	if maxc == math.Abs(p.Y) {
		return NewVector(0, p.Y, 0)
	}
	return NewVector(0, 0, p.Z)
}

func (s *Cube) Equals(t Shaper) bool {
	return s.ID == t.GetID()
}

func (s *Cube) SetTransform(t Matrix) {
	s.Transform = t
}

func (s *Cube) GetTransform() Matrix {
	return s.Transform
}

func (s *Cube) SetOrigin(t Tuple) {
	s.Origin = t
}

func (s *Cube) GetOrigin() Tuple {
	return s.Origin
}

func (s *Cube) GetMaterial() Material {
	return s.Material
}

func (s *Cube) SetMaterial(m Material) {
	s.Material = m
}

func (s *Cube) Bounds() *Bounds {
	b := NewBounds()
	b.Minimum = NewPoint(-1, -1, -1)
	b.Maximum = NewPoint(1, 1, 1)
	return b
}
