package di

import (
	editor2 "github.com/fe3dback/galaxy/scope/editor"
	editorcomponents "github.com/fe3dback/galaxy/scope/editor/components"
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

func (c *Container) ProvideEditorManager() *editor2.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	components := make([]editor2.Component, 0)
	components = append(components, editorcomponents.NewCamera())
	manager := editor2.NewManager(components)

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
