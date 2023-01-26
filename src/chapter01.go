package main

import "fmt"

func ChapterOne() {
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

	p := Projectile{
		NewPoint(0, 1, 0),
		NewVector(1, 1, 0).Normalize(),
	}
	e := Environment{
		NewVector(0, -0.1, 0),
		NewVector(-0.01, 0, 0.01),
	}
	fmt.Printf("P: %f,%f,%f ; V: %f,%f,%f, FIRE!\n", p.Position.X, p.Position.Y, p.Position.Z, p.Velocity.X, p.Velocity.Y, p.Velocity.Z)
	for {
		p = tick(e, p)
		fmt.Printf("P: %f,%f,%f\n", p.Position.X, p.Position.Y, p.Position.Z)
		if p.Position.Y <= 0 {
			break
		}
	}
	fmt.Print("Done\n")
}
