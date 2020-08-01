package render

import (
	"github.com/fe3dback/galaxy/engine"
)

type Camera struct {
	rect engine.Rect
}

func NewCamera(rect engine.Rect) *Camera {
	return &Camera{
		rect: rect,
	}
}

func (c *Camera) Rect() engine.Rect {
	return c.rect
}

func (c *Camera) MoveTo(p engine.Point) {
	c.rect.X = p.X
	c.rect.Y = p.Y
}

func (c *Camera) CenterOn(p engine.Point) {
	c.MoveTo(engine.Point{
		X: p.X - c.rect.W/2,
		Y: p.Y - c.rect.H/2,
	})
}
