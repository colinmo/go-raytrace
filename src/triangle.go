package main

import "math"

type Triangle struct {
	Shaper
	ID         int
	P1, P2, P3 Tuple
	E1, E2     Tuple
	Normal     Tuple
	Transform  Matrix
	Material   Material
	Parent     *Group
}

var triangleCount = 0

func NewTriangle(p1, p2, p3 Tuple) *Triangle {
	triangleCount++
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)
	return &Triangle{
		ID:        triangleCount,
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        e1,
		E2:        e2,
		Normal:    e2.CrossProduct(e1).Normalize(),
		Transform: BaseTransform,
		Material:  BaseMaterial,
		Parent:    nil,
	}
}

func (t *Triangle) GetPoint(name string) Tuple {
	switch name {
	case "P1":
		return t.P1
	case "P2":
		return t.P2
	case "P3":
		return t.P3
	default:
		return NewTuple(0, 0, 0, 0)
	}
}
func (t *Triangle) SetPoint(name string, p Tuple) {
	switch name {
	case "P1":
		t.P1 = p
	case "P2":
		t.P2 = p
	case "P3":
		t.P3 = p
	}
}
func (t *Triangle) GetVector(name string) Tuple {
	switch name {
	case "E1":
		return t.E1
	case "E2":
		return t.E2
	default:
		return NewTuple(0, 0, 0, 0)
	}
}
func (t *Triangle) SetVector(name string, v Tuple) {
	switch name {
	case "E1":
		t.E1 = v
	case "E2":
		t.E2 = v
	}
}

func (t *Triangle) GetNormal() Tuple {
	return t.Normal
}

func (t *Triangle) LocalNormalAt(p Tuple) Tuple {
	return t.Normal
}

func (t *Triangle) LocalIntersects(r Ray) map[int]Intersection {
	dirCrossE2 := r.Direction.CrossProduct(t.E2)
	det := t.E1.DotProduct(dirCrossE2)
	if math.Abs(det) < epsilon {
		return map[int]Intersection{}
	}

	f := 1.0 / det
	p1_to_origin := r.Origin.Subtract(t.P1)
	u := f * p1_to_origin.DotProduct(dirCrossE2)
	if u < 0 || u > 1 {
		return map[int]Intersection{}
	}

	originCrossE1 := p1_to_origin.CrossProduct(t.E1)
	v := f * r.Direction.DotProduct(originCrossE1)
	if v < 0 || (u+v) > 1 {
		return map[int]Intersection{}
	}
	t2 := f * t.E2.DotProduct(originCrossE1)
	return map[int]Intersection{0: NewIntersection(t2, t)}
}

func (t *Triangle) GetParent() *Group {
	return t.Parent
}

func (t *Triangle) SetParent(g *Group) {
	t.Parent = g
}

func (t *Triangle) GetID() int {
	return t.ID
}

func (s *Triangle) Bounds() *Bounds {
	b := NewBounds()
	b.Minimum = NewPoint(
		math.Min(s.P1.X, math.Min(s.P2.X, s.P3.X)),
		math.Min(s.P1.Y, math.Min(s.P2.Y, s.P3.Y)),
		math.Min(s.P1.Z, math.Min(s.P2.Z, s.P3.Z)),
	)

	b.Maximum = NewPoint(
		math.Max(s.P1.X, math.Max(s.P2.X, s.P3.X)),
		math.Max(s.P1.Y, math.Max(s.P2.Y, s.P3.Y)),
		math.Max(s.P1.Z, math.Max(s.P2.Z, s.P3.Z)),
	)

	return b
}

func (s *Triangle) GetTransform() Matrix {
	return IdentityMatrix()
}

func (s *Triangle) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Triangle) NormalAt(p Tuple) Tuple {
	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}
func (s *Triangle) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
}

func (s *Triangle) NormalToWorld(p Tuple) Tuple {
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

func (s *Triangle) GetMaterial() Material {
	return s.Material
}

func (s *Triangle) SetMaterial(m Material) {
	s.Material = m
}
