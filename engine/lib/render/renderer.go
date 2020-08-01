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
		return int32(x - int(r.camera.position.X))
	}

	return int32(x)
}

func (r *Renderer) transformY(y int) int32 {
	if r.renderMode == engine.RenderModeWorld {
		return int32(y - int(r.camera.position.Y))
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

	return !(rect.X > int(r.camera.position.X)+r.camera.width ||
		rect.X+rect.W < int(r.camera.position.X) ||
		rect.Y > int(r.camera.position.Y)+r.camera.height ||
		rect.Y+rect.H < int(r.camera.position.Y))
}

func (r *Renderer) isPointInsideCamera(point engine.Point) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return point.X >= int(r.camera.position.X) &&
		point.Y >= int(r.camera.position.Y) &&
		point.X <= int(r.camera.position.X)+r.camera.width &&
		point.Y <= int(r.camera.position.Y)+r.camera.height
}
