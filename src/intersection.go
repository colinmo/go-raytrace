package main

import "log"

type Intersection struct {
	T      float64
	Object interface{}
}

func NewIntersection(t float64, shape interface{}) Intersection {
	ok := false
	switch shape.(type) {
	case Sphere:
		ok = true
	}
	if ok {
		return Intersection{T: t, Object: shape}
	}
	log.Fatalf("Wrong shape type")
	return Intersection{}
}

func (i *Intersection) ObjectEquals(target interface{}) bool {
	switch i.Object.(type) {
	case Sphere:
		s := i.Object.(Sphere)
		t := target.(Sphere)
		return s.Equals(t)
	}
	return false
}
