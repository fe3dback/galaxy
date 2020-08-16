package lib

import (
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type SDLLib struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func (s *SDLLib) Window() *sdl.Window {
	return s.window
}

func (s *SDLLib) Renderer() *sdl.Renderer {
	return s.renderer
}

func NewSDLLib(closer *utils.Closer, defaultWidth, defaultHeight int, fullscreen bool) (*SDLLib, error) {
	defer utils.CheckPanic("sdl lib")

	// lib
	err := sdl.Init(sdl.INIT_VIDEO & sdl.INIT_EVENTS)
	utils.Check("init", err)

	err = ttf.Init()
	utils.Check("ttf init", err)

	closer.EnqueueClose(func() error {
		sdl.Quit()
		return nil
	})

	var winFlags uint32
	winFlags &= sdl.WINDOW_SHOWN
	winFlags &= sdl.WINDOW_ALLOW_HIGHDPI
	winFlags &= sdl.WINDOW_BORDERLESS

	winWidth := defaultWidth
	winHeight := defaultHeight

	if fullscreen {
		mode, displayModeErr := sdl.GetCurrentDisplayMode(0)
		utils.Check("get display mode", displayModeErr)

		winFlags &= sdl.WINDOW_FULLSCREEN
		winWidth = int(mode.W)
		winHeight = int(mode.H)
	}

	// window
	window, err := sdl.CreateWindow(
		"Galaxy",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(winWidth), int32(winHeight),
		winFlags,
	)
	utils.Check("create window", err)
	closer.EnqueueClose(window.Destroy)

	surface, err := window.GetSurface()
	utils.Check("window get surface", err)

	err = surface.FillRect(nil, 0)
	utils.Check("clear window surface", err)

	err = window.UpdateSurface()
	utils.Check("update window surface", err)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	utils.Check("create renderer", err)
	closer.EnqueueClose(renderer.Destroy)

	return &SDLLib{
		window:   window,
		renderer: renderer,
	}, nil
}
