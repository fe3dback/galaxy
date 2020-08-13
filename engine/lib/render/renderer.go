package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/lib/event"
	"github.com/fe3dback/galaxy/generated"
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
	gizmos         engine.Gizmos
	appState       *engine.AppState
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
	gizmos engine.Gizmos,
	appState *engine.AppState,
) *Renderer {
	renderer := &Renderer{
		window:         sdlWindow,
		ref:            sdlRenderer,
		fontManager:    fontManager,
		textureManager: textureManager,
		camera:         camera,
		gizmos:         gizmos,
		appState:       appState,
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

func (r *Renderer) SetDrawColor(color engine.Color) {
	utils.Check("set draw color", r.ref.SetDrawColor(color.Split()))
}

func (r *Renderer) Camera() engine.Camera {
	return r.camera
}

func (r *Renderer) Gizmos() engine.Gizmos {
	return r.gizmos
}

func (r *Renderer) InEditorMode() bool {
	return r.appState.InEditorState()
}

func (r *Renderer) SetRenderMode(renderMode engine.RenderMode) {
	r.renderMode = renderMode
}

func (r *Renderer) Origin() *sdl.Renderer {
	return r.ref
}

func (r *Renderer) TextureQuery(res generated.ResourcePath) engine.TextureInfo {
	tex := r.getTexture(res)

	return engine.TextureInfo{
		Width:  int(tex.Width),
		Height: int(tex.Height),
	}
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
	r.resetClippingRect()
}

func (r *Renderer) resetClippingRect() {
	err := r.ref.SetClipRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(r.camera.width),
		H: int32(r.camera.height),
	})
	utils.Check("set clip rect", err)
}
