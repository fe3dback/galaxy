package registry

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"

	"github.com/fe3dback/galaxy/engine/lib/event"
	"github.com/fe3dback/galaxy/system"
)

func (r registerFactory) registerDispatcher(
	onQuit event.HandlerQuit,
	onEditorSwitch event.HandlerKeyboard,
) *event.Dispatcher {
	dispatcher := event.NewEventDispatcher()
	dispatcher.OnQuit(onQuit)
	dispatcher.OnKeyBoard(onEditorSwitch)

	return dispatcher
}

func (r registerFactory) eventQuit(frames *system.Frames) event.HandlerQuit {
	return func(quit event.EvQuit) error {
		fmt.Printf("sdl quit event handled\n")
		frames.Interrupt()

		return nil
	}
}

func (r registerFactory) eventSwitchEditorState(ed *engine.AppState) event.HandlerKeyboard {
	return func(keyboard event.EvKeyboard) error {
		if keyboard.PressType != event.KeyboardPressTypePressed {
			return nil
		}

		if keyboard.Key != event.KeyF1 {
			return nil
		}

		ed.SwitchState()
		return nil
	}
}
