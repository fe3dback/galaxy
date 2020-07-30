package render

import (
	"github.com/fe3dback/galaxy/engine"
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
type Point = sdl.Point

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

func (r *Renderer) SetDrawColor(color engine.Color) {
	utils.Check("set draw color", r.ref.SetDrawColor(color.Split()))
}

func (r *Renderer) Origin() *sdl.Renderer {
	return r.ref
}

func (r *Renderer) transformRect(rect engine.Rect) Rect {
	return Rect{
		X: int32(rect.X),
		Y: int32(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}
}

func (r *Renderer) transformRectRef(rect engine.Rect) *Rect {
	rRect := r.transformRect(rect)
	return &rRect
}

func (r *Renderer) transformPoint(point engine.Point) Point {
	return Point{
		X: int32(point.X),
		Y: int32(point.Y),
	}
}

func (r *Renderer) transformPointRef(point engine.Point) *Point {
	rPoint := r.transformPoint(point)
	return &rPoint
}
