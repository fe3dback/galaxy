package collider

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
)

const (
	TypePoint Type = iota
	TypeRect
	TypeCircle
)

type (
	Type uint8

	Collider struct {
		colliderType Type
		point        *engine.Vec
		rect         *engine.Rect
		circle       *engine.Circle
	}
)

func NewPointCollider(p engine.Vec) *Collider {
	return &Collider{
		colliderType: TypePoint,
		point:        &p,
	}
}

func NewRectCollider(r engine.Rect) *Collider {
	return &Collider{
		colliderType: TypeRect,
		rect:         &r,
	}
}

func NewCircleCollider(c engine.Circle) *Collider {
	return &Collider{
		colliderType: TypeCircle,
		circle:       &c,
	}
}

func (c *Collider) Update(e Entity) {
	switch c.colliderType {
	case TypePoint:
		c.updatePoint(e)
	case TypeRect:
		c.updateRect(e)
	case TypeCircle:
		c.updateCircle(e)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}

func (c *Collider) updatePoint(e Entity) {
	*c.point = e.Position()
}

func (c *Collider) updateRect(e Entity) {
	w, h := c.rect.Width(), c.rect.Height()
	offset := engine.Vec{X: w / 2, Y: h / 2}

	c.rect.Min = e.Position().Sub(offset)
	c.rect.Max = e.Position().Add(offset)
}

func (c *Collider) updateCircle(e Entity) {
	c.circle.Pos = e.Position()
}
