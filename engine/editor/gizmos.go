package editor

import "github.com/fe3dback/galaxy/engine/event"

type DrawGizmos struct {
	system    bool
	primary   bool
	secondary bool
	debug     bool
	spam      bool

	// keyboard state
	ctrlPressed bool
}

func NewDrawGizmos(dispatcher *event.Dispatcher, debugMode bool) *DrawGizmos {
	gz := &DrawGizmos{
		system:    false,
		primary:   false,
		secondary: false,
		debug:     false,
		spam:      false,
	}

	if !debugMode {
		return gz
	}

	// default settings
	gz.system = true // todo move init settings to env
	gz.primary = true
	gz.secondary = true

	// subscribe to keyboard
	dispatcher.OnKeyBoard(func(keyboard event.KeyBoardEvent) error {
		gz.HandleCtrlKey(keyboard)
		gz.HandleKeyboard(keyboard)
		return nil
	})

	return gz
}

func (d *DrawGizmos) System() bool {
	return d.system
}

func (d *DrawGizmos) Primary() bool {
	return d.primary
}

func (d *DrawGizmos) Secondary() bool {
	return d.secondary
}

func (d *DrawGizmos) Debug() bool {
	return d.debug
}

func (d *DrawGizmos) Spam() bool {
	return d.spam
}

func (d *DrawGizmos) HandleCtrlKey(ev event.KeyBoardEvent) {
	if ev.Key != event.KeyLctrl {
		return
	}

	if ev.PressType == event.KeyboardPressTypeReleased {
		d.ctrlPressed = false
		return
	}

	d.ctrlPressed = true
}

func (d *DrawGizmos) HandleKeyboard(ev event.KeyBoardEvent) {
	if !d.ctrlPressed {
		return
	}

	if ev.PressType != event.KeyboardPressTypePressed {
		return
	}

	if ev.Key == event.Key1 {
		d.system = !d.system
	}

	if ev.Key == event.Key2 {
		d.primary = !d.primary
	}

	if ev.Key == event.Key3 {
		d.secondary = !d.secondary
	}

	if ev.Key == event.Key4 {
		d.debug = !d.debug
	}

	if ev.Key == event.Key5 {
		d.spam = !d.spam
	}
}
