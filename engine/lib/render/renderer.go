package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/fe3dback/galaxy/generated"
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

	dispatcher.OnWindow(func(window event.WindowEvent) error {
		if window.EventType == event.WindowEventTypeSizeChanged {
			renderer.onWindowResize()
		}

		return nil
	})

	dispatcher.OnCameraUpdate(func(cameraUpdateEvent event.CameraUpdateEvent) error {
		renderer.onCameraUpdate(
			int32(cameraUpdateEvent.Width),
			int32(cameraUpdateEvent.Height),
			cameraUpdateEvent.Zoom,
		)

		return nil
	})

	renderer.onWindowResize()
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

func (r *Renderer) onWindowResize() {
	width, height := r.window.GetSize()
	r.Camera().Resize(int(width), int(height))
}

func (r *Renderer) onCameraUpdate(width int32, height int32, zoom float64) {
	flags := r.window.GetFlags()
	fullScreen := flags&sdl.WINDOW_FULLSCREEN != 0

	fmt.Printf("Resize to [%dx%d] (fullScreen = %v, zoom = %v)\n",
		width,
		height,
		fullScreen,
		zoom,
	)

	err := r.ref.SetLogicalSize(width, height)
	utils.Check("set logical size", err)

	err = r.ref.SetViewport(&sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	})
	utils.Check("set viewport", err)

	err = r.ref.SetClipRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: width,
		H: height,
	})
	utils.Check("set clip rect", err)

	err = r.ref.SetScale(float32(r.camera.zoom), float32(r.camera.zoom))
	utils.Check("scale (zoom) rect", err)
}
