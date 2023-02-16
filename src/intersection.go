package main

import (
	"log"
	"math"
	"math/rand"
)

type Intersection struct {
	ID     int
	T      float64
	Object Shaper
}

func NewIntersection(t float64, shape Shaper) Intersection {
	return Intersection{ID: rand.Intn(100000), T: t, Object: shape}
}

func (i *Intersection) ObjectEquals(target Shaper) bool {
	return i.Object.GetID() == target.GetID()
}

func (i *Intersection) Equals(i2 Intersection) bool {
	return i.ID == i2.ID
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
	T          float64
	Object     Shaper
	Point      Tuple
	Eyev       Tuple
	Normalv    Tuple
	OverPoint  Tuple
	Reflectv   Tuple
	Inside     bool
	N1         float64
	N2         float64
	UnderPoint Tuple
}

func (i *Intersection) PrepareComputations(r Ray, xs map[int]Intersection) Computations {
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
	comps.Reflectv = r.Direction.Reflect(comps.Normalv)
	if math.IsNaN(comps.OverPoint.X) {
		log.Fatalf("OP: %v|%v|%v|%f", comps.OverPoint, comps.Point, comps.Normalv, epsilon)
	}
	comps.OverPoint = comps.Point.Add(comps.Normalv.MultiplyScalar(epsilon))
	comps.UnderPoint = comps.Point.Subtract(comps.Normalv.MultiplyScalar(epsilon))

	// Refraction/ reflection
	containers := []Shaper{}

outer:
	for it := 0; it < len(xs); it++ {
		inter := xs[it]

		if i.Equals(inter) {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		appended := false
		for ij, j := range containers {
			if j.Equals(inter.Object) {
				containers = append(containers[:ij], containers[ij+1:]...)
				appended = true
				break
			}
		}
		if !appended {
			containers = append(containers, inter.Object)
		}

		if i.Equals(inter) {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex

			}

			break outer
		}

	}

	return comps
}

func (c *Computations) Schlick() float64 {
	cos := c.Eyev.DotProduct(c.Normalv)
	if c.N1 > c.N2 {
		n := c.N1 / c.N2
		sin2T := n * n * (1.0 - cos*cos)
		if sin2T > 1.0 {
			return 1.0
		}
		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}
	r0 := math.Pow((c.N1-c.N2)/(c.N1+c.N2), 2.0)
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
