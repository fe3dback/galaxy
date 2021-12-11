package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

const avgTextWidthOptRender = 150
const avgTextHeightOptRender = 20

func (r *Renderer) internalDrawText(fontId consts.AssetsPath, color galx.Color, text string, pos sdl.Point) {
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
		X: pos.X,
		Y: pos.Y,
		W: surface.W,
		H: surface.H,
	}

	err = r.ref.Copy(texture, &src, &dest)
	utils.Check("copy font texture", err)
}
