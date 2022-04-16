package di

import (
	"github.com/fe3dback/galaxy/internal/engine/node"
)

func (c *Container) provideEngineNodeComponentsRegistry() *node.ComponentsRegistry {
	if c.memstate.node.componentsRegistry != nil {
		return c.memstate.node.componentsRegistry
	}

	reg := node.NewComponentsRegistry()
	for _, component := range c.flags.Components() {
		reg.RegisterComponent(component)
	}

	c.memstate.node.componentsRegistry = reg
	return c.memstate.node.componentsRegistry
}

func (c *Container) provideEngineNodeQuery() *node.ObjectQuery {
	if c.memstate.node.query != nil {
		return c.memstate.node.query
	}

	c.memstate.node.query = node.NewObjectQuery(
		c.ProvideEngineScenesManager(),
		c.provideRenderCameraOLD(),
	)
	return c.memstate.node.query
}
