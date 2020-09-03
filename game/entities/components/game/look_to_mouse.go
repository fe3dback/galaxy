package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type LookToMouse struct {
	entity *entity.Entity
}

func NewLookToMouse(entity *entity.Entity) *LookToMouse {
	return &LookToMouse{
		entity: entity,
	}
}

func (r *LookToMouse) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *LookToMouse) OnUpdate(s engine.State) error {
	mouseWorld := s.Mouse().MouseCoords().Add(s.Camera().Position())

	r.entity.SetRotation(
		r.entity.Position().AngleTo(mouseWorld),
	)

	return nil
}
