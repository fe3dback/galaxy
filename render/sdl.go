package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLLib struct {
	window *sdl.Window
}

func (s *SDLLib) Close() error {
	err := s.window.Destroy()
	sdl.Quit()

	return err
}

func (s *SDLLib) Window() *sdl.Window {
	return s.window
}

func CreateSDL() (lib *SDLLib, sdlError error) {
	defer func() {
		if err := recover(); err != nil {
			sdlError = fmt.Errorf("sdl: %v", err)
			return
		}
	}()

	// lib
	err := sdl.Init(sdl.INIT_EVERYTHING)
	check(err, "init")

	// window
	window, err := sdl.CreateWindow(
		"Galaxy",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	check(err, "create window")

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
