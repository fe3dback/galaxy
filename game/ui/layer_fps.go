package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/render"
)

type LayerFPS struct {
	fpsProvider FramesProvider
}

func NewLayerFPS(fpsProvider FramesProvider) *LayerFPS {
	return &LayerFPS{
		fpsProvider: fpsProvider,
	}
}

func (l *LayerFPS) OnDraw(r *render.Renderer) (err error) {
	r.DrawText(
		render.FontDefaultMono,
		engine.ColorGreen,
		fmt.Sprintf("fps: %d / %d", l.fpsProvider.FPS(), l.fpsProvider.TotalFPS()),
		5, 5,
	)

	return nil
}
