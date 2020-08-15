package components

import (
	"github.com/fe3dback/galaxy/engine"
)

const cameraSpeed = 50.0

type Camera struct {
	zoom float64
}

func NewCamera() *Camera {
	return &Camera{
		zoom: 1,
	}
}

func (c *Camera) OnUpdate(state engine.State) error {
	// zoom
	lastScroll := state.Mouse().ScrollLastOffset()
	if lastScroll != 0 {
		c.zoom = engine.Clamp(c.zoom+lastScroll*0.5, 0.5, 4)
		state.Camera().ZoomView(c.zoom)
	}

	// move camera
	speed := cameraSpeed
	if state.Movement().Shift() {
		speed *= 5
	}

	state.Camera().MoveTo(
		state.Camera().Position().Add(state.Movement().Vector().Scale(speed)),
	)

	return nil
}

func (c *Camera) OnDraw(_ engine.Renderer) error {
	return nil
}
