package game

import (
	"github.com/fe3dback/galaxy/galx"
)

type LookToMouse struct {
	entity galx.GameObject
}

func NewLookToMouse(entity galx.GameObject) *LookToMouse {
	return &LookToMouse{
		entity: entity,
	}
}

func (r *LookToMouse) OnDraw(_ galx.Renderer) error {
	return nil
}

func (r *LookToMouse) OnUpdate(s galx.State) error {
	mouseWorld := s.Mouse().MouseCoords().Add(s.Camera().Position())

	r.entity.SetRotation(
		r.entity.Position().AngleTo(mouseWorld),
	)

	return nil
}
