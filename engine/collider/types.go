package collider

import (
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
		layer        uint8
		colliderType Type
		point        *engine.Vec
		rect         *engine.Rect
		circle       *engine.Circle
	}
)

func NewPointCollider(layer uint8, p engine.Vec) *Collider {
	return &Collider{
		layer:        layer,
		colliderType: TypePoint,
		point:        &p,
	}
}

func NewRectCollider(layer uint8, r engine.Rect) *Collider {
	return &Collider{
		layer:        layer,
		colliderType: TypeRect,
		rect:         &r,
	}
}

func NewCircleCollider(layer uint8, c engine.Circle) *Collider {
	return &Collider{
		layer:        layer,
		colliderType: TypeCircle,
		circle:       &c,
	}
}
