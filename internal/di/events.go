package di

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/fe3dback/galaxy/internal/engine/event"
)

type (
	systemPoller struct{}
)

func (p *systemPoller) PollEvents() {
	glfw.PollEvents()
}

func (c *Container) ProvideEventDispatcher() *event.Dispatcher {
	if c.memstate.eventDispatcher != nil {
		return c.memstate.eventDispatcher
	}

	dispatcher := event.NewDispatcher(&systemPoller{})
	dispatcher.OnQuit(c.createEventQuit())

	c.memstate.eventDispatcher = dispatcher
	return c.memstate.eventDispatcher
}

func (c *Container) createEventQuit() event.HandlerQuit {
	return func(quit event.QuitEvent) error {
		log.Print("sdl quit event handled")
		c.ProvideFrames().Interrupt()

		return nil
	}
}
