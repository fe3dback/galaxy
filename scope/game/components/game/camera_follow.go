package game

import (
	"github.com/fe3dback/galaxy/galx"
)

type CameraFollower struct {
	entity galx.GameObject
	cam    galx.Camera
}

func (r CameraFollower) Id() string {
	return "4c8b05b5-a006-4d44-8cf8-2132cf1010de"
}

func (r CameraFollower) Title() string {
	return "Game.Camera Follower"
}

func (r CameraFollower) Description() string {
	return "Component will lock camera on entity, and update it position each frame"
}

func (r *CameraFollower) OnCreated(entity galx.GameObject) {
	r.entity = entity
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
