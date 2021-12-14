package di

import (
	"github.com/fe3dback/galaxy/scope/editor"
	"github.com/fe3dback/galaxy/scope/editor/components/control"
	"github.com/fe3dback/galaxy/scope/editor/components/debug"
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

func (c *Container) ProvideEditorManager() *editor.Manager {
	if c.memstate.editor.manager != nil {
		return c.memstate.editor.manager
	}

	// todo auto register all editor components
	componentList := make([]editor.Component, 0)
	componentList = append(componentList, debug.NewGuiHelp())
	componentList = append(componentList, debug.NewGrid(c.createEditorUILayerSettings()))
	componentList = append(componentList, control.NewCamera(c.createEditorUILayerSettings()))
	componentList = append(componentList, control.NewTransform(c.createEditorUILayerSettings())) // control level: 1
	componentList = append(componentList, control.NewSelect(c.provideEngineNodeQuery()))         // control level: 2
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
		c.createEditorUILayerSettings(),
		c.createEditorUILayerEntities(),
	)
	return c.memstate.editor.ui
}
