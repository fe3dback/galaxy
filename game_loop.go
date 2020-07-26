package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type SdlQuitErr struct{}

func (s SdlQuitErr) Error() string {
	panic("sdl quit")
}

func gameLoop(params *gameParams) error {
	var err error

	frames := params.frames
	world := params.world
	ui := params.ui

	for frames.Ready() {
		frames.Begin()

		// -- game loop

		err = world.OnUpdate(frames.DeltaTime())
		if err != nil {
			return fmt.Errorf("can`t update world: %v", err)
		}

		err = world.OnDraw()
		if err != nil {
			return fmt.Errorf("can`t draw world: %v", err)
		}

		err = ui.OnDraw()
		if err != nil {
			return fmt.Errorf("can`t draw ui: %v", err)
		}

		debug(params)

		// -- handle events for next frame

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			err := handleSdlEvent(event)

			if err != nil {
				switch err.(type) {
				case SdlQuitErr:
					frames.Interrupt()
				default:
					return fmt.Errorf("can`t process sdl event: %v", err)
				}
			}
		}

		frames.End()
	}

	return nil
}

func handleSdlEvent(ev sdl.Event) error {
	switch ev.(type) {
	case *sdl.QuitEvent:
		fmt.Printf("sdl quit event handled\n")
		return SdlQuitErr{}
	}

	return nil
}
