package ui

import (
	"github.com/fe3dback/galaxy/engine"
)

type LayerCamera struct {
	mouse engine.Point
}

func NewLayerCamera() *LayerCamera {
	return &LayerCamera{}
}

func (l *LayerCamera) OnUpdate(s engine.State) error {
	l.mouse = s.Mouse().MouseCoords()

	s.Camera().CenterOn(l.mouse)
	return nil
}

func (l *LayerCamera) OnDraw(r engine.Renderer) (err error) {
	r.DrawCrossLines(engine.ColorOrange, 5, l.mouse)

	return nil
}
