package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/di"
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
)

func gameLoop(c *di.Container) error {
	var err error

	// engine
	appState := c.ProvideEngineAppState()
	frames := c.ProvideFrames()
	renderer := c.ProvideEngineRenderer()
	dispatcher := c.ProvideEventDispatcher()
	scenesManager := c.ProvideEngineScenesManager()

	// shared
	gameState := c.ProvideEngineGameState()

	// game
	gameSceneLoader := c.ProvideGameScenesLoader()
	gameSceneLoader.LoadGameScenes()
	gameSceneLoader.EnterToFirstScene()
	gameUI := c.ProvideGameUI()

	// editor
	editorManager := c.ProvideEditorManager()
	editorUI := c.ProvideEditorUI()

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
		// update state
		// -----------------------------------
		if appState.InEditorState() {
			err = editorManager.OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update editor: %w", err)
			}
		} else {
			err = scenesManager.Current().OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update game scene: %w", err)
			}
		}

		// -----------------------------------
		// update ui
		// -----------------------------------
		if appState.InEditorState() {
			err = editorUI.OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update editor ui: %v", err)
			}
		} else {
			err = gameUI.OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update game ui: %v", err)
			}
		}

		// -----------------------------------
		// draw
		// -----------------------------------
		renderer.Clear(engine.ColorBackground)

		renderer.SetRenderMode(engine.RenderModeWorld)
		err = scenesManager.Current().OnDraw(renderer)
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
		debug(c)

		// -----------------------------------
		// finalize frame
		// -----------------------------------
		dispatcher.PublishEventFrameEnd(event.FrameEndEvent{})
		frames.End()
	}

	return nil
}
