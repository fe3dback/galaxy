package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/render"
)

type LayerFPS struct {
	renderer    *render.Renderer
	fpsProvider FramesProvider
}

func NewLayerFPS(renderer *render.Renderer, fpsProvider FramesProvider) *LayerFPS {
	return &LayerFPS{
		renderer:    renderer,
		fpsProvider: fpsProvider,
	}
}

func (l *LayerFPS) OnDraw() (err error) {
	l.renderer.DrawText(
		engine.ColorGreen,
		render.FontDefaultMono,
		fmt.Sprintf("fps: %d / %d", l.fpsProvider.FPS(), l.fpsProvider.TotalFPS()),
		5, 5,
	)

	return nil
}
