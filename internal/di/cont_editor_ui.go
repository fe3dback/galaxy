package di

import (
	"github.com/fe3dback/galaxy/scope/editor/ui"
)

func (c *Container) createEditorUILayerEntities() *ui.LayerHierarchy {
	if c.memstate.editor.uiLayers.hierarchy != nil {
		return c.memstate.editor.uiLayers.hierarchy
	}

	c.memstate.editor.uiLayers.hierarchy = ui.NewLayerHierarchy()
	return c.memstate.editor.uiLayers.hierarchy
}

func (c *Container) createEditorUILayerSettings() *ui.LayerSettings {
	if c.memstate.editor.uiLayers.settings != nil {
		return c.memstate.editor.uiLayers.settings
	}

	c.memstate.editor.uiLayers.settings = ui.NewLayerSettings()
	return c.memstate.editor.uiLayers.settings
}
