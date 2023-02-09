package main

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
}

func NewSphere() *Sphere {
	rand.Seed(time.Now().UnixNano())
	return &Sphere{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
	}
}

func NewGlassSphere() *Sphere {
	rand.Seed(time.Now().UnixNano())
	me := Sphere{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
	}
	me.Material.Transparency = 1.0
	me.Material.RefractiveIndex = 1.5
	return &me
}

func (s *Sphere) GetSavedRay() Ray {
	return NewRay(NewTuple(0, 0, 0, 0), NewTuple(0, 0, 0, 1))
}
func (s *Sphere) GetID() int {
	return s.ID
}

func (s *Sphere) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Sphere) LocalIntersects(r Ray) map[int]Intersection {
	sphereToRay := r.Origin.Subtract(NewPoint(0, 0, 0))
	a := r.Direction.DotProduct(r.Direction)
	b := r.Direction.DotProduct(sphereToRay) * 2
	c := sphereToRay.DotProduct(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return Intersections()
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return map[int]Intersection{0: NewIntersection(t1, s), 1: NewIntersection(t2, s)}
}

func (s *Sphere) NormalAt(p Tuple) Tuple {
	return NormalAt(s, p)
}

func (s *Sphere) LocalNormalAt(p Tuple) Tuple {
	return p.Subtract(NewPoint(0, 0, 0))
}

func (s *Sphere) Equals(t Shaper) bool {
	return s.Transform.EqualsMatrix(t.GetTransform()) &&
		s.Origin.EqualsTuple(t.GetOrigin()) &&
		s.Material.Equals(t.GetMaterial())
}

func (s *Sphere) SetTransform(t Matrix) {
	s.Transform = t
}

func (s *Sphere) GetTransform() Matrix {
	return s.Transform
}

func (s *Sphere) SetOrigin(t Tuple) {
	s.Origin = t
}

func (s *Sphere) GetOrigin() Tuple {
	return s.Origin
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}
