package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) DrawSquare(color engine.Color, rect engine.Rect) {
	if !r.isRectInsideCamera(rect) {
		return
	}

	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		r.screenPoint(engine.Point{X: rect.X, Y: rect.Y}),
		r.screenPoint(engine.Point{X: rect.X + rect.W, Y: rect.Y}),
		r.screenPoint(engine.Point{X: rect.X + rect.W, Y: rect.Y + rect.H}),
		r.screenPoint(engine.Point{X: rect.X, Y: rect.Y + rect.H}),
		r.screenPoint(engine.Point{X: rect.X, Y: rect.Y}),
	})
	utils.Check("draw square", err)
}

func (r *Renderer) DrawLine(color engine.Color, line engine.Line) {
	if !r.isLineInsideCamera(line) {
		return
	}

	r.SetDrawColor(color)
	err := r.ref.DrawLines(r.screenLine(line))
	utils.Check("draw line", err)
}

func (r *Renderer) DrawCrossLines(color engine.Color, size int, point engine.Point) {
	r.DrawLine(
		color,
		engine.Line{
			A: engine.Point{X: point.X - size, Y: point.Y - size},
			B: engine.Point{X: point.X + size, Y: point.Y + size},
		},
	)
	r.DrawLine(
		color,
		engine.Line{
			A: engine.Point{X: point.X - size, Y: point.Y + size},
			B: engine.Point{X: point.X + size, Y: point.Y - size},
		},
	)
}

func (r *Renderer) DrawPoint(color engine.Color, point engine.Point) {
	if !r.isPointInsideCamera(point) {
		return
	}

	r.SetDrawColor(color)
	err := r.ref.DrawPoint(
		r.screenX(point.X),
		r.screenY(point.Y),
	)
	utils.Check("draw point", err)
}
