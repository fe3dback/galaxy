package galaxy

import (
	"fmt"

	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/event"
)

const (
	defaultColor = 0x282A36FF
)

func gameLoop(game *Game) error {
	var err error
	c := game.container

	// engine
	engineState := c.ProvideEngineState()
	frames := c.ProvideFrames()
	renderer := c.ProvideEngineRenderer()
	dispatcher := c.ProvideEventDispatcher()
	scenesManager := c.ProvideEngineScenesManager()
	scenesManager.LoadScenes()

	// shared
	gameState := c.ProvideEngineGameState()

	// game
	gameUI := c.ProvideGameUI()

	// editor
	editorManager := c.ProvideEditorManager()
	editorUI := c.ProvideEditorUI()

	// clear first time screen (fix copy texture from underlying memory)
	renderer.Clear(defaultColor)
	renderer.EndEngineFrame()

	for frames.Ready() {
		// -----------------------------------
		// start frame
		// -----------------------------------
		frames.Begin()
		renderer.StartGUIFrame()
		dispatcher.PublishEventFrameStart(event.FrameStartEvent{})

		// -----------------------------------
		// handle events
		// -----------------------------------
		dispatcher.Dispatch()

		// -----------------------------------
		// update state
		// -----------------------------------
		if engineState.InEditorMode() {
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
		if engineState.InEditorMode() {
			err = editorUI.OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update editor ui: %w", err)
			}
		} else {
			err = gameUI.OnUpdate(gameState)
			if err != nil {
				return fmt.Errorf("can`t update game ui: %w", err)
			}
		}

		// -----------------------------------
		// draw
		// -----------------------------------
		renderer.Clear(defaultColor)

		renderer.SetRenderMode(engine.RenderModeWorld)
		err = scenesManager.Current().OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw world: %w", err)
		}

		renderer.SetRenderMode(engine.RenderModeUI)

		if engineState.InEditorMode() {
			err = editorManager.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor: %w", err)
			}

			err = editorUI.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor ui: %w", err)
			}
		} else {
			err = gameUI.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw game ui: %w", err)
			}
		}

		renderer.EndEngineFrame()
		renderer.EndGUIFrame()
		renderer.UpdateGPU()

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
