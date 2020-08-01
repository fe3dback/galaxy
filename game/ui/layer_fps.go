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

	// Table Pos:
	uiInfoX      = 5
	uiInfoY      = 5
	uiInfoFpsX   = uiInfoX
	uiInfoFpsY   = uiInfoY
	uiInfoCamX   = uiInfoX
	uiInfoCamY   = uiInfoFpsY + 15
	uiInfoGraphX = uiInfoX
	uiInfoGraphY = uiInfoCamY + 20
)

type LayerFPS struct {
	moment      engine.Moment
	mousePos    engine.Point
	graph       [graphCapacity]uint8
	graphCursor int
}

func NewLayerFPS() *LayerFPS {
	return &LayerFPS{
		graph:       [graphCapacity]uint8{},
		graphCursor: graphIndexFirst,
	}
}

func (l *LayerFPS) OnUpdate(s engine.State) error {
	l.moment = s.Moment()
	l.mousePos = s.Mouse().MouseCoords()

	metricRate := l.moment.FrameDuration().Seconds() / l.moment.LimitDuration().Seconds()
	metric := uint8(graphHeight * metricRate)
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
		fmt.Sprintf("fps: %d / %s",
			l.moment.FPS(),
			l.moment.FrameDuration().String(),
		),
		engine.Point{
			X: uiInfoFpsX,
			Y: uiInfoFpsY,
		},
	)
	r.DrawText(
		generated.ResourcesFontsJetBrainsMonoRegular,
		engine.ColorGreen,
		fmt.Sprintf("cam: %.2f, %.2f",
			r.Camera().Position().X,
			r.Camera().Position().Y,
		),
		engine.Point{
			X: uiInfoCamX,
			Y: uiInfoCamY,
		},
	)

	l.drawGraph(r)

	// draw mouse
	r.DrawCrossLines(engine.ColorOrange, 3, l.mousePos)

	return nil
}

func (l *LayerFPS) drawGraph(r engine.Renderer) {
	xl := uiInfoGraphX
	xr := uiInfoGraphX + graphWidth
	yb := uiInfoGraphY + graphHeight

	// draw graph bottom border
	r.DrawLine(
		engine.ColorSelection,
		engine.Line{
			A: engine.Point{X: xl, Y: yb + 2},
			B: engine.Point{X: xr, Y: yb + 2},
		},
	)

	// draw graph
	xOffset := 0
	var graphColor engine.Color
	for i := l.graphCursor - graphWidth; i < l.graphCursor; i++ {
		ratePercent := float32(l.graph[i]) / graphHeight

		if ratePercent > 0.75 {
			graphColor = engine.ColorRed
		} else if ratePercent > 0.5 {
			graphColor = engine.ColorOrange
		} else if ratePercent > 0.25 {
			graphColor = engine.ColorYellow
		} else {
			graphColor = engine.ColorSelection
		}

		r.DrawLine(
			graphColor,
			engine.Line{
				A: engine.Point{X: xl + xOffset, Y: yb},
				B: engine.Point{X: xl + xOffset, Y: yb - int(l.graph[i])},
			},
		)

		xOffset++
	}
}
