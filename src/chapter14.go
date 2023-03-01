package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func ChapterFourteen() {

	hexagonCorner := func() Shaper {
		corner := NewSphere()
		b := NewTranslation(0, 0, -1)
		b = b.MultiplyMatrix(NewScaling(0.25, 0.25, 0.25))
		corner.SetTransform(b)
		return corner
	}

	hexagonEdge := func() Shaper {
		edge := NewCylinder()
		edge.SetMinimum(0)
		edge.SetMaximum(1)

		b := IdentityMatrix()
		b = b.MultiplyMatrix(NewTranslation(0, 0, -1))
		b = b.MultiplyMatrix(NewRotationY(-math.Pi / 6))
		b = b.MultiplyMatrix(NewRotationZ(-math.Pi / 2))
		b = b.MultiplyMatrix(NewScaling(0.25, 1, 0.25))

		edge.SetTransform(b)
		return edge
	}

	hexagonSide := func() *Group {
		side := NewGroup()
		b := hexagonCorner()
		side.AddShape(&b)
		c := hexagonEdge()
		side.AddShape(&c)
		return side
	}

	hexagon := func() *Group {
		hex := NewGroup()

		for n := 0.0; n < 6.0; n += 1.0 {
			side := hexagonSide()
			side.SetTransform(NewRotationY(n * math.Pi / 3))
			hex.AddGroup(side)
		}
		return hex
	}

	bob := hexagon()
	b := NewRotationX(2.4 * math.Pi / 6)
	b = b.MultiplyMatrix(NewTranslation(0, 5, 1))
	bob.SetTransform(b)
	world := NewWorld()
	world.SetLight(NewLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)))
	world.Objects = append(world.Objects, bob)
	camera := NewCamera(600, 300, math.Pi/3)
	camera.SetTransform(ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0)))

	canvas := camera.Render(world)

	tempFile := filepath.Join(os.TempDir(), "chapter14.ppm")
	os.WriteFile(tempFile, []byte(canvas.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
