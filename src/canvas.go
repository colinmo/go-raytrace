package main

import (
	"fmt"
	"math"
)

// Color is a 3 element array representing a color as Red, Green, Blue
type Color struct {
	Red, Green, Blue float64
}

func (c1 Color) add(c2 Color) Color {
	return Color{c1.Red + c2.Red, c1.Green + c2.Green, c1.Blue + c2.Blue}
}

func (c1 Color) sub(c2 Color) Color {
	return Color{c1.Red - c2.Red, c1.Green - c2.Green, c1.Blue - c2.Blue}
}

func (c1 Color) multS(scalar float64) Color {
	return Color{c1.Red * scalar, c1.Green * scalar, c1.Blue * scalar}
}
func (c1 Color) multC(c2 Color) Color {
	return Color{c1.Red * c2.Red, c1.Green * c2.Green, c1.Blue * c2.Blue}
}

func (c1 Color) equals(c2 Color) bool {
	return epsilonEquals(c1.Red, c2.Red) &&
		epsilonEquals(c1.Green, c2.Green) &&
		epsilonEquals(c1.Blue, c2.Blue)
}

// Canvas is the 2 dimensional array of pixels (Color)
type Canvas struct {
	W, H int
	P    [][]Color
}

func (c Canvas) writePixel(x, y int, col Color) {
	c.P[x][y] = col
}

func (c Canvas) pixelAt(x, y int) Color {
	return c.P[x][y]
}

func clampColor(col float64) int {
	col = col * 255
	clamp := int(math.Round(math.Max(math.Min(col, 255), 0)))
	return clamp
}
func (c Canvas) dump() {
	fmt.Println(c)
}
func (c Canvas) toPPM() string {
	s := fmt.Sprintf("P3\n%d %d\n%d\n", c.W, c.H, 255)
	for j := 0; j < c.H; j++ {
		l := ""
		for i := 0; i < c.W; i++ {
			l = fmt.Sprintf("%s%d %d %d ",
				l,
				clampColor(c.pixelAt(i, j).Red),
				clampColor(c.pixelAt(i, j).Green),
				clampColor(c.pixelAt(i, j).Blue))
		}
		// Ensure no line is too long
		if len(l) > 70 {
			for len(l) > 70 {
				splitPoint := 70
				for l[splitPoint:splitPoint+1] != " " {
					splitPoint = splitPoint - 1
				}
				s = fmt.Sprintf("%s%s\n", s, l[:splitPoint])
				l = l[splitPoint+1:]
			}

		}
		s = fmt.Sprintf("%s%s\n", s, l[:len(l)-1])
	}
	return s
}

// NewCanvas creates a Canvas
func NewCanvas(x, y int) Canvas {
	p := [][]Color{}
	for i := 0; i < x; i++ {
		p = append(p, make([]Color, y))
	}
	return Canvas{x, y, p}
}
