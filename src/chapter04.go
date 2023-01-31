package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ChapterFour() {
	rayOrigin := NewPoint(0, 0, -5)
	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 200.0
	pixelSize := wallSize / canvasPixels
	half := wallSize / 2

	canvas := NewCanvas(int(canvasPixels), int(canvasPixels))
	color := NewColor(1, 0, 0)
	shape := NewSphere()

	for y := 0.0; y < canvasPixels; y = y + 1.0 {
		worldY := half - pixelSize*y
		for x := 0.0; x < canvasPixels; x = x + 1.0 {
			worldX := -half + pixelSize*x
			position := NewPoint(worldX, worldY, wallZ)
			r := NewRay(rayOrigin, (position.Subtract(rayOrigin)).Normalize())
			xs := shape.Intersects(r)

			hit, _ := Hit(xs)
			if hit {
				canvas.WritePixel(int(x), int(y), color)
			}
		}
	}

	tempFile := filepath.Join(os.TempDir(), "chapter04.ppm")
	os.WriteFile(tempFile, []byte(canvas.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
