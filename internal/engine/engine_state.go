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
	mode        uint8
	EngineState struct {
		switchQueued bool
		mode         mode
	}
)

func NewEngineState(dispatcher *event.Dispatcher) *EngineState {
	as := &EngineState{
		switchQueued: false,
		mode:         modeGame,
	}

	dispatcher.OnKeyBoard(as.handleKeyboard)
	dispatcher.OnFrameStart(as.handleFrameStart)

	return as
}

func (as *EngineState) handleKeyboard(keyboard event.KeyBoardEvent) error {
	if keyboard.PressType != event.KeyboardPressTypePressed {
		return nil
	}

	if keyboard.Key != event.KeyF1 {
		return nil
	}

	as.switchQueued = true
	return nil
}

func (as *EngineState) handleFrameStart(_ event.FrameStartEvent) error {
	if !as.switchQueued {
		return nil
	}

	as.switchQueued = false
	as.switchState()

	return nil
}

func (as *EngineState) switchState() {
	switch as.mode {
	case modeEditor:
		as.mode = modeGame
	case modeGame:
		as.mode = modeEditor
	default:
		panic(fmt.Sprintf("unknown app mode `%d`", as.mode))
	}
}

func (as *EngineState) InEditorMode() bool {
	return as.mode == modeEditor
}
