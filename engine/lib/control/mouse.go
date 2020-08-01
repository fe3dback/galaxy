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

func (m Mouse) MouseCoords() engine.Point {
	x, y, _ := sdl.GetMouseState()

	return engine.Point{
		X: int(x),
		Y: int(y),
	}
}
