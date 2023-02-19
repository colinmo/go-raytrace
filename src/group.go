package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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
	Closed    bool
	Shapes    map[int]Shaper
}

func NewGroup() *Group {
	rand.Seed(time.Now().UnixNano())
	return &Group{
		ID:        rand.Intn(100000),
		Transform: IdentityMatrix(),
		Origin:    NewPoint(0, 0, 0),
		Material:  NewMaterial(),
		Shapes:    make(map[int]Shaper),
	}
}

func (s *Group) GetShapesCount() int {
	return len(s.Shapes)
}

func (s *Group) GetTransform() Matrix {
	return s.Transform
}

func (s *Group) AddShape(t *Shaper) {
	id := (*t).GetID()
	s.Shapes[id] = *t
	(*t).SetParent(s)
}

func (s *Group) GetID() int {
	return s.ID
}

func (s *Group) GetShapes() map[int]Shaper {
	return s.Shapes
}

func (s *Group) Equals(t Shaper) bool {
	return s.ID == t.GetID()
}

func (s *Group) LocalIntersects(r Ray) map[int]Intersection {
	oks := 0
	mep := ""
	xs := make(map[int]Intersection)
	for _, o := range s.Shapes {
		xs2 := o.Intersects(r)
		mep = mep + fmt.Sprintf("Looking at shape %d\n", o.GetID())
		for _, b := range xs2 {
			xs[len(xs)] = b
			oks++
		}
	}
	keys := make([]int, 0, len(xs))

	for k := range xs {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	xs2 := make(map[int]Intersection)
	for i, k := range keys {
		xs2[i] = xs[k]
	}
	return xs2
}
