package event

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/internal/utils"
)

func (d *Dispatcher) pullSDLEvents() {
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
	case *sdl.MouseButtonEvent:
		d.PublishEventMouseButton(assembleMouseButton(sdlEvent))
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

func assembleMouseButton(mouseButtonEvent *sdl.MouseButtonEvent) MouseButtonEvent {
	return MouseButtonEvent{
		IsLeft:     mouseButtonEvent.Button == sdl.BUTTON_LEFT,
		IsRight:    mouseButtonEvent.Button == sdl.BUTTON_RIGHT,
		IsPressed:  mouseButtonEvent.State == sdl.PRESSED,
		IsReleased: mouseButtonEvent.State == sdl.RELEASED,
	}
}
