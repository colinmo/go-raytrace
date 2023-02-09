package main

import "math"

type Camera struct {
	HSize       int64
	VSize       int64
	FieldOfView float64
	Transform   Matrix
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
}

func NewCamera(h, v int64, f float64) Camera {
	bob := Camera{
		HSize:       h,
		VSize:       v,
		FieldOfView: f,
		Transform:   IdentityMatrix(),
	}
	bob.CalcPixelSize()
	return bob
}

func (c *Camera) CalcPixelSize() {
	halfView := math.Tan(c.FieldOfView / 2)
	aspect := float64(c.HSize) / float64(c.VSize)
	if aspect > 1 {
		c.HalfWidth = halfView
		c.HalfHeight = halfView / aspect
	} else {
		c.HalfWidth = halfView * aspect
		c.HalfHeight = halfView
	}
	c.PixelSize = c.HalfWidth * 2 / float64(c.HSize)

}

func (c *Camera) RayForPixel(px, py int64) Ray {
	xoffset := (float64(px) + 0.5) * c.PixelSize
	yoffset := (float64(py) + 0.5) * c.PixelSize
	worldX := c.HalfWidth - xoffset
	worldY := c.HalfHeight - yoffset

	m := c.Transform.Inverse()
	pixel := m.MultiplyTuple(NewPoint(worldX, worldY, -1))
	origin := m.MultiplyTuple(NewPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()

	return NewRay(origin, direction)
}

func (c *Camera) SetTransform(t Matrix) {
	c.Transform = t
}

func (c *Camera) Render(w World) Canvas {
	image := NewCanvas(int(c.HSize), int(c.VSize))

	var y, x int64
	for y = 0; y < c.VSize; y++ {
		for x = 0; x < c.HSize; x++ {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray, maxReflects)
			image.WritePixel(int(x), int(y), color)
		}
	}
	return image
}
