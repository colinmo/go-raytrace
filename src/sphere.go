package main

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
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
		Transform: IdentityMatrix(),
		Origin:    NewPoint(0, 0, 0),
		Material:  NewMaterial(),
	}
}

func (s *Sphere) GetID() int {
	return s.ID
}

func (s *Sphere) Intersects(r Ray) map[int]Intersection {
	r2 := r.Transform(s.Transform.Inverse())

	sphereToRay := r2.Origin.Subtract(NewPoint(0, 0, 0))
	a := r2.Direction.DotProduct(r2.Direction)
	b := r2.Direction.DotProduct(sphereToRay) * 2
	c := sphereToRay.DotProduct(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return Intersections()
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return map[int]Intersection{0: NewIntersection(t1, s), 1: NewIntersection(t2, s)}
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

func (s *Sphere) NormalAt(p Tuple) Tuple {
	v := s.Transform.Inverse()
	objectPoint := v.MultiplyTuple(p)
	objectNormal := objectPoint.Subtract(NewPoint(0, 0, 0))
	trans := v.Transpose()
	worldNormal := trans.MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m Material) {
	s.Material = m
}
