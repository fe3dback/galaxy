package ui

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
)

const (
	graphWidth       = 200
	graphHeight      = 15 // 255 max
	graphCapacity    = graphWidth * 2
	graphIndexFirst  = graphCapacity / 2
	graphIndexLast   = graphCapacity - 1
	graphWindowWidth = graphCapacity / 2
	graphX           = 5
	graphY           = 20
)

type LayerFPS struct {
	moment      engine.Moment
	graph       [graphCapacity]uint8
	graphCursor int
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

func (l *LayerFPS) OnDraw(r engine.Renderer) (err error) {
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorGreen,
		fmt.Sprintf("fps: %d / %s", l.moment.FPS(), l.moment.FrameDuration().String()),
		5, 5,
	)

	l.drawGraph(r)

	return nil
}

func (l *LayerFPS) drawGraph(r engine.Renderer) {
	xl := graphX
	xr := graphX + graphWidth
	yb := graphY + graphHeight

	// draw graph bottom border
	r.DrawLine(
		engine.ColorSelection,
		engine.Point{X: xl, Y: yb + 2},
		engine.Point{X: xr, Y: yb + 2},
	)

	// draw graph
	xOffset := 0
	for i := l.graphCursor - graphWidth; i < l.graphCursor; i++ {
		r.DrawLine(
			engine.ColorOrange,
			engine.Point{X: xl + xOffset, Y: yb},
			engine.Point{X: xl + xOffset, Y: yb - int(l.graph[i])},
		)

		xOffset++
	}
}
