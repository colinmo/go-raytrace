package main

type Shaper interface {
	Equals(t Shaper) bool
	Intersects(r Ray) map[int]Intersection
	GetID() int
	SetOrigin(t Tuple)
	GetOrigin() Tuple
	SetTransform(t Matrix)
	GetTransform() Matrix
	NormalAt(t Tuple) Tuple

	GetMaterial() Material
	SetMaterial(m Material)
}
