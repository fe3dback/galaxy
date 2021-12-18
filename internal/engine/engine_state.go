package engine

import (
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/event"
)

const (
	modeGame mode = iota
	modeEditor
)

type (
	mode  uint8
	State struct {
		sceneManager galx.SceneManager
		switchQueued bool
		mode         mode
	}
)

func NewEngineState(
	dispatcher *event.Dispatcher,
	sceneManager galx.SceneManager,
	isGameMode bool,
	includeEditor bool,
) *State {
	es := &State{
		sceneManager: sceneManager,
		switchQueued: false,
		mode:         defaultEngineMode(isGameMode, includeEditor),
	}

	if includeEditor {
		dispatcher.OnKeyBoard(es.handleKeyboard)
		dispatcher.OnFrameStart(es.handleFrameStart)
	}

	return es
}

func (as *State) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF1 {
		return nil
	}

	as.switchQueued = true
	return nil
}

func (as *State) handleFrameStart(_ event.FrameStartEvent) error {
	if !as.switchQueued {
		return nil
	}

	as.switchQueued = false
	as.switchState()

	return nil
}

func (as *State) switchState() {
	if as.mode == modeGame {
		// game -> editor
		as.sceneManager.StateToEditorMode()
		as.mode = modeEditor
		return
	}

	// editor -> game
	as.sceneManager.StateToGameMode()
	as.mode = modeGame
}

func (as *State) InEditorMode() bool {
	return as.mode == modeEditor
}

func defaultEngineMode(isGameMode bool, includeEditor bool) mode {
	if !includeEditor {
		return modeGame
	}

	if isGameMode {
		return modeGame
	}

	return modeEditor
}
