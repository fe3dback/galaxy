package event

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowEventTypeResize      WindowEventType = sdl.WINDOWEVENT_RESIZED
	WindowEventTypeSizeChanged WindowEventType = sdl.WINDOWEVENT_SIZE_CHANGED
)

type (
	WindowEventType = uint8

	EvWindow struct {
		WindowID  uint32
		EventType WindowEventType
		Data1     int32
		Data2     int32
	}

	// todo: codegen
	HandlerWindow func(window EvWindow) error
)

// todo: codegen
func (d *Dispatcher) OnWindow(h HandlerWindow) {
	d.registryHandler(typeWindow, func(e sdl.Event) error {
		evWindow, ok := e.(*sdl.WindowEvent)
		if !ok {
			panic(fmt.Sprintf("can`t handle `OnWindow` unexpected event type `%s`", reflect.TypeOf(e)))
		}

		return h(d.assembleWindow(evWindow))
	})
}

func (d *Dispatcher) assembleWindow(ev *sdl.WindowEvent) EvWindow {
	return EvWindow{
		WindowID:  ev.WindowID,
		EventType: ev.Event,
		Data1:     ev.Data1,
		Data2:     ev.Data2,
	}
}
