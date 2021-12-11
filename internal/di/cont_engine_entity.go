package di

import (
	"github.com/fe3dback/galaxy/internal/engine/entity"
)

func (c *Container) provideEngineEntityComponentsRegistry() *entity.ComponentsRegistry {
	if c.memstate.entity.componentsRegistry != nil {
		return c.memstate.entity.componentsRegistry
	}

	reg := entity.NewComponentsRegistry()
	for _, component := range c.flags.Components() {
		reg.RegisterComponent(component)
	}

	c.memstate.entity.componentsRegistry = reg
	return c.memstate.entity.componentsRegistry
}
