package render

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

const avgTextWidthOptRender = 150
const avgTextHeightOptRender = 20

func (r *Renderer) cacheText(fontId consts.AssetsPath, color galx.Color, text string) *cachedText {
	key := fmt.Sprintf("%x", md5.Sum([]byte(fontId+strconv.Itoa(int(color))+text)))
	if tex, ok := r.textCache[key]; ok {
		return tex
	}

	r.SetDrawColor(color)

	font := r.fontManager.Get(fontId)
	surface := font.RenderText(text, color)
	defer surface.Free()

	texture, err := r.ref.CreateTextureFromSurface(surface)
	if err != nil {
		utils.Check("create font texture from surface", err)
	}

	r.textCache[key] = &cachedText{
		tex:    texture,
		width:  surface.W,
		height: surface.H,
	}
	return r.textCache[key]
}

func (r *Renderer) internalDrawText(fontId consts.AssetsPath, color galx.Color, text string, pos sdl.Point) {
	tex := r.cacheText(fontId, color, text)

	src := Rect{
		X: 0,
		Y: 0,
		W: tex.width,
		H: tex.height,
	}

	dest := Rect{
		X: pos.X,
		Y: pos.Y,
		W: tex.width,
		H: tex.height,
	}

	err := r.ref.Copy(tex.tex, &src, &dest)
	utils.Check("copy font texture", err)
}
