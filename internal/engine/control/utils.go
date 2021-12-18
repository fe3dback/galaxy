package control

type (
	state = uint8

	engineGUI interface {
		CaptureMouse() bool
		CaptureKeyboard() bool
		CursorOnWindow() bool
	}
)

const (
	stateUnknown state = iota
	statePressed
	stateReleased
	stateDown
	stateUp
)
