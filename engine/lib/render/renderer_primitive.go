package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) internalDrawSquare(color engine.Color, rect sdl.Rect) {
	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		{X: rect.X, Y: rect.Y},
		{X: rect.X + rect.W, Y: rect.Y},
		{X: rect.X + rect.W, Y: rect.Y + rect.H},
		{X: rect.X, Y: rect.Y + rect.H},
		{X: rect.X, Y: rect.Y},
	})
	utils.Check("draw square", err)
}

func (r *Renderer) internalDrawLines(color engine.Color, line []sdl.Point) {
	r.SetDrawColor(color)
	err := r.ref.DrawLines(line)
	utils.Check("draw line", err)
}

func (r *Renderer) internalDrawPoint(color engine.Color, pos sdl.Point) {
	r.SetDrawColor(color)
	err := r.ref.DrawPoint(pos.X, pos.Y)
	utils.Check("draw point", err)
}

// -------------------------------------------
// extended function (based on originals)
// -------------------------------------------

func (r *Renderer) internalDrawCrossLines(color engine.Color, size int32, pos sdl.Point) {
	r.internalDrawLines(
		color,
		[]sdl.Point{
			{X: pos.X - size, Y: pos.Y - size},
			{X: pos.X + size, Y: pos.Y + size},
		},
	)
	r.internalDrawLines(
		color,
		[]sdl.Point{
			{X: pos.X - size, Y: pos.Y + size},
			{X: pos.X + size, Y: pos.Y - size},
		},
	)
}
