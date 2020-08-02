package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type CameraFollower struct {
	entity *entity.Entity
}

func NewCameraFollower(entity *entity.Entity) *CameraFollower {
	return &CameraFollower{
		entity: entity,
	}
}

func (r *CameraFollower) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *CameraFollower) OnUpdate(s engine.State) error {
	s.Camera().CenterOn(r.entity.Position())

	return nil
}
