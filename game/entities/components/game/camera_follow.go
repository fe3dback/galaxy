package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type CameraFollower struct {
	entity *entity.Entity
	cam    engine.Camera
}

func NewCameraFollower(entity *entity.Entity) *CameraFollower {
	return &CameraFollower{
		entity: entity,
	}
}

func (r *CameraFollower) OnDraw(d engine.Renderer) error {
	d.DrawSquare(engine.ColorPink, engine.Rect{
		Min: r.cam.Position(),
		Max: engine.Vec{
			X: float64(r.cam.Width()),
			Y: float64(r.cam.Height()),
		},
	})

	return nil
}

func (r *CameraFollower) OnUpdate(s engine.State) error {
	r.cam = s.Camera()
	// s.Camera().CenterOn(r.entity.Position()) // todo
	s.Camera().CenterOn(s.Mouse().MouseCoords())

	return nil
}
