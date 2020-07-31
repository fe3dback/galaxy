package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) DrawSquare(color engine.Color, x, y, w, h int) {
	r.SetDrawColor(color)
	err := r.ref.DrawLines([]sdl.Point{
		{X: int32(x), Y: int32(y)},
		{X: int32(x) + int32(w), Y: int32(y)},
		{X: int32(x) + int32(w), Y: int32(y) + int32(h)},
		{X: int32(x), Y: int32(y) + int32(h)},
		{X: int32(x), Y: int32(y)},
	})
	utils.Check("draw square", err)
}

func (r *Renderer) DrawLine(col engine.Color, a, b engine.Point) {
	r.SetDrawColor(col)
	err := r.ref.DrawLines([]sdl.Point{
		r.transformPoint(a),
		r.transformPoint(b),
	})
	utils.Check("draw line", err)
}
