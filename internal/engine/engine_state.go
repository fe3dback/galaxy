package engine

import (
	"fmt"

	"github.com/fe3dback/galaxy/internal/engine/event"
)

const (
	modeGame mode = iota
	modeEditor
)

type (
	mode  uint8
	State struct {
		switchQueued bool
		mode         mode
	}
)

func NewEngineState(dispatcher *event.Dispatcher, isGameMode bool) *State {
	es := &State{
		switchQueued: false,
		mode:         defaultEngineMode(isGameMode),
	}

	dispatcher.OnKeyBoard(es.handleKeyboard)
	dispatcher.OnFrameStart(es.handleFrameStart)

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
	switch as.mode {
	case modeEditor:
		as.mode = modeGame
	case modeGame:
		as.mode = modeEditor
	default:
		panic(fmt.Sprintf("unknown app mode `%d`", as.mode))
	}
}

func (as *State) InEditorMode() bool {
	return as.mode == modeEditor
}

func defaultEngineMode(isGameMode bool) mode {
	if isGameMode {
		return modeGame
	}

	return modeEditor
}
