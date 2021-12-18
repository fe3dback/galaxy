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
	engineGUI := c.ProvideEngineGUI()
	dispatcher := c.ProvideEventDispatcher()
	scenesManager := c.ProvideEngineScenesManager()

	// shared
	gameState := c.ProvideEngineGameState()

	// game
	gameUI := c.ProvideGameUI()

	// editor
	editorManager := c.ProvideEditorManager()

	// clear first time screen (fix copy texture from underlying memory)
	renderer.Clear(defaultColor)
	renderer.EndEngineFrame()

	for frames.Ready() {
		// -----------------------------------
		// start frame
		// -----------------------------------
		frames.Begin()
		engineGUI.StartGUIFrame()
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
		// update game ui
		// -----------------------------------
		if !engineState.InEditorMode() {
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
		if engineState.InEditorMode() {
			err = editorManager.OnBeforeDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor (before): %w", err)
			}
		}

		err = scenesManager.Current().OnDraw(renderer)
		if err != nil {
			return fmt.Errorf("can`t draw game world: %w", err)
		}

		if engineState.InEditorMode() {
			err = editorManager.OnAfterDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw editor (after): %w", err)
			}
		}

		renderer.SetRenderMode(engine.RenderModeUI)
		if !engineState.InEditorMode() {
			err = gameUI.OnDraw(renderer)
			if err != nil {
				return fmt.Errorf("can`t draw game ui: %w", err)
			}
		}

		renderer.EndEngineFrame()
		engineGUI.EndGUIFrame()
		renderer.UpdateGPU()

		// -----------------------------------
		// debug
		// -----------------------------------
		debug(c)

		// -----------------------------------
		// finalize frame
		// -----------------------------------
		dispatcher.PublishEventFrameEnd(event.FrameEndEvent{
			FrameID:   frames.FrameId(),
			DeltaTime: frames.DeltaTime(),
		})
		frames.End()
	}

	return nil
}
