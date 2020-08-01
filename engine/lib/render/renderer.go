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
	camera         *Camera
	renderMode     engine.RenderMode
}

type Rect = sdl.Rect
type Point = sdl.Point

func NewRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *FontManager,
	textureManager *TextureManager,
	camera *Camera,
) *Renderer {
	return &Renderer{
		window:         sdlWindow,
		ref:            sdlRenderer,
		fontManager:    fontManager,
		textureManager: textureManager,
		camera:         camera,
	}
}

func (r *Renderer) SetDrawColor(color engine.Color) {
	utils.Check("set draw color", r.ref.SetDrawColor(color.Split()))
}

func (r *Renderer) Camera() engine.Camera {
	return r.camera
}

func (r *Renderer) SetRenderMode(renderMode engine.RenderMode) {
	r.renderMode = renderMode
}

func (r *Renderer) Origin() *sdl.Renderer {
	return r.ref
}

// -- Base transforms (include camera relative pos)

func (r *Renderer) transformX(x int) int32 {
	if r.renderMode == engine.RenderModeWorld {
		return int32(x - r.camera.rect.X)
	}

	return int32(x)
}

func (r *Renderer) transformY(y int) int32 {
	if r.renderMode == engine.RenderModeWorld {
		return int32(y - r.camera.rect.Y)
	}

	return int32(y)
}

// -- Complex transforms (should depend on base transforms)

func (r *Renderer) transformRect(rect engine.Rect) Rect {
	return Rect{
		X: r.transformX(rect.X),
		Y: r.transformY(rect.Y),
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
		X: r.transformX(point.X),
		Y: r.transformY(point.Y),
	}
}

func (r *Renderer) transformPointRef(point engine.Point) *Point {
	rPoint := r.transformPoint(point)
	return &rPoint
}

func (r *Renderer) transformLine(line engine.Line) []sdl.Point {
	return []sdl.Point{
		r.transformPoint(line.A),
		r.transformPoint(line.B),

		// close lines back will fix render glitches
		r.transformPoint(line.B),
		r.transformPoint(line.A),
	}
}

// -- Camera visibility checks

func (r *Renderer) isLineInsideCamera(line engine.Line) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return r.isRectInsideCamera(engine.Rect{
		X: line.A.X,
		Y: line.A.Y,
		W: line.B.X - line.A.X,
		H: line.B.Y - line.A.Y,
	})
}

func (r *Renderer) isRectInsideCamera(rect engine.Rect) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	cam := r.camera.rect
	return !(rect.X > cam.X+cam.W ||
		rect.X+rect.W < cam.X ||
		rect.Y > cam.Y+cam.H ||
		rect.Y+rect.H < cam.Y)
}

func (r *Renderer) isPointInsideCamera(point engine.Point) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	cam := r.camera.rect
	return point.X >= cam.X &&
		point.Y >= cam.Y &&
		point.X <= cam.X+cam.W &&
		point.Y <= cam.Y+cam.H
}
