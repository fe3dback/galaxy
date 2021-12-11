package game

import (
	"github.com/fe3dback/galaxy/galx"
)

type LookToMouse struct {
	entity galx.GameObject
}

func (r LookToMouse) Id() string {
	return "e40e3f52-db31-45cd-a6de-752ba942bd8e"
}

func (r LookToMouse) Title() string {
	return "Game.Look to mouse"
}

func (r LookToMouse) Description() string {
	return "Will lock entity rotation towards current mouse position"
}

func (r *LookToMouse) OnCreated(entity galx.GameObject) {
	r.entity = entity
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
