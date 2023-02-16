package main

import (
	"math"
	"sort"
)

type World struct {
	Lights  []Light
	Objects []Shaper
}

func NewWorld() World {
	return World{
		Lights:  []Light{},
		Objects: []Shaper{},
	}
}

func DefaultWorld() World {
	s1 := NewSphere()
	s1.Material.Color = NewColor(0.8, 1, 0.6)
	s1.Material.Diffuse = 0.7
	s1.Material.Specular = 0.2
	s2 := NewSphere()
	s2.Transform = NewScaling(0.5, 0.5, 0.5)
	return World{
		Lights:  []Light{NewLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))},
		Objects: []Shaper{s1, s2},
	}
}

func (w World) Contains(s Shaper) bool {
	for _, x := range w.Objects {
		if x.Equals(s) {
			return true
		}
	}
	return false
}

func (w World) GetLight() Light {
	return w.Lights[0]
}

func (w *World) SetLight(l Light) {
	w.Lights = make([]Light, 1)
	w.Lights[0] = l
}

func (w *World) Intersect(r Ray) map[int]Intersection {
	inters := []Intersection{}
	for _, o := range w.Objects {
		mep := o.Intersects(r)
		for _, j := range mep {
			inters = append(inters, j)
		}
	}

	sort.SliceStable(inters, func(i, j int) bool {
		return inters[i].T < inters[j].T
	})

	inters2 := map[int]Intersection{}
	for x, o := range inters {
		inters2[x] = o
	}
	/*
		if STOPHERE {
			log.Fatalf(
				"Objects: %d\n\nRay: %v\nO3: %s:%v\nO4: %s:%v\n\nC: %d: %v|%v",
				len(w.Objects),
				r,
				w.Objects[2].GetType(),
				w.Objects[2],
				w.Objects[3].GetType(),
				w.Objects[3],
				len(inters),
				inters,
				inters2)
		}
	*/
	return inters2
}

var STOPHERE = false

func (w *World) ShadeHit(comps Computations, remaining int) Color {
	inShadow := w.IsShadowed(comps.OverPoint)
	surface := Lighting(
		comps.Object.GetMaterial(),
		comps.Object,
		w.GetLight(),
		comps.OverPoint,
		comps.Eyev,
		comps.Normalv,
		inShadow)
	reflected := w.ReflectedColor(comps, remaining)
	refracted := w.RefractedColor(comps, remaining)

	material := comps.Object.GetMaterial()
	if material.Reflective > 0 && material.Transparency > 0 {
		reflectance := comps.Schlick()
		return surface.Add(reflected.MultiplyScalar(reflectance)).Add(refracted.MultiplyScalar(1 - reflectance))
	}
	return surface.Add(reflected).Add(refracted)
}

func (w *World) ColorAt(r Ray, remaining int) Color {
	i := w.Intersect(r)
	hit, is := Hit(i)

	if !hit {
		return NewColor(0, 0, 0)
	}
	comps := is.PrepareComputations(r, i)
	return w.ShadeHit(comps, remaining)
}

func (w *World) IsShadowed(p Tuple) bool {
	v := w.Lights[0].Position.Subtract(p)
	distance := v.Magnitude()
	direction := v.Normalize()

	r := NewRay(p, direction)
	intersections := w.Intersect(r)

	ishit, h := Hit(intersections)
	if ishit && h.T < distance {
		return true
	}
	return false
}

func (w *World) ReflectedColor(comps Computations, remaining int) Color {
	if remaining < 1 {
		return NewColor(0, 0, 0)
	}
	if comps.Object.GetMaterial().Reflective == 0 {
		return NewColor(0, 0, 0)
	}
	reflectRay := NewRay(comps.OverPoint, comps.Reflectv)
	color := w.ColorAt(reflectRay, remaining-1)
	return color.MultiplyScalar(comps.Object.GetMaterial().Reflective)
}

func (w *World) RefractedColor(comps Computations, remaining int) Color {
	if remaining < 1 {
		return NewColor(0, 0, 0)
	}
	if comps.Object.GetMaterial().Transparency == 0 {
		return NewColor(0, 0, 0)
	}
	nRatio := comps.N1 / comps.N2
	cosI := comps.Eyev.DotProduct(comps.Normalv)
	sin2T := nRatio * nRatio * (1 - cosI*cosI)
	if sin2T > 1 {
		return NewColor(0, 0, 0)
	}
	cosT := math.Sqrt(1.0 - sin2T)
	direction := comps.Normalv.MultiplyScalar(nRatio*cosI - cosT).Subtract(comps.Eyev.MultiplyScalar(nRatio))
	refractRay := NewRay(comps.UnderPoint, direction)
	/*
		if STOPHERE {
			c1 := w.ColorAt(refractRay, remaining-1)
			color := c1.MultiplyScalar(comps.Object.GetMaterial().Transparency)
			log.Fatalf(
				"REFRACT FAIL %v\n%v\n%v\nN %f, C %f, S %f, CT %f\nP: %v, D: %v",
				refractRay,
				c1,
				color,
				nRatio,
				cosI,
				sin2T,
				cosT,
				comps.UnderPoint,
				direction)
		}
	*/
	color := w.ColorAt(refractRay, remaining-1).MultiplyScalar(comps.Object.GetMaterial().Transparency)
	return color
}
