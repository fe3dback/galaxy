package event

import (
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (d *Dispatcher) PullSDLEvents() {
	defer utils.CheckPanic("handle events")

	for {
		event := sdl.PollEvent()
		if event == nil {
			return
		}

		d.dispatchSDLEvent(event)
	}
}

func (d *Dispatcher) dispatchSDLEvent(ev sdl.Event) {
	switch sdlEvent := ev.(type) {
	case *sdl.QuitEvent:
		d.PublishEventQuit(assembleQuit(sdlEvent))
	case *sdl.KeyboardEvent:
		d.PublishEventKeyBoard(assembleKeyboard(sdlEvent))
	case *sdl.WindowEvent:
		d.PublishEventWindow(assembleWindow(sdlEvent))
	case *sdl.MouseWheelEvent:
		d.PublishEventMouseWheel(assembleMouseWheel(sdlEvent))
	}
}

func assembleQuit(_ *sdl.QuitEvent) QuitEvent {
	return QuitEvent{}
}

func assembleKeyboard(ev *sdl.KeyboardEvent) KeyBoardEvent {
	var pressType KeyboardPressType

	if ev.Type == sdl.KEYDOWN {
		pressType = KeyboardPressTypePressed
	} else {
		pressType = KeyboardPressTypeReleased
	}

	return KeyBoardEvent{
		PressType: pressType,
		Key:       ev.Keysym.Sym,
	}
}

func assembleMouseWheel(mouseWheelEvent *sdl.MouseWheelEvent) MouseWheelEvent {
	return MouseWheelEvent{
		ScrollOffset: float64(mouseWheelEvent.Y),
	}
}

func assembleWindow(ev *sdl.WindowEvent) WindowEvent {
	return WindowEvent{
		WindowID:  ev.WindowID,
		EventType: ev.Event,
		Data1:     ev.Data1,
		Data2:     ev.Data2,
	}
}
