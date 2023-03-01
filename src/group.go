package main

import (
	"fmt"
	"math"
	"sort"
)

type Group struct {
	Shaper
	ID        int
	Origin    Tuple
	Radius    float64
	Transform Matrix
	Material  Material
	Minimum   float64
	Maximum   float64
	Shapes    []Shaper
	Parent    *Group
}

var groupIdNumber = 0

func NewGroup() *Group {
	groupIdNumber++
	return &Group{
		ID:        groupIdNumber,
		Transform: IdentityMatrix(),
		Origin:    NewPoint(0, 0, 0),
		Material:  NewMaterial(),
		Shapes:    []Shaper{},
		Parent:    nil,
	}
}

func (s *Group) GetShapesCount() int {
	return len(s.Shapes)
}

func (s *Group) SetTransform(t Matrix) {
	s.Transform = s.Transform.MultiplyMatrix(t)
}
func (s *Group) GetTransform() Matrix {
	return s.Transform
}

func (g *Group) AddShape(s *Shaper) {
	g.Shapes = append(g.Shapes, *s)
	(*s).SetParent(g)
}

func (s *Group) AddGroup(t *Group) {
	s.Shapes = append(s.Shapes, t)
	(*t).SetParent(s)
}

func (s *Group) AddTriangle(t *Triangle) {
	s.Shapes = append(s.Shapes, t)
	(*t).SetParent(s)
}

func (s *Group) SetParent(g *Group) {
	s.Parent = g
}

func (s *Group) GetID() int {
	return s.ID
}

func (s *Group) GetShapes() []Shaper {
	return s.Shapes
}

func (s *Group) Equals(t Shaper) bool {
	return s.ID == t.GetID()
}

func (s *Group) Intersects(r Ray) map[int]Intersection {
	return Intersect(s, r)
}

func (s *Group) WorldToObject(p Tuple) Tuple {
	if s.Parent != nil {
		p = s.Parent.WorldToObject(p)
	}
	b := s.GetTransform()
	c := b.Inverse()
	return c.MultiplyTuple(p)
}
func (s *Group) NormalToWorld(p Tuple) Tuple {
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
func (s *Group) LocalIntersects(r Ray) map[int]Intersection {
	xs := make(map[int]Intersection)
	bG := s.Bounds()
	xx := bG.Intersects(r)
	if len(xx) == 0 {
		return xs
	}
	oks := 0
	mep := ""
	for _, o := range s.Shapes {
		xs2 := o.Intersects(r)
		mep = mep + fmt.Sprintf("Looking at shape %d\n", o.GetID())
		for _, b := range xs2 {
			xs[len(xs)] = b
			oks++
		}
	}

	var sortxs []Intersection
	for _, k := range xs {
		sortxs = append(sortxs, k)
	}
	sort.Slice(sortxs, func(i, j int) bool {
		return sortxs[i].T < sortxs[j].T
	})

	xs = make(map[int]Intersection)
	for i, k := range sortxs {
		xs[i] = k
	}
	return xs
}

func (s *Group) NormalAt(p Tuple) Tuple {
	localPoint := s.WorldToObject(p)
	localNormal := s.LocalNormalAt(localPoint)
	return s.NormalToWorld(localNormal)
}

func (s *Group) Bounds() *Bounds {
	b := NewBounds()
	// Check each containing object
	for _, o := range s.Shapes {
		oB := o.Bounds().AsCube()
		for _, c := range oB {
			// Convert to Object Space
			d := o.GetTransform()
			c = d.MultiplyTuple(c)
			b.Minimum.X = math.Min(c.X, b.Minimum.X)
			b.Minimum.Y = math.Min(c.Y, b.Minimum.Y)
			b.Minimum.Z = math.Min(c.Z, b.Minimum.Z)
			b.Maximum.X = math.Max(c.X, b.Maximum.X)
			b.Maximum.Y = math.Max(c.Y, b.Maximum.Y)
			b.Maximum.Z = math.Max(c.Z, b.Maximum.Z)
		}
	}

	return b
}
