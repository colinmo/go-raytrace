package main

import "log"

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func NewRay(origin Tuple, direction Tuple) Ray {
	if origin.isVector() || direction.isPoint() {
		log.Fatal("Wrong Bits")
	}
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r *Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}
