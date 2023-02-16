package main

import (
	"math"
	"math/rand"
	"time"
)

type Cone struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Minimum   float64
	Maximum   float64
	Closed    bool
}

func NewCone() *Cone {
	rand.Seed(time.Now().UnixNano())
	return &Cone{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
	}
}

func (s *Cone) GetType() string { return "cone" }
func (s *Cone) GetID() int      { return s.ID }
func (s *Cone) SetTransform(t Matrix) {
	s.Transform = s.Transform.MultiplyMatrix(t)
}
func (s *Cone) GetTransform() Matrix {
	return s.Transform
}
func (s *Cone) SetMaterial(m Material) {
	s.Material = m
}
func (s *Cone) GetMaterial() Material {
	return s.Material
}
func (s *Cone) Equals(s1 Shaper) bool {
	return s.ID == s1.GetID()
}

func (s *Cone) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Cone) LocalIntersects(r Ray) map[int]Intersection {
	xs := map[int]Intersection{}
	a := math.Pow(r.Direction.X, 2) - math.Pow(r.Direction.Y, 2) + math.Pow(r.Direction.Z, 2)
	b := 2*r.Origin.X*r.Direction.X -
		2*r.Origin.Y*r.Direction.Y +
		2*r.Origin.Z*r.Direction.Z
	c := math.Pow(r.Origin.X, 2) - math.Pow(r.Origin.Y, 2) + math.Pow(r.Origin.Z, 2)

	if a == 0 {
		if b != 0 {
			xs[len(xs)] = NewIntersection(-c/(2*b), s)
			//return xs
		}
	}
	xs = s.IntersectCaps(r, xs)

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
func (s *Cone) NormalAt(p Tuple) Tuple {
	return NormalAt(s, p)
}
func (s *Cone) LocalNormalAt(point Tuple) Tuple {
	dist := math.Pow(point.X, 2) + math.Pow(point.Z, 2)
	if dist < 1 && point.Y >= s.Maximum-epsilon {
		return NewVector(0, 1, 0)
	}
	if dist < 1 && point.Y <= s.Minimum+epsilon {
		return NewVector(0, -1, 0)
	}
	y := math.Sqrt(math.Pow(point.X, 2) + math.Pow(point.Z, 2))
	if point.Y > 0 {
		y *= -1
	}
	return NewVector(point.X, y, point.Z)
}

func (s *Cone) GetMinimum() float64 {
	return s.Minimum
}

func (s *Cone) GetMaximum() float64 {
	return s.Maximum
}

func (s *Cone) SetMinimum(m float64) {
	s.Minimum = m
}
func (s *Cone) SetMaximum(m float64) {
	s.Maximum = m
}

func (s *Cone) GetClosed() bool {
	return s.Closed
}

func (s *Cone) SetClosed(b bool) {
	s.Closed = b
}

func CheckCapCylinder(r Ray, t float64, y float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z
	return (math.Pow(x, 2) + math.Pow(z, 2)) <= math.Abs(y)
}

func (s *Cone) IntersectCaps(r Ray, xs map[int]Intersection) map[int]Intersection {

	if !s.Closed || epsilonEquals(math.Abs(r.Direction.Y), 0) {
		return xs
	}

	t := (s.Minimum - r.Origin.Y) / r.Direction.Y
	if CheckCapCylinder(r, t, s.Minimum) {
		xs[len(xs)] = NewIntersection(t, s)
	}
	t = (s.Maximum - r.Origin.Y) / r.Direction.Y
	if CheckCapCylinder(r, t, s.Maximum) {
		xs[len(xs)] = NewIntersection(t, s)
	}
	return xs
}
