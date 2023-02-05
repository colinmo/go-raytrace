package main

import (
	"fmt"
	"math"
	"strings"
)

type Canvas struct {
	Width  int
	Height int
	Pixels map[int]map[int]Color
}

func NewCanvas(w, h int) Canvas {
	me := Canvas{
		Width:  w,
		Height: h,
	}
	defaultColour := NewColor(0, 0, 0)
	me.Pixels = make(map[int]map[int]Color, h)
	for i := 0; i < w; i++ {
		me.Pixels[i] = make(map[int]Color, w)
		for j := 0; j < h; j++ {
			me.Pixels[i][j] = defaultColour
		}
	}
	return me
}

func (c *Canvas) WritePixel(x, y int, col Color) {
	if x < 0 || x > c.Width || y < 0 || y > c.Height {
		return
	}
	c.Pixels[x][y] = col
}

func (c *Canvas) PixelAt(x, y int) Color {
	return c.Pixels[x][y]
}

func (c *Canvas) ToPPM() string {
	colorDepth := 255
	body := []string{}
	for y := 0; y < c.Height; y++ {
		newRow := ""
		for x := 0; x < c.Width; x++ {
			cell := c.Pixels[x][y]
			red := colorToDepth(cell.Red, colorDepth)
			green := colorToDepth(cell.Green, colorDepth)
			blue := colorToDepth(cell.Blue, colorDepth)
			asString := fmt.Sprintf("%d %d %d ", red, green, blue)
			newRow = newRow + asString
		}
		body = append(body, strings.Trim(newRow, " "))
	}
	toReturn := ""
	for _, row := range body {
	inner:
		for {
			if len(row) < 71 {
				toReturn = toReturn + row + "\n"
				break inner
			}

			splitCol := 69
		checker:
			for {
				if row[splitCol:splitCol+1] == " " {
					toReturn = toReturn + row[0:splitCol] + "\n"
					row = row[splitCol+1:]
					break checker
				}
				splitCol = splitCol - 1
			}
		}
	}

	return fmt.Sprintf("P3\n%d %d\n%d\n%s\n", c.Width, c.Height, colorDepth, toReturn)
}

func colorToDepth(code float64, depth int) int {
	x := int(math.Round(code * float64(depth)))
	if x < 0 {
		x = 0
	}
	if x > depth {
		x = depth
	}
	return x
}
