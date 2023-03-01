package main

import (
	"math"
	"math/rand"
	"time"
)

type Cylinder struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Minimum   float64
	Maximum   float64
	Closed    bool
	Parent    *Group
}

func NewCylinder() *Cylinder {
	rand.Seed(time.Now().UnixNano())
	return &Cylinder{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Radius:    1,
		Material:  BaseMaterial,
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
		Parent:    nil,
	}
}

func (s *Cylinder) GetType() string { return "cube" }
func (s *Cylinder) GetID() int      { return s.ID }
func (s *Cylinder) SetTransform(t Matrix) {
	s.Transform = s.Transform.MultiplyMatrix(t)
}
func (s *Cylinder) GetTransform() Matrix {
	return s.Transform
}
func (s *Cylinder) SetMaterial(m Material) {
	s.Material = m
}
func (s *Cylinder) GetMaterial() Material {
	return s.Material
}
func (s *Cylinder) Equals(s1 Shaper) bool {
	return s.ID == s1.GetID()
}

func (s *Cylinder) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Cylinder) LocalIntersects(r Ray) map[int]Intersection {
	a := math.Pow(r.Direction.X, 2) + math.Pow(r.Direction.Z, 2)
	xs := map[int]Intersection{}

	if math.Floor(a) == 0 { //} epsilon {
		xs = s.IntersectCaps(r, xs)
	}

	b := 2*r.Origin.X*r.Direction.X +
		2*r.Origin.Z*r.Direction.Z
	c := math.Pow(r.Origin.X, 2) + math.Pow(r.Origin.Z, 2) - 1

	disc := math.Pow(b, 2) - 4*a*c
	if disc < 0 {
		return map[int]Intersection{}
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		x := t0
		t0 = t1
		t1 = x
	}

	y0 := r.Origin.Y + t0*r.Direction.Y
	if s.Minimum < y0 && y0 < s.Maximum {
		xs[len(xs)] = NewIntersection(t0, s)
	}
	y1 := r.Origin.Y + t1*r.Direction.Y
	if s.Minimum < y1 && y1 < s.Maximum {
		xs[len(xs)] = NewIntersection(t1, s)
	}

	return xs
}

func (s *Cylinder) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
}
func (s *Cylinder) NormalToWorld(p Tuple) Tuple {
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
func (s *Cylinder) NormalAt(p Tuple) Tuple {

	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}
func (s *Cylinder) LocalNormalAt(point Tuple) Tuple {
	dist := math.Pow(point.X, 2) + math.Pow(point.Z, 2)
	if dist < 1 && point.Y >= s.Maximum-epsilon {
		return NewVector(0, 1, 0)
	}
	if dist < 1 && point.Y <= s.Minimum+epsilon {
		return NewVector(0, -1, 0)
	}
	return NewVector(point.X, 0, point.Z)
}

func (s *Cylinder) GetMinimum() float64 {
	return s.Minimum
}

func (s *Cylinder) GetMaximum() float64 {
	return s.Maximum
}

func (s *Cylinder) SetMinimum(m float64) {
	s.Minimum = m
}
func (s *Cylinder) SetMaximum(m float64) {
	s.Maximum = m
}

func (s *Cylinder) GetClosed() bool {
	return s.Closed
}

func (s *Cylinder) SetClosed(b bool) {
	s.Closed = b
}

func CheckCap(r Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z
	return (math.Pow(x, 2) + math.Pow(z, 2)) <= 1
}

func (s *Cylinder) IntersectCaps(r Ray, xs map[int]Intersection) map[int]Intersection {

	if !s.Closed || epsilonEquals(math.Abs(r.Direction.Y), 0) {
		return xs
	}

	t := (s.Minimum - r.Origin.Y) / r.Direction.Y
	if CheckCap(r, t) {
		xs[len(xs)] = NewIntersection(t, s)
	}
	t = (s.Maximum - r.Origin.Y) / r.Direction.Y
	if CheckCap(r, t) {
		xs[len(xs)] = NewIntersection(t, s)
	}
	return xs
}

func (s *Cylinder) Bounds() *Bounds {
	b := NewBounds()
	b.Minimum = NewPoint(-1, -1, s.GetMinimum())
	b.Maximum = NewPoint(1, 1, s.GetMaximum())

	return b
}

func (s *Cylinder) SetParent(g *Group) {
	s.Parent = g
}
