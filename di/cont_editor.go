package di

import (
	"github.com/fe3dback/galaxy/editor"
	editorcomponents "github.com/fe3dback/galaxy/editor/components"
	"github.com/fe3dback/galaxy/shared/ui"
)

func (c *Container) ProvideEditorManager() *editor.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	components := make([]editor.Component, 0)
	components = append(components, editorcomponents.NewCamera())
	manager := editor.NewManager(components)

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
