package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) FillRect(rect engine.Rect) {
	utils.Check("fill", r.ref.FillRect(r.screenRectPtr(rect)))
}

func (r *Renderer) Clear(color engine.Color) {
	r.SetDrawColor(color)

	//err := r.ref.SetClipRect(&sdl.Rect{
	//	X: 0,
	//	Y: 0,
	//	W: int32(r.camera.width),
	//	H: int32(r.camera.height),
	//})
	//utils.Check("set clip rect", err)

	err := r.ref.Clear()
	utils.Check("clear", err)
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
