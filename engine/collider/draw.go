package collider

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

const color = engine.ColorGreen

func (c *Collider) OnDraw(r engine.Renderer) {
	switch c.colliderType {
	case TypePoint:
		c.drawPoint(r)
	case TypeRect:
		c.drawRect(r)
	case TypeCircle:
		c.drawCircle(r)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}

func (c *Collider) drawPoint(r engine.Renderer) {
	r.DrawPoint(color, *c.point)
}

func (c *Collider) drawRect(r engine.Renderer) {
	rs := c.rect.Screen()
	r.DrawSquare(color, engine.Rect{
		Min: rs.Min,
		Max: engine.Vec{
			X: rs.Width(),
			Y: rs.Height(),
		},
	})
}

func (c *Collider) drawCircle(_ engine.Renderer) {
	panic("drawCircle not implemented")
}
