package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/registry"
	"github.com/veandco/go-sdl2/sdl"
)

type SdlQuitErr struct{}

func (s SdlQuitErr) Error() string {
	panic("sdl quit")
}

func gameLoop(provider *registry.Provider) error {
	var err error

	frames := provider.Registry.Game.Frames
	world := provider.Registry.Game.World
	gameUI := provider.Registry.Game.Ui
	renderer := provider.Registry.Engine.Renderer

	// clear first time screen (fix copy texture from underlying memory)
	renderer.Clear(engine.ColorBackground)
	renderer.Present()

	// render frames
	for frames.Ready() {
		// start frame
		frames.Begin()

		// update
		err = world.OnUpdate(frames)
		if err != nil {
			return fmt.Errorf("can`t update world: %v", err)
		}

		err = gameUI.OnUpdate(frames)
		if err != nil {
			return fmt.Errorf("can`t update ui: %v", err)
		}

		// draw
		renderer.Clear(engine.ColorBackground)

		err = world.OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw world: %v", err)
		}

		err = gameUI.OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw ui: %v", err)
		}

		renderer.Present()

		// debug
		debug(provider)

		// handle events
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

		// finalize frame
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
