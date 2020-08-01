package event

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	maxHandlers = 1024

	typeQuit     handlerType = "quit"
	typeKeyboard handlerType = "Keyboard"
	typeWindow   handlerType = "Window"
)

type (
	handlerType string
	handlerFunc func(event sdl.Event) error

	Dispatcher struct {
		handlers map[handlerType][]handlerFunc
	}
)

func NewEventDispatcher() *Dispatcher {
	d := &Dispatcher{
		handlers: map[handlerType][]handlerFunc{},
	}

	sdl.SetEventFilter(d, struct{}{})

	return d
}

func (d *Dispatcher) FilterEvent(e sdl.Event, _ interface{}) bool {
	switch e.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.KeyboardEvent:
		return true
	case *sdl.MouseMotionEvent:
		return true
	case *sdl.WindowEvent:
		return true
	}

	return false
}

func (d *Dispatcher) HandleQueue() {
	defer utils.CheckPanic("handle events")

	for {
		event := sdl.PollEvent()
		if event == nil {
			return
		}

		dispatchErr := d.dispatch(event)
		utils.Check(fmt.Sprintf("process sdl event `%s`", reflect.TypeOf(event)), dispatchErr)
	}
}

func (d *Dispatcher) registryHandler(t handlerType, h handlerFunc) {
	count := len(d.handlers[t])
	if count >= maxHandlers {
		panic(fmt.Sprintf("Can`t register more than %d `%s` handlers", maxHandlers, t))
	}

	fmt.Printf("registered `%s` handler\n", t)
	d.handlers[t] = append(d.handlers[t], h)
}

// todo: codegen
func (d *Dispatcher) dispatch(ev sdl.Event) error {
	switch ev.(type) {
	case *sdl.QuitEvent:
		return d.send(typeQuit, ev)
	case *sdl.KeyboardEvent:
		return d.send(typeKeyboard, ev)
	case *sdl.WindowEvent:
		return d.send(typeWindow, ev)
	}

	return nil
}

func (d *Dispatcher) send(t handlerType, ev sdl.Event) error {
	handlers, ok := d.handlers[t]
	if !ok {
		return nil
	}

	for _, handler := range handlers {
		err := handler(ev)
		utils.Check("handle", err)
	}

	return nil
}
