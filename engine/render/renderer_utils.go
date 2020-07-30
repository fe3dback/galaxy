package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) FillRect(rect engine.Rect) {
	utils.Check("fill", r.ref.FillRect(r.transformRectRef(rect)))
}

func (r *Renderer) Clear(color engine.Color) {
	r.SetDrawColor(color)
	utils.Check("clear", r.ref.Clear())
}

func (r *Renderer) Present() {
	r.ref.Present()
}

func transformColor(color engine.Color) sdl.Color {
	r, g, b, a := color.Split()
	return sdl.Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
