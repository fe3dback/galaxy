package registry

import (
	"log"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/fe3dback/galaxy/system"
)

func (r registerFactory) registerDispatcher(
	onQuit event.HandlerQuit,
	onEditorSwitch event.HandlerKeyBoard,
) *event.Dispatcher {
	dispatcher := event.NewDispatcher()
	dispatcher.OnQuit(onQuit)
	dispatcher.OnKeyBoard(onEditorSwitch)

	return dispatcher
}

func (r registerFactory) eventQuit(frames *system.Frames) event.HandlerQuit {
	return func(quit event.QuitEvent) error {
		log.Print("sdl quit event handled")
		frames.Interrupt()

		return nil
	}
}

func (r registerFactory) eventSwitchEditorState(ed *engine.AppState) event.HandlerKeyBoard {
	return func(keyboard event.KeyBoardEvent) error {
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
