package di

import (
	control2 "github.com/fe3dback/galaxy/internal/engine/control"
)

func (c *Container) provideEngineControlMouse() *control2.Mouse {
	if c.memstate.control.mouse != nil {
		return c.memstate.control.mouse
	}

	c.memstate.control.mouse = control2.NewMouse(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.control.mouse
}

func (c *Container) provideEngineControlMovement() *control2.Movement {
	if c.memstate.control.movement != nil {
		return c.memstate.control.movement
	}

	c.memstate.control.movement = control2.NewMovement(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.control.movement
}
