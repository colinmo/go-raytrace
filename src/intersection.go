package main

import (
	"log"
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

type Computations struct {
	T         float64
	Object    Shaper
	Point     Tuple
	Eyev      Tuple
	Normalv   Tuple
	OverPoint Tuple
	Inside    bool
}

func (i *Intersection) PrepareComputations(r Ray) Computations {
	comps := Computations{
		T:      i.T,
		Object: i.Object,
	}
	comps.Point = r.Position(comps.T)
	comps.Eyev = r.Direction.Negative()
	comps.Normalv = comps.Object.NormalAt(comps.Point)
	if comps.Normalv.DotProduct(comps.Eyev) < 0 {
		comps.Inside = true
		comps.Normalv = comps.Normalv.Negative()
	} else {
		comps.Inside = false
	}
	if math.IsNaN(comps.OverPoint.X) {
		log.Fatalf("OP: %v|%v|%v|%f", comps.OverPoint, comps.Point, comps.Normalv, epsilon)
	}
	comps.OverPoint = comps.Point.Add(comps.Normalv.MultiplyScalar(epsilon))
	return comps
}
