package di

import (
	"github.com/fe3dback/galaxy/internal/engine/control"
)

func (c *Container) provideEngineControlMouse() *control.Mouse {
	if c.memstate.control.mouse != nil {
		return c.memstate.control.mouse
	}

	c.memstate.control.mouse = control.NewMouse(
		c.ProvideEventDispatcher(),
		c.provideSDL().GUI(),
	)
	return c.memstate.control.mouse
}

func (c *Container) provideEngineControlKeyboard() *control.Keyboard {
	if c.memstate.control.keyboard != nil {
		return c.memstate.control.keyboard
	}

	c.memstate.control.keyboard = control.NewKeyboard(
		c.ProvideEventDispatcher(),
	)
	return c.memstate.control.keyboard
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
