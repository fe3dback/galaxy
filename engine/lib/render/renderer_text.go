package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
)

const avgTextWidthOptRender = 150
const avgTextHeightOptRender = 20

func (r *Renderer) DrawText(fontId generated.ResourcePath, color engine.Color, text string, p engine.Point) {
	if !r.isRectInsideCamera(engine.Rect{
		X: p.X,
		Y: p.Y,
		W: avgTextWidthOptRender,
		H: avgTextHeightOptRender,
	}) {
		return
	}

	r.SetDrawColor(color)

	font := r.fontManager.Get(fontId)
	surface := font.RenderText(text, color)
	defer surface.Free()

	texture, err := r.ref.CreateTextureFromSurface(surface)
	if err != nil {
		utils.Check("create font texture from surface", err)
	}
	defer func() {
		err = texture.Destroy()
		utils.Check("font texture destroy", err)
	}()

	src := Rect{
		X: 0,
		Y: 0,
		W: surface.W,
		H: surface.H,
	}

	dest := Rect{
		X: r.screenX(p.X),
		Y: r.screenY(p.Y),
		W: surface.W,
		H: surface.H,
	}

	err = r.ref.Copy(texture, &src, &dest)
	utils.Check("copy font texture", err)
}
