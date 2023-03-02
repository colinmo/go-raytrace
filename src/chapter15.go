package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func ChapterFifteen() {

	p := NewParserFromFile(`/private/tmp/go-raytrace/src/fixtures/teapot.obj`)

	world := NewWorld()
	world.SetLight(NewLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)))
	g := p.ToGroup()
	g.SetTransform(NewRotationY(math.Pi / 4))
	world.Objects = append(world.Objects, p.ToGroup())
	camera := NewCamera(200, 100, math.Pi/3)
	camera.SetTransform(ViewTransform(NewPoint(0, 1.5, -4), NewPoint(0, 1, 0), NewVector(0, 1, -5)))

	canvas := camera.Render(world)

	tempFile := filepath.Join(os.TempDir(), "chapter15.ppm")
	os.WriteFile(tempFile, []byte(canvas.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
