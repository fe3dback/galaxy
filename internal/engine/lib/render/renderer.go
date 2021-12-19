package render

import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/utils"
)

// should have at least 1 surface for primary rendering
const surfacesCount galx.RenderTarget = 8

type (
	Renderer struct {
		window         *sdl.Window
		ref            *sdl.Renderer
		fontManager    *FontsManager
		textureManager *TextureManager
		camera         *Camera
		renderMode     galx.RenderMode
		gizmos         galx.Gizmos
		appState       *engine.State
		renderTarget   *renderTarget

		textCache map[string]*cachedText
	}

	renderTarget struct {
		width         int32
		height        int32
		scale         float32
		screenTexture *sdl.Texture // texture for screen draw
		engineGUI     *sdl.Texture // texture for engine GUI
		textureLayers [surfacesCount]*sdl.Texture
	}

	cachedText struct {
		tex    *sdl.Texture
		width  int32
		height int32
	}
)

type Rect = sdl.Rect
type Point = sdl.Point

func NewRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *FontsManager,
	textureManager *TextureManager,
	camera *Camera,
	dispatcher *event.Dispatcher,
	gizmos galx.Gizmos,
	appState *engine.State,
) *Renderer {
	const renderTargetSize = 32

	renderer := &Renderer{
		window:         sdlWindow,
		ref:            sdlRenderer,
		fontManager:    fontManager,
		textureManager: textureManager,
		camera:         camera,
		renderMode:     galx.RenderModeWorld,
		gizmos:         gizmos,
		appState:       appState,
		renderTarget: &renderTarget{
			screenTexture: sdlRenderer.GetRenderTarget(),
			width:         renderTargetSize,
			height:        renderTargetSize,
			scale:         1.0,
		},
		textCache: map[string]*cachedText{},
	}

	// create all render targets
	renderer.renderTarget.scale = 1.0
	renderer.renderTarget.engineGUI = renderer.createSurfaceTexture(renderTargetSize, renderTargetSize)
	for i := galx.RenderTarget(0); i < surfacesCount; i++ {
		renderer.renderTarget.textureLayers[i] = renderer.createSurfaceTexture(renderTargetSize, renderTargetSize)
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
			float32(cameraUpdateEvent.Scale),
		)

		return nil
	})

	renderer.onWindowResize()
	return renderer
}

func (r *Renderer) SetRenderTarget(id galx.RenderTarget) {
	if id+1 > surfacesCount {
		panic(fmt.Sprintf("can`t draw to surface ID#%d, max surfaces: %d (overflow)",
			id,
			surfacesCount,
		))
	}

	r.renderTo(r.renderTarget.textureLayers[id])
}

func (r *Renderer) MainEngineTexture() uintptr {
	// this return pointer to texture for rendering driver
	// engine GUI will use it for scene image draw
	// all GUI draw bindings, should use it for rendering SDL texture into GPU
	return uintptr(unsafe.Pointer(r.renderTarget.textureLayers[0]))
}

func (r *Renderer) SetDrawColor(color galx.Color) {
	utils.Check("set draw color", r.ref.SetDrawColor(color.Split()))
}

func (r *Renderer) Gizmos() galx.Gizmos {
	return r.gizmos
}

func (r *Renderer) InEditorMode() bool {
	return r.appState.InEditorMode()
}

func (r *Renderer) SetRenderMode(renderMode galx.RenderMode) {
	r.renderMode = renderMode

	if renderMode == galx.RenderModeWorld {
		err := r.ref.SetScale(r.renderTarget.scale, r.renderTarget.scale)
		utils.Check("set render camera world scale", err)
	}

	if renderMode == galx.RenderModeUI {
		// ui scale is always 100%
		err := r.ref.SetScale(1, 1)
		utils.Check("set render camera UI scale", err)
	}
}

func (r *Renderer) Origin() *sdl.Renderer {
	return r.ref
}

func (r *Renderer) TextureQuery(res consts.AssetsPath) galx.TextureInfo {
	tex := r.getTexture(res)

	return galx.TextureInfo{
		Width:  int(tex.Width),
		Height: int(tex.Height),
	}
}

func (r *Renderer) onWindowResize() {
	width, height := r.window.GetSize()
	r.camera.Resize(int(width), int(height))
}

func (r *Renderer) onCameraUpdate(width int32, height int32, scale float32) {
	flags := r.window.GetFlags()
	fullScreen := flags&sdl.WINDOW_FULLSCREEN != 0

	log.Printf("Resize to [%dx%d] (fullScreen = %v, scale = %v)\n",
		width,
		height,
		fullScreen,
		scale,
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

	// resize all surfaces
	r.renderTarget.scale = scale
	r.renderTarget.width = width
	r.renderTarget.height = height

	r.renderTarget.engineGUI = r.resizeSurfaceTexture(r.renderTarget.engineGUI, width, height)
	for i := galx.RenderTarget(0); i < surfacesCount; i++ {
		r.renderTarget.textureLayers[i] = r.resizeSurfaceTexture(r.renderTarget.textureLayers[i], width, height)
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

	err = tex.SetBlendMode(sdl.BLENDMODE_NONE)
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
