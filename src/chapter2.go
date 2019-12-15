package main

import (
	"fmt"
	"io/ioutil"
)

type chapterTwotickprojectile struct {
	P, V Tuple
}
type chapterTwotickenvironment struct {
	G, W Tuple
}

func chapterTwotick(env chapterTwotickenvironment, proj chapterTwotickprojectile) chapterTwotickprojectile {
	position := proj.P.add(proj.V)
	velocity := proj.V.add(env.G).add(env.W)

	return chapterTwotickprojectile{position, velocity}
}

func chapterTwo() {
	start := point(0, 1, 0)
	velocity := vector(1, 1.8, 0).normalize().mult(11.25)
	p := chapterTwotickprojectile{start, velocity}

	gravity := vector(0, -0.1, 0)
	wind := vector(-0.01, 0, 0)
	e := chapterTwotickenvironment{gravity, wind}

	c := NewCanvas(900, 550)

	red := Color{1, 0, 0}

	for p.P.Y > 0 {
		c.writePixel(int(p.P.X), 550-int(p.P.Y), red)
		p = chapterTwotick(e, p)
	}
	fmt.Println("Writing to file...")
	ioutil.WriteFile("ChapterTwo.ppm", []byte(c.toPPM()), 0644)
}
