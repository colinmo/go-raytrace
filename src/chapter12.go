package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func ChapterTwelve() {

	floor := NewPlane()
	floor.SetTransform(NewScaling(10, 0.01, 10))
	floor.SetMaterial(NewMaterial())
	//floor.Material.SetPattern(NewCheckerPattern(NewColor(1, 1, 1), NewColor(0, 0, 0)))
	floor.Material.SetPattern(NewStripePattern(NewColor(1, 1, 0), NewColor(0, 1, 1)))
	//floor.Material.Color = NewColor(1, 0.9, 0.9)
	floor.Material.Specular = 0

	middle := NewSphere()
	middle.SetTransform(NewTranslation(-0.5, 1, 0.5))
	middle.Material = NewMaterial()
	middle.Material.Color = NewColor(0.1, 1, 0.5)
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3
	middle.Material.Transparency = 0.9
	middle.Material.RefractiveIndex = 0.4
	middle.Material.Reflective = 0.9

	right := NewCube()
	right.SetTransform(NewTranslation(1.5, 0.5, -0.5))
	right.SetTransform(right.Transform.MultiplyMatrix(NewScaling(0.5, 0.5, 0.5)))
	right.Material = NewMaterial()
	right.Material.Color = NewColor(0.5, 1, 0.1)
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3
	//right.Material.SetPattern(NewStripePattern(NewColor(1, 1, 0), NewColor(0, 1, 1)))
	right.Material.SetPattern(NewCheckerPattern(NewColor(1, 1, 1), NewColor(0, 0, 0)))

	left := NewSphere()
	left.SetTransform(NewTranslation(-1.5, 0.33, -0.75))
	left.SetTransform(left.Transform.MultiplyMatrix(NewScaling(0.33, 0.33, 0.33)))
	left.Material = NewMaterial()
	left.Material.Color = NewColor(1, 0.8, 0.1)
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3
	left.Material.SetPattern(NewStripePattern(NewColor(1, 0.7, 0.2), NewColor(0, 1, 0)))

	world := NewWorld()
	world.SetLight(NewLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)))
	world.Objects = append(world.Objects, left)
	world.Objects = append(world.Objects, middle)
	world.Objects = append(world.Objects, right)
	world.Objects = append(world.Objects, floor)
	camera := NewCamera(600, 300, math.Pi/3)
	camera.SetTransform(ViewTransform(NewPoint(0, 1.5, -5), NewPoint(0, 1, 0), NewVector(0, 1, 0)))

	canvas := camera.Render(world)

	tempFile := filepath.Join(os.TempDir(), "chapter12.ppm")
	os.WriteFile(tempFile, []byte(canvas.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
