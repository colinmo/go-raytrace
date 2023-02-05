package main

import (
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
	return inters2
}

func (w *World) ShadeHit(comps Computations) Color {
	return Lighting(
		comps.Object.GetMaterial(),
		w.GetLight(),
		comps.Point,
		comps.Eyev,
		comps.Normalv)
}

func (w *World) ColorAt(r Ray) Color {
	i := w.Intersect(r)
	hit, is := Hit(i)
	if !hit {
		return NewColor(0, 0, 0)
	}
	comps := is.PrepareComputations(r)
	return w.ShadeHit(comps)
}