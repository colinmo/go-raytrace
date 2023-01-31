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
}

func NewSphere() *Sphere {
	rand.Seed(time.Now().UnixNano())
	return &Sphere{
		ID:        rand.Intn(100000),
		Transform: IdentityMatrix(),
		Origin:    NewPoint(0, 0, 0),
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
	return s.ID == t.GetID()
}

func (s *Sphere) SetTransform(t Matrix) {
	s.Transform = t
}

func (s *Sphere) GetTransform() Matrix {
	return s.Transform
}
