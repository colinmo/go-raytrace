package main

type Shaper interface {
	Equals(t Shaper) bool
	Intersects(r Ray) map[int]Intersection
	GetID() int
	SetTransform(t Matrix)
	GetTransform() Matrix
}
