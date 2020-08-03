package control

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/veandco/go-sdl2/sdl"
)

type Mouse struct {
}

func NewMouse() *Mouse {
	return &Mouse{}
}

func (m Mouse) MouseCoords() engine.Vec {
	x, y, _ := sdl.GetMouseState()

	return engine.Vec{
		X: float64(x),
		Y: float64(y),
	}
}
