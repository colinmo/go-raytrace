package main

import (
	"math"
	"math/rand"
	"time"
)

type Plane struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Parent    *Group
}

func NewPlane() *Plane {
	rand.Seed(time.Now().UnixNano())
	return &Plane{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
		Parent:    nil,
	}
}
func (s *Plane) GetType() string { return "plane" }
func (s *Plane) GetID() int {
	return s.ID
}

func (s *Plane) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Plane) LocalIntersects(r Ray) map[int]Intersection {
	if math.Abs(r.Direction.Y) < epsilon {
		return map[int]Intersection{}
	}
	t := -r.Origin.Y / r.Direction.Y
	return map[int]Intersection{0: NewIntersection(t, s)}
}

func (s *Plane) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
}
func (s *Plane) NormalToWorld(p Tuple) Tuple {
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
func (s *Plane) NormalAt(p Tuple) Tuple {

	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}

func (s *Plane) LocalNormalAt(p Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (s *Plane) Equals(t Shaper) bool {
	return s.Transform.EqualsMatrix(t.GetTransform()) &&
		s.Origin.EqualsTuple(t.GetOrigin()) &&
		s.Material.Equals(t.GetMaterial())
}

func (s *Plane) SetTransform(t Matrix) {
	s.Transform = t
}

func (s *Plane) GetTransform() Matrix {
	return s.Transform
}

func (s *Plane) SetOrigin(t Tuple) {
	s.Origin = t
}

func (s *Plane) GetOrigin() Tuple {
	return s.Origin
}

func (s *Plane) GetMaterial() Material {
	return s.Material
}

func (s *Plane) SetMaterial(m Material) {
	s.Material = m
}
