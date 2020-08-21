package collider

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/collision"
)

func (c *Collider) IsCollideWith(other *Collider) bool {
	switch c.colliderType {
	case TypePoint:
		return other.isCollideWithPoint(*c.point)
	case TypeRect:
		return other.isCollideWithRect(*c.rect)
	case TypeCircle:
		return other.isCollideWithCircle(*c.circle)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}

func (c *Collider) isCollideWithPoint(p engine.Vec) bool {
	switch c.colliderType {
	case TypePoint:
		return c.point.X == p.X && c.point.Y == p.Y
	case TypeRect:
		return collision.Rect2Point(*c.rect, p)
	case TypeCircle:
		return collision.Circle2Point(*c.circle, p)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}

func (c *Collider) isCollideWithRect(r engine.Rect) bool {
	switch c.colliderType {
	case TypePoint:
		return collision.Rect2Point(r, *c.point)
	case TypeRect:
		return collision.Rect2Rect(r, *c.rect)
	case TypeCircle:
		return collision.Rect2Circle(r, *c.circle)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}

func (c *Collider) isCollideWithCircle(d engine.Circle) bool {
	switch c.colliderType {
	case TypePoint:
		return collision.Circle2Point(d, *c.point)
	case TypeRect:
		return collision.Rect2Circle(*c.rect, d)
	case TypeCircle:
		return collision.Circle2Circle(*c.circle, d)
	default:
		panic(fmt.Sprintf("Unknown collider type %d, %T", c.colliderType, c.colliderType))
	}
}
