package engine

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine/event"
)

const (
	modeGame mode = iota
	modeEditor
)

type (
	mode     uint8
	AppState struct {
		switchQueued bool
		mode         mode
	}
)

func NewAppState(dispatcher *event.Dispatcher) *AppState {
	as := &AppState{
		switchQueued: false,
		mode:         modeGame,
	}

	dispatcher.OnKeyBoard(as.handleKeyboard)
	dispatcher.OnFrameStart(as.handleFrameStart)

	return as
}

func (as *AppState) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF1 {
		return nil
	}

	as.switchQueued = true
	return nil
}

func (as *AppState) handleFrameStart(_ event.FrameStartEvent) error {
	if !as.switchQueued {
		return nil
	}

	as.switchQueued = false
	as.switchState()

	return nil
}

func (as *AppState) switchState() {
	switch as.mode {
	case modeEditor:
		as.mode = modeGame
	case modeGame:
		as.mode = modeEditor
	default:
		panic(fmt.Sprintf("unknown app mode `%d`", as.mode))
	}
}

func (as *AppState) InEditorState() bool {
	return as.mode == modeEditor
}
