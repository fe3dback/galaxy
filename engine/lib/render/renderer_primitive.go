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

func (r *Renderer) DrawSquareEx(color engine.Color, angle engine.Angle, rect engine.Rect) {
	orig := engine.Vector2D{
		X: float64(rect.X),
		Y: float64(rect.Y),
	}

	hw := rect.W / 2
	hh := rect.H / 2

	tl := engine.Vector2D{X: float64(rect.X - hw), Y: float64(rect.Y - hh)}.RotateAround(orig, angle)
	tr := engine.Vector2D{X: float64(rect.X - hw + rect.W), Y: float64(rect.Y - hh)}.RotateAround(orig, angle)
	br := engine.Vector2D{X: float64(rect.X - hw + rect.W), Y: float64(rect.Y - hh + rect.H)}.RotateAround(orig, angle)
	bl := engine.Vector2D{X: float64(rect.X - hw), Y: float64(rect.Y - hh + rect.H)}.RotateAround(orig, angle)

	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		r.screenPoint(tl.ToPoint()),
		r.screenPoint(tr.ToPoint()),
		r.screenPoint(br.ToPoint()),
		r.screenPoint(bl.ToPoint()),
		r.screenPoint(tl.ToPoint()),
	})
	utils.Check("draw square angled", err)
}

func (r *Renderer) DrawLine(color engine.Color, line engine.Line) {
	if !r.isLineInsideCamera(line) {
		return
	}

	r.SetDrawColor(color)
	err := r.ref.DrawLines(r.screenLine(line))
	utils.Check("draw line", err)
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

// -------------------------------------------
// extended function (based on originals)
// -------------------------------------------

func (r *Renderer) DrawVector(color engine.Color, dist float64, vec engine.Vector2D, angle engine.Angle) {
	target := vec.PolarOffset(dist, angle)

	line := engine.Line{
		A: vec.ToPoint(),
		B: target.ToPoint(),
	}

	counterDeg := angle.Add(180)
	arrowLeft := engine.Line{
		A: target.ToPoint(),
		B: target.PolarOffset(6, counterDeg-30).ToPoint(),
	}
	arrowRight := engine.Line{
		A: target.ToPoint(),
		B: target.PolarOffset(6, counterDeg+30).ToPoint(),
	}

	r.DrawLine(color, line)
	r.DrawLine(color, arrowLeft)
	r.DrawLine(color, arrowRight)
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
