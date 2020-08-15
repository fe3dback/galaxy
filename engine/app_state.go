package engine

import "fmt"

const (
	modeGame mode = iota
	modeEditor
)

type (
	mode     uint8
	AppState struct {
		mode mode
	}
)

func NewAppState() *AppState {
	return &AppState{
		mode: modeGame,
	}
}

func (e *AppState) SwitchState() {
	switch e.mode {
	case modeEditor:
		e.mode = modeGame
	case modeGame:
		e.mode = modeEditor
	default:
		panic(fmt.Sprintf("unknown app mode `%d`", e.mode))
	}
}

func (e *AppState) InEditorState() bool {
	return e.mode == modeEditor
}
