package registry

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine/lib/event"
	"github.com/fe3dback/galaxy/system"
)

func (r registerFactory) registerDispatcher(onQuit event.HandlerQuit) *event.Dispatcher {
	dispatcher := event.NewEventDispatcher()
	dispatcher.OnQuit(onQuit)

	return dispatcher
}

func (r registerFactory) eventQuit(frames *system.Frames) event.HandlerQuit {
	return func(quit event.EvQuit) error {
		fmt.Printf("sdl quit event handled\n")
		frames.Interrupt()

		return nil
	}
}
