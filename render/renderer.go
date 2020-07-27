package render

import (
	"image/color"

	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type Renderer struct {
	ref   *sdl.Renderer
	fonts *FontsCollection
}

type Rect = sdl.Rect

func NewRenderer(window *sdl.Window, fonts *FontsCollection, closer *utils.Closer) *Renderer {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	utils.Check("create renderer", err)
	closer.Enqueue(renderer.Destroy)

	return &Renderer{
		ref:   renderer,
		fonts: fonts,
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

func (r *Renderer) FillRect(rect *Rect) {
	utils.Check("fill", r.ref.FillRect(rect))
}

func (r *Renderer) Clear(color color.RGBA) {
	r.SetDrawColor(color)
	utils.Check("clear", r.ref.Clear())
}

func (r *Renderer) DrawText(color color.RGBA, fontId FontId, text string, x, y int32) {
	r.SetDrawColor(color)

	font := r.fonts.Get(fontId)
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

	src := sdl.Rect{
		X: 0,
		Y: 0,
		W: surface.W,
		H: surface.H,
	}

	dest := sdl.Rect{
		X: x,
		Y: y,
		W: surface.W,
		H: surface.H,
	}

	err = r.ref.Copy(texture, &src, &dest)
	utils.Check("copy font texture", err)
}

func (r *Renderer) Present() {
	r.ref.Present()
}
