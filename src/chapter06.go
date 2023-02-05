package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChapterSix() {
	rayOrigin := NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 200.0
	pixelSize := wallSize / canvasPixels
	half := wallSize / 2

	canvas := NewCanvas(int(canvasPixels), int(canvasPixels))
	color := NewColor(1, 0, 0)
	shape := NewSphere()
	shape.Material = NewMaterial()
	shape.Material.Color = NewColor(1, 0.2, 1)

	lightPosition := NewPoint(-10, 10, -10)
	lightColor := NewColor(1, 1, 1)
	light := NewLight(lightPosition, lightColor)

	for y := 0.0; y < canvasPixels; y = y + 1.0 {
		worldY := half - pixelSize*y
		for x := 0.0; x < canvasPixels; x = x + 1.0 {
			worldX := -half + pixelSize*x
			position := NewPoint(worldX, worldY, wallZ)
			r := NewRay(rayOrigin, (position.Subtract(rayOrigin)).Normalize())
			xs := shape.Intersects(r)

			hit, hitInter := Hit(xs)
			if hit {
				point := r.Position(hitInter.T)
				normal := hitInter.Object.NormalAt(point)
				eye := r.Direction.Negative()
				inShadow := false
				color = Lighting(hitInter.Object.GetMaterial(), light, point, eye, normal, inShadow)
				canvas.WritePixel(int(x), int(y), color)
			}
		}
	}

	tempFile := filepath.Join(os.TempDir(), "chapter06.ppm")
	os.WriteFile(tempFile, []byte(canvas.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
