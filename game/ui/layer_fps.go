package ui

import (
	"math/rand"
	"strconv"

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
	//fps := strconv.Itoa(l.fpsProvider.FPS())
	fps := strconv.Itoa(60 + rand.Intn(60))

	l.renderer.Clear(engine.ColorBlack)
	l.renderer.DrawText(engine.ColorGreen, render.FontDefaultMono, fps, 5, 5)
	l.renderer.Present()

	return nil
}
