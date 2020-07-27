package render

import (
	"image/color"

	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	window         *sdl.Window
	ref            *sdl.Renderer
	fontManager    *FontManager
	textureManager *TextureManager
}

type Rect = sdl.Rect

func NewRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *FontManager,
	textureManager *TextureManager,
) *Renderer {
	return &Renderer{
		window:         sdlWindow,
		ref:            sdlRenderer,
		fontManager:    fontManager,
		textureManager: textureManager,
	}
}

func (r *Renderer) SetDrawColor(color color.RGBA) {
	utils.Check("set draw color", r.ref.SetDrawColor(
		color.R,
		color.G,
		color.B,
		color.A,
	))
}

func (r *Renderer) Origin() *sdl.Renderer {
	return r.ref
}
