package game

import (
	"github.com/fe3dback/galaxy/galx"
)

type CameraFollower struct {
	entity galx.GameObject
	cam    galx.Camera
}

func NewCameraFollower(entity galx.GameObject) *CameraFollower {
	return &CameraFollower{
		entity: entity,
	}
}

func (r *CameraFollower) OnDraw(d galx.Renderer) error {
	if !d.Gizmos().Debug() {
		return nil
	}

	d.DrawSquare(galx.ColorPink, galx.Rect{
		Min: r.cam.Position(),
		Max: galx.Vec{
			X: float64(r.cam.Width() - 1),
			Y: float64(r.cam.Height() - 1),
		},
	})

	return nil
}

func (r *CameraFollower) OnUpdate(s galx.State) error {
	r.cam = s.Camera()
	s.Camera().CenterOn(r.entity.Position())

	return nil
}
