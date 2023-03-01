package main

import (
	"math"
	"math/rand"
)

type Sphere struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Parent    *Group
}

var SphereCount = 0

func NewSphere() *Sphere {
	return &Sphere{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
		Parent:    nil,
	}
}

func NewGlassSphere() *Sphere {
	me := Sphere{
		ID:        rand.Intn(100000),
		Transform: BaseTransform,
		Origin:    BaseOrigin,
		Material:  BaseMaterial,
		Parent:    nil,
	}
	me.Material.Transparency = 1.0
	me.Material.RefractiveIndex = 1.5
	return &me
}
func (s *Sphere) GetType() string { return "sphere" }

func (s *Sphere) GetSavedRay() Ray {
	return NewRay(NewTuple(0, 0, 0, 0), NewTuple(0, 0, 0, 1))
}
func (s *Sphere) GetID() int {
	return s.ID
}
func (s *Sphere) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
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

	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}

func (s *Sphere) LocalNormalAt(p Tuple) Tuple {
	return p.Subtract(NewPoint(0, 0, 0))
}

func (s *Sphere) Equals(t Shaper) bool {
	equals := s.Transform.EqualsMatrix(t.GetTransform()) &&
		s.Origin.EqualsTuple(t.GetOrigin())
	if s.Material.HasPattern && t.GetMaterial().HasPattern {
		equals = equals && s.Material.Equals(t.GetMaterial())
	} else if s.Material.HasPattern || t.GetMaterial().HasPattern {
		return false
	}
	return equals
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

func (s *Sphere) GetParent() *Group {
	return s.Parent
}

func (s *Sphere) SetParent(g *Group) {
	s.Parent = g
}

func (s *Sphere) NormalToWorld(p Tuple) Tuple {
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

func (s *Sphere) Bounds() *Bounds {
	b := NewBounds()
	b.Minimum = NewPoint(-1, -1, -1)
	b.Maximum = NewPoint(1, 1, 1)
	return b
}
