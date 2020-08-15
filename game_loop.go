package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine/event"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/registry"
)

func gameLoop(provider *registry.Provider) error {
	var err error

	// engine
	appState := provider.Registry.Engine.AppState
	frames := provider.Registry.Game.Frames
	renderer := provider.Registry.Engine.Renderer
	dispatcher := provider.Registry.Engine.Dispatcher

	// shared
	worldState := provider.Registry.State

	// game
	worldManager := provider.Registry.Game.WorldManager
	gameUI := provider.Registry.Game.Ui

	// editor
	editorManager := provider.Registry.Editor.Manager
	editorUI := provider.Registry.Editor.Ui

	// clear first time screen (fix copy texture from underlying memory)
	renderer.Clear(engine.ColorBackground)
	renderer.Present()

	for frames.Ready() {
		// -----------------------------------
		// start frame
		// -----------------------------------
		frames.Begin()
		dispatcher.PublishEventFrameStart(event.FrameStartEvent{})

		// -----------------------------------
		// handle events
		// -----------------------------------
		dispatcher.Dispatch()

		// -----------------------------------
		// update
		// -----------------------------------
		worldManager.OnFrameStart()

		if appState.InEditorState() {
			err = editorManager.OnUpdate(worldState)
			if err != nil {
				return fmt.Errorf("can`t update editor: %v", err)
			}
		} else {
			err = worldManager.CurrentWorld().OnUpdate(worldState)
			if err != nil {
				return fmt.Errorf("can`t update game world: %v", err)
			}

			err = gameUI.OnUpdate(worldState)
			if err != nil {
				return fmt.Errorf("can`t update game ui: %v", err)
			}
		}

		// -----------------------------------
		// draw
		// -----------------------------------
		renderer.Clear(engine.ColorBackground)

		renderer.SetRenderMode(engine.RenderModeWorld)
		err = worldManager.CurrentWorld().OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw world: %v", err)
		}

		renderer.SetRenderMode(engine.RenderModeUI)

		if appState.InEditorState() {
			err = editorManager.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor: %v", err)
			}

			err = editorUI.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor ui: %v", err)
			}
		} else {
			err = gameUI.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw game ui: %v", err)
			}
		}

		renderer.Present()

		// -----------------------------------
		// debug
		// -----------------------------------
		debug(provider)

		// -----------------------------------
		// finalize frame
		// -----------------------------------
		dispatcher.PublishEventFrameEnd(event.FrameEndEvent{})
		frames.End()
	}

	return nil
}
