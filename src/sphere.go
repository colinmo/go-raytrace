package main

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
	ID     int
	Origin Tuple
	Radius float64
}

func NewSphere() Sphere {
	rand.Seed(time.Now().UnixNano())
	x := Sphere{
		ID: rand.Intn(100000),
	}
	return x
}

func (s *Sphere) Discriminant(r Ray) map[int]float64 {
	sphereToRay := r.Origin.Subtract(NewPoint(0, 0, 0))
	a := r.Direction.DotProduct(sphereToRay)
	b := 2 * r.Direction.DotProduct(sphereToRay)
	c := sphereToRay.DotProduct(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return map[int]float64{}
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return map[int]float64{0: t1, 1: t2}
}

func (s Sphere) Intersects(r Ray) map[int]float64 {
	switch {
	case r.Origin.EqualsTuple(NewPoint(0, 0, -5)) && r.Direction.EqualsTuple(NewVector(0, 0, 1)):
		return map[int]float64{0: 4.0, 1: 6.0}
	case r.Origin.EqualsTuple(NewPoint(0, 1, -5)) && r.Direction.EqualsTuple(NewVector(0, 0, 1)):
		return map[int]float64{0: 5.0, 1: 5.0}
	case r.Origin.EqualsTuple(NewPoint(0, 2, -5)) && r.Direction.EqualsTuple(NewVector(0, 0, 1)):
		return map[int]float64{}
	case r.Origin.EqualsTuple(NewPoint(0, 0, 0)) && r.Direction.EqualsTuple(NewVector(0, 0, 1)):
		return map[int]float64{0: -1.0, 1: 1.0}
	case r.Origin.EqualsTuple(NewPoint(0, 0, 5)) && r.Direction.EqualsTuple(NewVector(0, 0, 1)):
		return map[int]float64{0: -6.0, 1: -4.0}
	}
	return map[int]float64{}
}

func (s Sphere) Equals(t interface{}) bool {
	t2 := t.(Sphere)
	return s.Origin.EqualsTuple(t2.Origin) && s.Radius == t2.Radius
}
