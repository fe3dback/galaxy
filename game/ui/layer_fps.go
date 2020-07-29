package ui

import (
	"fmt"
	"image/color"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/render"
)

const (
	graphWidth       = 200
	graphHeight      = 15 // 255 max
	graphCapacity    = graphWidth * 2
	graphIndexFirst  = graphCapacity / 2
	graphIndexLast   = graphCapacity - 1
	graphWindowWidth = graphCapacity / 2
	graphX           = int32(5)
	graphY           = int32(20)
)

var (
	graphBorderColor = color.RGBA{R: 0, G: 255, B: 127, A: 50}
	graphGraphColor  = color.RGBA{R: 255, G: 127, B: 0, A: 30}
)

type LayerFPS struct {
	moment      engine.Moment
	graph       [graphCapacity]uint8
	graphCursor int32
}

func NewLayerFPS() *LayerFPS {
	return &LayerFPS{
		graph:       [graphCapacity]uint8{},
		graphCursor: graphIndexFirst,
	}
}

func (l *LayerFPS) OnUpdate(moment engine.Moment) error {
	l.moment = moment

	metric := graphHeight * uint8(moment.FrameDuration()/moment.LimitDuration())
	if metric > graphHeight {
		metric = graphHeight
	}

	l.graph[l.graphCursor] = metric
	l.graph[l.graphCursor-graphWindowWidth] = metric

	l.graphCursor++

	if l.graphCursor > graphIndexLast {
		l.graphCursor = graphIndexFirst
	}

	return nil
}

func (l *LayerFPS) OnDraw(r *render.Renderer) (err error) {
	r.DrawText(
		render.FontDefaultMono,
		engine.ColorGreen,
		fmt.Sprintf("fps: %d / %s", l.moment.FPS(), l.moment.FrameDuration().String()),
		5, 5,
	)

	l.drawGraph(r)

	return nil
}

func (l *LayerFPS) drawGraph(r *render.Renderer) {
	xl := graphX
	xr := graphX + graphWidth
	yb := graphY + graphHeight

	// draw graph bottom border
	r.DrawLine(
		graphBorderColor,
		sdl.Point{X: xl, Y: yb + 2},
		sdl.Point{X: xr, Y: yb + 2},
	)

	// draw graph
	xOffset := int32(0)
	for i := l.graphCursor - graphWidth; i < l.graphCursor; i++ {
		r.DrawLine(
			graphGraphColor,
			sdl.Point{X: xl + xOffset, Y: yb},
			sdl.Point{X: xl + xOffset, Y: yb - int32(l.graph[i])},
		)

		xOffset++
	}
}
