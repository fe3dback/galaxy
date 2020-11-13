package render

import (
	"fmt"
	"log"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const surfacesCount = 4

type (
	Renderer struct {
		window         *sdl.Window
		ref            *sdl.Renderer
		fontManager    *FontManager
		textureManager *TextureManager
		camera         *Camera
		renderMode     engine.RenderMode
		gizmos         engine.Gizmos
		appState       *engine.AppState
		renderTarget   renderTarget
	}

	renderTarget struct {
		width     int32
		height    int32
		primary   *sdl.Texture
		secondary [surfacesCount]*sdl.Texture
	}
)

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
		renderMode:     engine.RenderModeWorld,
		gizmos:         gizmos,
		appState:       appState,
		renderTarget: renderTarget{
			primary: sdlRenderer.GetRenderTarget(),
		},
	}

	// create all render targets
	for i := 0; i < surfacesCount; i++ {
		renderer.renderTarget.secondary[i] = renderer.createSurfaceTexture(32, 32)
	}

	// subscribe to events
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
			float32(cameraUpdateEvent.Zoom),
		)

		return nil
	})

	renderer.onWindowResize()
	return renderer
}

func (r *Renderer) SetRenderTarget(id uint8) {
	if id == 0 {
		r.renderTo(r.renderTarget.primary)
		return
	}

	if id > surfacesCount {
		panic(fmt.Sprintf("can`t draw to surface #%d, max surfaces: %d",
			id,
			surfacesCount,
		))
	}

	r.renderTo(r.renderTarget.secondary[id-1])
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

func (r *Renderer) onCameraUpdate(width int32, height int32, zoom float32) {
	flags := r.window.GetFlags()
	fullScreen := flags&sdl.WINDOW_FULLSCREEN != 0

	log.Printf("Resize to [%dx%d] (fullScreen = %v, zoom = %v)\n",
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

	err = r.ref.SetScale(zoom, zoom)
	utils.Check("scale (zoom) rect", err)

	// resize all surfaces
	r.renderTarget.width = width
	r.renderTarget.height = height

	for i := 0; i < surfacesCount; i++ {
		r.renderTarget.secondary[i] = r.resizeSurfaceTexture(r.renderTarget.secondary[i], width, height)
	}
}

func (r *Renderer) createSurfaceTexture(width int32, height int32) *sdl.Texture {
	tex, err := r.ref.CreateTexture(
		uint32(sdl.PIXELFORMAT_RGBA32),
		sdl.TEXTUREACCESS_TARGET,
		width,
		height,
	)
	utils.Check("create surface texture", err)

	err = tex.SetBlendMode(sdl.BLENDMODE_ADD)
	utils.Check("set surface texture blend mode", err)

	return tex
}

func (r *Renderer) resizeSurfaceTexture(old *sdl.Texture, width int32, height int32) *sdl.Texture {
	newSurface := r.createSurfaceTexture(width, height)
	r.renderTo(newSurface)

	_, _, oldWidth, oldHeight, err := old.Query()
	utils.Check("query old surface", err)

	src := sdl.Rect{
		X: 0,
		Y: 0,
		W: oldWidth,
		H: oldHeight,
	}

	err = r.ref.Copy(old, &src, &src)
	utils.Check("copy surface texture to new surface", err)

	err = old.Destroy()
	utils.Check("destroy old surface", err)

	return newSurface
}

func (r *Renderer) renderTo(tex *sdl.Texture) {
	err := r.ref.SetRenderTarget(tex)
	utils.Check("set new render target", err)
}
