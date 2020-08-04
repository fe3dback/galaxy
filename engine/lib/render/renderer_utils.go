package render

import (
	"github.com/fe3dback/galaxy/engine"

	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) FillRect(rect engine.Rect) {
	utils.Check("fill", r.ref.FillRect(r.transRectPtr(rect)))
}

func (r *Renderer) Clear(color engine.Color) {
	r.SetDrawColor(color)
	err := r.ref.Clear()
	utils.Check("clear", err)
}

func (r *Renderer) Present() {
	r.ref.Present()
}
