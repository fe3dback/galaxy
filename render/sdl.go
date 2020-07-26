package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/fe3dback/galaxy/utils"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLLib struct {
	window *sdl.Window
}

func (s *SDLLib) Window() *sdl.Window {
	return s.window
}

func NewSDLLib(closer *utils.Closer) (lib *SDLLib, sdlError error) {
	defer func() {
		if err := recover(); err != nil {
			sdlError = fmt.Errorf("sdl: %v", err)
			return
		}
	}()

	// lib
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err, "init")

	err = ttf.Init()
	check(err, "ttf init")

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
	check(err, "create window")
	closer.Enqueue(window.Destroy)

	surface, err := window.GetSurface()
	check(err, "window get surface")

	err = surface.FillRect(nil, 0)
	check(err, "clear window surface")

	err = window.UpdateSurface()
	check(err, "update window surface")

	return &SDLLib{
		window: window,
	}, nil
}

func check(err error, explain string) {
	if err == nil {
		return
	}

	panic(fmt.Sprintf("%s: %v", explain, err))
}
