package render

import (
	"github.com/fe3dback/galaxy/engine"
)

type Camera struct {
	position engine.Vector2D
	width    int
	height   int
}

func NewCamera(position engine.Vector2D, width, height int) *Camera {
	return &Camera{
		position: position,
		width:    width,
		height:   height,
	}
}

func (c *Camera) Position() engine.Vector2D {
	return c.position
}

func (c *Camera) Width() int {
	return c.width
}

func (c *Camera) Height() int {
	return c.height
}

func (c *Camera) MoveTo(p engine.Vector2D) {
	c.position = p
}

func (c *Camera) CenterOn(p engine.Vector2D) {
	c.MoveTo(engine.Vector2D{
		X: p.X - float64(c.width)/2,
		Y: p.Y - float64(c.height)/2,
	})
}
