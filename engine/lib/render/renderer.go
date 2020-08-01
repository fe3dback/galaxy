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

	r.Camera().Resize(int(width), int(height))
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

func (r *Renderer) screenX(x int) int32 {
	if r.renderMode == engine.RenderModeUI {
		return int32(x)
	}

	return int32(x - int(r.camera.position.X))
}

func (r *Renderer) screenY(y int) int32 {
	if r.renderMode == engine.RenderModeUI {
		return int32(y)
	}

	return int32(y - int(r.camera.position.Y))
}

// -- Complex transforms (should depend on base transforms)

func (r *Renderer) screenRect(rect engine.Rect) Rect {
	return Rect{
		X: r.screenX(rect.X),
		Y: r.screenY(rect.Y),
		W: int32(rect.W),
		H: int32(rect.H),
	}
}

func (r *Renderer) screenRectPtr(rect engine.Rect) *Rect {
	rRect := r.screenRect(rect)
	return &rRect
}

func (r *Renderer) screenPoint(point engine.Point) Point {
	return Point{
		X: r.screenX(point.X),
		Y: r.screenY(point.Y),
	}
}

func (r *Renderer) screenPointPtr(point engine.Point) *Point {
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
