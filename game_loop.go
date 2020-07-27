package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"

	"github.com/veandco/go-sdl2/sdl"
)

type SdlQuitErr struct{}

func (s SdlQuitErr) Error() string {
	panic("sdl quit")
}

func gameLoop(provider *provider) error {
	var err error

	frames := provider.registry.game.frames
	world := provider.registry.game.world
	gameUI := provider.registry.game.ui
	renderer := provider.registry.engine.renderer

	for frames.Ready() {
		// start frame
		frames.Begin()

		// update
		err = world.OnUpdate(frames.DeltaTime())
		if err != nil {
			return fmt.Errorf("can`t update world: %v", err)
		}

		// draw
		renderer.Clear(engine.ColorBlack)

		err = world.OnDraw()
		if err != nil {
			return fmt.Errorf("can`t draw world: %v", err)
		}

		err = gameUI.OnDraw()
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
