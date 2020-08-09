package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/registry"
)

func gameLoop(provider *registry.Provider) error {
	var err error

	frames := provider.Registry.Game.Frames
	worldManager := provider.Registry.Game.WorldManager
	gameUI := provider.Registry.Game.Ui
	renderer := provider.Registry.Engine.Renderer
	dispatcher := provider.Registry.Engine.Dispatcher
	gameState := provider.Registry.Game.State

	// clear first time screen (fix copy texture from underlying memory)
	renderer.Clear(engine.ColorBackground)
	renderer.Present()

	// render frames
	for frames.Ready() {
		// start frame
		frames.Begin()

		// update
		worldManager.OnFrameStart()

		err = worldManager.CurrentWorld().OnUpdate(gameState)
		if err != nil {
			return fmt.Errorf("can`t update world: %v", err)
		}

		err = gameUI.OnUpdate(gameState)
		if err != nil {
			return fmt.Errorf("can`t update ui: %v", err)
		}

		// draw
		renderer.Clear(engine.ColorBackground)

		renderer.SetRenderMode(engine.RenderModeWorld)
		err = worldManager.CurrentWorld().OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw world: %v", err)
		}

		renderer.SetRenderMode(engine.RenderModeUI)
		err = gameUI.OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw ui: %v", err)
		}

		renderer.Present()

		// debug
		debug(provider)

		// handle events
		dispatcher.HandleQueue()

		// finalize frame
		frames.End()
	}

	return nil
}
