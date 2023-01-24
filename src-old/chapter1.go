package main

import "fmt"

type projectile struct {
	P, V Tuple
}
type environment struct {
	G, W Tuple
}

func tick(env environment, proj projectile) projectile {
	position := proj.P.add(proj.V)
	velocity := proj.V.add(env.G).add(env.W)

	return projectile{position, velocity}
}

func chapterOne() {
	p := projectile{point(0, 1, 0), vector(1, 1, 0).normalize()}
	e := environment{vector(0, -0.1, 0), vector(-0.01, 0, 0)}

	for p.P.Y > 0 {
		fmt.Printf("Y: %f\n", p.P.Y)
		p = tick(e, p)
	}
}
