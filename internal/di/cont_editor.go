package di

import (
	"github.com/fe3dback/galaxy/scope/editor"
	"github.com/fe3dback/galaxy/scope/editor/components"
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

func (c *Container) ProvideEditorManager() *editor.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	componentList := make([]editor.Component, 0)
	componentList = append(componentList, components.NewCamera())
	componentList = append(componentList, components.NewGrid())
	componentList = append(componentList, components.NewSelect(c.provideEngineNodeQuery()))
	manager := editor.NewManager(componentList)

	c.memstate.editor.manager = manager
	return c.memstate.editor.manager
}

func (c *Container) ProvideEditorUI() *ui.UI {
	if c.memstate.editor.ui != nil {
		return c.memstate.editor.ui
	}

	c.memstate.editor.ui = ui.NewUI(
		c.createUILayerFPS(),
		c.createEditorUILayerEntities(),
	)
	return c.memstate.editor.ui
}
