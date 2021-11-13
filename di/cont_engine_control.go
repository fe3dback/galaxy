package di

import "github.com/fe3dback/galaxy/engine/control"

func (c *Container) provideEngineControlMouse() *control.Mouse {
	if c.memstate.control.mouse != nil {
		return c.memstate.control.mouse
	}

	c.memstate.control.mouse = control.NewMouse(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.control.mouse
}

func (c *Container) provideEngineControlMovement() *control.Movement {
	if c.memstate.control.movement != nil {
		return c.memstate.control.movement
	}

	c.memstate.control.movement = control.NewMovement(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.control.movement
}
