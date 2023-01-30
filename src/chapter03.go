package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
)

func ChapterThree() {
	c := NewCanvas(100, 100)
	segs := math.Pi / 6
	white := NewColor(1, 1, 1)
	for i := 0.0; i < 12.0; i++ {
		p := NewPoint(0, 0, 1)
		step := NewRotationY(i * segs)
		me := step.MultiplyTuple(p)
		me.X *= 40
		me.Z *= 40
		c.WritePixel(50+int(me.X), 50+int(me.Z), white)
	}
	tempFile := filepath.Join(os.TempDir(), "chapter03.ppm")
	os.WriteFile(tempFile, []byte(c.ToPPM()), 0666)
	fmt.Printf("Open %s\n", tempFile)
}
