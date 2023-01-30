package main

type Shaper interface {
	Equals(t interface{}) bool
	Intersects(r Ray) map[int]float64
}
