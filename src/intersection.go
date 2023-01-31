package main

import (
	"math"
)

type Intersection struct {
	T      float64
	Object Shaper
}

func NewIntersection(t float64, shape Shaper) Intersection {
	return Intersection{T: t, Object: shape}
}

func (i *Intersection) ObjectEquals(target Shaper) bool {
	return i.Object.GetID() == target.GetID()
}

func (i *Intersection) Equals(i2 Intersection) bool {
	return i.T == i2.T && i.Object.GetID() == i2.Object.GetID()
}

func Intersections(x ...Intersection) map[int]Intersection {
	z := map[int]Intersection{}
	for i, y := range x {
		z[i] = y
	}
	return z

}

func Hit(inters map[int]Intersection) (bool, Intersection) {
	var hit = Intersection{T: math.Inf(1)}
	var set = false
	for _, y := range inters {
		if y.T > 0 && y.T < hit.T {
			hit = y
			set = true
		}
	}
	return set, hit
}
