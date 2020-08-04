package engine

import "github.com/fe3dback/galaxy/engine/lib/event"

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
	gz.system = true
	gz.primary = true

	// subscribe to keyboard
	dispatcher.OnKeyBoard(func(keyboard event.EvKeyboard) error {
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

func (d *DrawGizmos) HandleCtrlKey(ev event.EvKeyboard) {
	if ev.Key != event.KeyLctrl {
		return
	}

	if ev.PressType == event.KeyboardPressTypeReleased {
		d.ctrlPressed = false
		return
	}

	d.ctrlPressed = true
}

func (d *DrawGizmos) HandleKeyboard(ev event.EvKeyboard) {
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
