package lib

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/utils"
)

type SDLLib struct {
	window     *sdl.Window
	renderer   *sdl.Renderer
	renderTech engine.RenderTech
}

func (s *SDLLib) Window() *sdl.Window {
	return s.window
}

func (s *SDLLib) Renderer() *sdl.Renderer {
	return s.renderer
}

func (s *SDLLib) RenderTech() engine.RenderTech {
	return s.renderTech
}

func (s *SDLLib) quit() {
	if s.renderer != nil {
		_ = s.renderer.Destroy()
	}

	if s.window != nil {
		_ = s.window.Destroy()
	}

	sdl.Quit()
}

func NewSDLLib(closer *utils.Closer, defaultWidth, defaultHeight int, fullscreen bool) (*SDLLib, error) {
	platform := &SDLLib{
		renderTech: engine.RenderTechOpenGL2, // strict to opengl2 for now
	}
	closer.EnqueueFree(platform.quit)

	defer utils.CheckPanicWith("sdl lib", func() {
		platform.quit()
	})

	// init sdl
	err := sdl.Init(sdl.INIT_VIDEO & sdl.INIT_EVENTS)
	utils.Check("sdl init", err)

	err = ttf.Init()
	utils.Check("ttf init", err)

	// create main window
	var winFlags uint32
	winFlags |= sdl.WINDOW_SHOWN
	winFlags |= sdl.WINDOW_ALLOW_HIGHDPI
	winFlags |= sdl.WINDOW_BORDERLESS

	if platform.renderTech == engine.RenderTechOpenGL2 || platform.renderTech == engine.RenderTechOpenGL3 {
		winFlags |= sdl.WINDOW_OPENGL
	}

	if platform.renderTech == engine.RenderTechVulkan {
		winFlags |= sdl.WINDOW_VULKAN
	}

	winWidth := defaultWidth
	winHeight := defaultHeight

	if fullscreen {
		mode, displayModeErr := sdl.GetCurrentDisplayMode(0)
		utils.Check("get display mode", displayModeErr)

		winFlags &= sdl.WINDOW_FULLSCREEN
		winWidth = int(mode.W)
		winHeight = int(mode.H)
	}

	window, err := sdl.CreateWindow(
		"Galaxy",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(winWidth), int32(winHeight),
		winFlags,
	)
	utils.Check("create window", err)
	closer.EnqueueClose(window.Destroy)
	platform.window = window

	// set openGL attributes
	switch platform.renderTech {
	case engine.RenderTechOpenGL2:
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 2)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	case engine.RenderTechOpenGL3:
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 2)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_FLAGS, sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG)
		_ = sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	default:
		// currently, only opengl supported
		// todo: vulkan
		panic(fmt.Errorf("unknown render client API: %s", platform.renderTech))
	}

	// set additional openGL attributes
	_ = sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)
	_ = sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 24)
	_ = sdl.GLSetAttribute(sdl.GL_STENCIL_SIZE, 8)

	// create openGL context
	glContext, err := window.GLCreateContext()
	utils.Check("failed to create OpenGL context", err)

	err = window.GLMakeCurrent(glContext)
	utils.Check("failed to set current OpenGL context", err)

	_ = sdl.GLSetSwapInterval(1)

	// create engine renderer
	surface, err := window.GetSurface()
	utils.Check("window get surface", err)

	err = surface.FillRect(nil, 0)
	utils.Check("clear window surface", err)

	err = window.UpdateSurface()
	utils.Check("update window surface", err)

	renderer, err := window.GetRenderer()
	if renderer == nil {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	}

	utils.Check("create engine renderer", err)
	closer.EnqueueClose(renderer.Destroy)
	platform.renderer = renderer

	return platform, nil
}
