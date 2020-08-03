package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/lib/event"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const fullScreenScaleFactor = 1

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
	dispatcher *event.Dispatcher,
) *Renderer {
	renderer := &Renderer{
		window:         sdlWindow,
		ref:            sdlRenderer,
		fontManager:    fontManager,
		textureManager: textureManager,
		camera:         camera,
	}

	dispatcher.OnWindow(func(window event.EvWindow) error {
		if window.EventType == event.WindowEventTypeSizeChanged {
			renderer.resetView()
		}

		return nil
	})

	renderer.resetView()
	return renderer
}

func (r *Renderer) resetView() {
	width, height := r.window.GetSize()
	flags := r.window.GetFlags()
	fullscreen := flags&sdl.WINDOW_FULLSCREEN != 0

	if fullscreen {
		err := r.ref.SetViewport(&sdl.Rect{
			X: 0,
			Y: 0,
			W: width * fullScreenScaleFactor,
			H: height * fullScreenScaleFactor,
		})
		utils.Check("set viewport", err)

		err = r.ref.SetScale(fullScreenScaleFactor, fullScreenScaleFactor)
		utils.Check("set scale", err)
	}

	fmt.Printf("Resize to %d, %d (fullscreen = %v, scale = %d)\n", width, height, fullscreen, fullScreenScaleFactor)

	// r.Camera().Resize(int(width), int(height)) // todo
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

func (r *Renderer) screenX(x float64) float64 {
	return x //todo
	if r.renderMode == engine.RenderModeUI {
		return x
	}

	return x - r.camera.position.X
}

func (r *Renderer) screenY(y float64) float64 {
	return y
	if r.renderMode == engine.RenderModeUI {
		return y
	}

	return y - r.camera.position.Y
}

// -- Complex transforms (should depend on base transforms)

func (r *Renderer) screenRect(rect engine.Rect) Rect {
	return Rect{
		X: int32(r.screenX(rect.Min.X)),
		Y: int32(r.screenY(rect.Min.Y)),
		W: int32(rect.Max.X),
		H: int32(rect.Max.Y),
	}
}

func (r *Renderer) screenRectPtr(rect engine.Rect) *Rect {
	rRect := r.screenRect(rect)
	return &rRect
}

func (r *Renderer) screenPoint(point engine.Vec) Point {
	return Point{
		X: int32(r.screenX(point.X)),
		Y: int32(r.screenY(point.Y)),
	}
}

func (r *Renderer) screenPointPtr(point engine.Vec) *Point {
	rPoint := r.screenPoint(point)
	return &rPoint
}

func (r *Renderer) screenLine(line engine.Line) []sdl.Point {
	return []sdl.Point{
		r.screenPoint(line.A),
		r.screenPoint(line.B),

		// close lines back will fix render glitches
		r.screenPoint(line.B),
		r.screenPoint(line.A),
	}
}

// -- Camera visibility checks

func (r *Renderer) isLineInsideCamera(line engine.Line) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return r.isRectInsideCamera(engine.Rect{
		Min: line.A,
		Max: line.B.Sub(line.A),
	}.Screen())
}

func (r *Renderer) isRectInsideCamera(rect engine.Rect) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	xB := rect.Min.X > r.camera.position.X+float64(r.camera.width)
	xL := rect.Min.X+rect.Max.X < r.camera.position.X

	yB := rect.Min.Y > r.camera.position.Y+float64(r.camera.height)
	yL := rect.Min.Y+rect.Max.Y < r.camera.position.Y

	return !(xB || xL || yB || yL)
}

func (r *Renderer) isPointInsideCamera(vec engine.Vec) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return vec.X >= r.camera.position.X &&
		vec.Y >= r.camera.position.Y &&
		vec.X <= r.camera.position.X+float64(r.camera.width) &&
		vec.Y <= r.camera.position.Y+float64(r.camera.height)
}
