package event

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	EvQuit struct {
	}

	// todo: codegen
	HandlerQuit func(quit EvQuit) error
)

// todo: codegen
func (d *Dispatcher) OnQuit(h HandlerQuit) {
	d.registryHandler(typeQuit, func(e sdl.Event) error {
		evQuit, ok := e.(*sdl.QuitEvent)
		if !ok {
			panic(fmt.Sprintf("can`t handle `OnQuit` unexpected event type `%s`", reflect.TypeOf(e)))
		}

		return h(d.assembleQuit(evQuit))
	})
}

func (d *Dispatcher) assembleQuit(_ *sdl.QuitEvent) EvQuit {
	return EvQuit{}
}
