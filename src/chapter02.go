package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChapterTwo() {
	type Projectile struct {
		Position Tuple
		Velocity Tuple
	}

	type Environment struct {
		Gravity Tuple
		Wind    Tuple
	}

	tick := func(env Environment, proj Projectile) Projectile {
		position := proj.Position.Add(proj.Velocity)
		velocity := proj.Velocity.Add(env.Gravity).Add(env.Wind)
		return Projectile{position, velocity}
	}

	c := NewCanvas(900, 550)
	start := NewPoint(0, 1, 0)
	velocity := NewVector(1, 1.8, .0).Normalize().MultiplyScalar(11.25)
	p := Projectile{start, velocity}

	gravity := NewVector(0, -0.1, 0)
	wind := NewVector(-0.01, 0, 0)
	e := Environment{gravity, wind}

	pink := NewColor(1, 0.5, 0.5)
	x := 0
	for {
		x++
		p = tick(e, p)
		c.WritePixel(int(p.Position.X), c.Height-int(p.Position.Y), pink)
		if p.Position.Y <= 0 {
			break
		}
		fmt.Print(".")
	}
	tempFile := filepath.Join(os.TempDir(), "chapter02.ppm")
	os.WriteFile(tempFile, []byte(c.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
	fmt.Printf("%d iterations\n", x)
}
