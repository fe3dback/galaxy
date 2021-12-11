package di

import (
	"log"

	event2 "github.com/fe3dback/galaxy/internal/engine/event"
)

func (c *Container) ProvideEventDispatcher() *event2.Dispatcher {
	if c.memstate.eventDispatcher != nil {
		return c.memstate.eventDispatcher
	}

	dispatcher := event2.NewDispatcher()
	dispatcher.OnQuit(c.createEventQuit())

	c.memstate.eventDispatcher = dispatcher
	return c.memstate.eventDispatcher
}

func (c *Container) createEventQuit() event2.HandlerQuit {
	return func(quit event2.QuitEvent) error {
		log.Print("sdl quit event handled")
		c.ProvideFrames().Interrupt()

		return nil
	}
}
