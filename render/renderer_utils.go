package render

import (
	"image/color"

	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) FillRect(rect *Rect) {
	utils.Check("fill", r.ref.FillRect(rect))
}

func (r *Renderer) Clear(color color.RGBA) {
	r.SetDrawColor(color)
	utils.Check("clear", r.ref.Clear())
}

func (r *Renderer) Present() {
	r.ref.Present()
}
