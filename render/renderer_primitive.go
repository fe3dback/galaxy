package render

import (
	"image/color"

	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) DrawSquare(col color.RGBA, x, y, w, h int) {
	// todo make all engine staff compatible with sdl
	// int -> int32, etc..

	err := r.ref.SetDrawColor(col.R, col.G, col.B, col.A)
	utils.Check("set color", err)

	err = r.ref.DrawLines([]sdl.Point{
		{X: int32(x), Y: int32(y)},
		{X: int32(x) + int32(w), Y: int32(y)},
		{X: int32(x) + int32(w), Y: int32(y) + int32(h)},
		{X: int32(x), Y: int32(y) + int32(h)},
		{X: int32(x), Y: int32(y)},
	})
	utils.Check("draw square", err)
}

func (r *Renderer) DrawLine(col color.RGBA, a, b sdl.Point) {
	err := r.ref.SetDrawColor(col.R, col.G, col.B, col.A)
	utils.Check("set color", err)

	err = r.ref.DrawLines([]sdl.Point{a, b})
	utils.Check("draw line", err)
}