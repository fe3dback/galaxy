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

func NewSDLLib(closer *utils.Closer) (*SDLLib, error) {
	defer utils.CheckPanic("sdl lib")

	// lib
	err := sdl.Init(sdl.INIT_EVERYTHING)
	utils.Check("init", err)

	err = ttf.Init()
	utils.Check("ttf init", err)

	closer.Enqueue(func() error {
		sdl.Quit()
		return nil
	})

	// window
	window, err := sdl.CreateWindow(
		"Galaxy",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	utils.Check("create window", err)
	closer.Enqueue(window.Destroy)

	surface, err := window.GetSurface()
	utils.Check("window get surface", err)

	err = surface.FillRect(nil, 0)
	utils.Check("clear window surface", err)

	err = window.UpdateSurface()
	utils.Check("update window surface", err)

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	utils.Check("create renderer", err)
	closer.Enqueue(renderer.Destroy)

	return &SDLLib{
		window:   window,
		renderer: renderer,
	}, nil
}
