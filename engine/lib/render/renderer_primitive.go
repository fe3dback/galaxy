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
		r.screenPoint(engine.Vec{X: rect.Min.X, Y: rect.Min.Y}),
		r.screenPoint(engine.Vec{X: rect.Min.X + rect.Max.X, Y: rect.Min.Y}),
		r.screenPoint(engine.Vec{X: rect.Min.X + rect.Max.X, Y: rect.Min.Y + rect.Max.Y}),
		r.screenPoint(engine.Vec{X: rect.Min.X, Y: rect.Min.Y + rect.Max.Y}),
		r.screenPoint(engine.Vec{X: rect.Min.X, Y: rect.Min.Y}),
	})
	utils.Check("draw square", err)
}

func (r *Renderer) DrawSquareEx(color engine.Color, angle engine.Angle, rect engine.Rect) {
	hw := rect.Max.X / 2
	hh := rect.Max.Y / 2

	tl := engine.Vec{X: rect.Min.X - hw, Y: rect.Min.Y - hh}.RotateAround(rect.Min, angle)
	tr := engine.Vec{X: rect.Min.X - hw + rect.Max.X, Y: rect.Min.Y - hh}.RotateAround(rect.Min, angle)
	br := engine.Vec{X: rect.Min.X - hw + rect.Max.X, Y: rect.Min.Y - hh + rect.Max.Y}.RotateAround(rect.Min, angle)
	bl := engine.Vec{X: rect.Min.X - hw, Y: rect.Min.Y - hh + rect.Max.Y}.RotateAround(rect.Min, angle)

	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		r.screenPoint(tl),
		r.screenPoint(tr),
		r.screenPoint(br),
		r.screenPoint(bl),
		r.screenPoint(tl),
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

func (r *Renderer) DrawPoint(color engine.Color, vec engine.Vec) {
	if !r.isPointInsideCamera(vec) {
		return
	}

	r.SetDrawColor(color)
	err := r.ref.DrawPoint(
		int32(r.screenX(vec.X)),
		int32(r.screenY(vec.Y)),
	)
	utils.Check("draw point", err)
}

// -------------------------------------------
// extended function (based on originals)
// -------------------------------------------

func (r *Renderer) DrawVector(color engine.Color, dist float64, vec engine.Vec, angle engine.Angle) {
	target := vec.PolarOffset(dist, angle)

	line := engine.Line{
		A: vec,
		B: target,
	}

	counterDeg := angle.Add(engine.NewAngle(180))
	arrowLeft := engine.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(engine.NewAngle(-30))),
	}
	arrowRight := engine.Line{
		A: target,
		B: target.PolarOffset(6, counterDeg.Add(engine.NewAngle(+30))),
	}

	r.DrawLine(color, line)
	r.DrawLine(color, arrowLeft)
	r.DrawLine(color, arrowRight)
}

func (r *Renderer) DrawCrossLines(color engine.Color, size int, vec engine.Vec) {
	sf := float64(size)

	r.DrawLine(
		color,
		engine.Line{
			A: engine.Vec{X: vec.X - sf, Y: vec.Y - sf},
			B: engine.Vec{X: vec.X + sf, Y: vec.Y + sf},
		},
	)
	r.DrawLine(
		color,
		engine.Line{
			A: engine.Vec{X: vec.X - sf, Y: vec.Y + sf},
			B: engine.Vec{X: vec.X + sf, Y: vec.Y - sf},
		},
	)
}
