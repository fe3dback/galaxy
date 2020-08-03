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
	uiInfoX           = 5
	uiInfoY           = 5
	uiInfoFpsX        = uiInfoX
	uiInfoFpsY        = uiInfoY
	uiInfoCamX        = uiInfoX
	uiInfoCamY        = uiInfoFpsY + 15
	uiInfoGraphFpsX   = uiInfoX
	uiInfoGraphFpsY   = uiInfoCamY + 20
	uiInfoGraphDeltaX = uiInfoX
	uiInfoGraphDeltaY = uiInfoGraphFpsY + 20
)

const (
	graphIdFps graphId = iota
	graphIdDelta
)

type (
	graphId = uint8

	graph struct {
		x, y   int
		cursor int
		metric [graphCapacity]uint8
		width  int
		height int
		middle bool
	}

	LayerFPS struct {
		moment            engine.Moment
		mousePos          engine.Vec
		graphs            map[graphId]*graph
		previousDeltaTime float64
	}
)

func NewLayerFPS() *LayerFPS {
	return &LayerFPS{
		graphs: createGraphs(),
	}
}

func createGraphs() map[graphId]*graph {
	graphs := make(map[graphId]*graph)

	graphs[graphIdFps] = &graph{
		x:      uiInfoGraphFpsX,
		y:      uiInfoGraphFpsY,
		cursor: graphIndexFirst,
		metric: [graphCapacity]uint8{},
		width:  graphWidth,
		height: graphHeight,
		middle: false,
	}

	graphs[graphIdDelta] = &graph{
		x:      uiInfoGraphDeltaX,
		y:      uiInfoGraphDeltaY,
		cursor: graphIndexFirst,
		metric: [graphCapacity]uint8{},
		width:  graphWidth,
		height: graphHeight,
		middle: true,
	}

	return graphs
}

func (g *graph) write(metricRate float64) {
	if metricRate < 0 {
		metricRate = 0
	}

	if metricRate > 1 {
		metricRate = 1
	}

	metric := uint8(float64(g.height) * metricRate)
	if metric > uint8(g.height) {
		metric = uint8(g.height)
	}

	g.metric[g.cursor] = metric
	g.metric[g.cursor-graphWindowWidth] = metric

	g.cursor++

	if g.cursor > graphIndexLast {
		g.cursor = graphIndexFirst
	}
}

func (l *LayerFPS) OnUpdate(s engine.State) error {
	l.moment = s.Moment()
	l.mousePos = s.Mouse().MouseCoords()

	fpsRate := l.moment.FrameDuration().Seconds() / l.moment.LimitDuration().Seconds()
	l.graphs[graphIdFps].write(fpsRate)

	deltaA := l.previousDeltaTime
	deltaB := l.moment.DeltaTime()

	deltaRate := (((deltaA + deltaB) / 2) / deltaB) - 0.5
	l.graphs[graphIdDelta].write(deltaRate)

	l.previousDeltaTime = l.moment.DeltaTime()
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
		engine.Vec{
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
		engine.Vec{
			X: uiInfoCamX,
			Y: uiInfoCamY,
		},
	)

	l.drawGraph(r, l.graphs[graphIdFps])
	l.drawGraph(r, l.graphs[graphIdDelta])

	// draw mouse
	r.DrawCrossLines(engine.ColorOrange, 3, l.mousePos)

	return nil
}

func (l *LayerFPS) drawGraph(r engine.Renderer, g *graph) {
	xl := g.x
	xr := g.x + g.width
	yb := g.y + g.height

	if g.middle {
		yb -= g.height / 2
	}

	// draw graph bottom border
	r.DrawLine(
		engine.ColorSelection,
		engine.Line{
			A: engine.Vec{X: float64(xl), Y: float64(g.y + g.height + 2)},
			B: engine.Vec{X: float64(xr), Y: float64(g.y + g.height + 2)},
		},
	)

	// draw graph
	xOffset := 0
	var graphColor engine.Color
	for i := g.cursor - g.width; i < g.cursor; i++ {
		ratePercent := float32(g.metric[i]) / float32(g.height)

		if g.middle {
			ratePercent -= 0.5
			graphColor = engine.ColorPink
		} else {
			if ratePercent > 0.75 {
				graphColor = engine.ColorRed
			} else if ratePercent > 0.5 {
				graphColor = engine.ColorOrange
			} else if ratePercent > 0.25 {
				graphColor = engine.ColorYellow
			} else {
				graphColor = engine.ColorSelection
			}
		}

		r.DrawLine(
			graphColor,
			engine.Line{
				A: engine.Vec{X: float64(xl + xOffset), Y: float64(yb)},
				B: engine.Vec{X: float64(xl + xOffset), Y: float64(yb - int(ratePercent*float32(g.height)))},
			},
		)

		xOffset++
	}
}
