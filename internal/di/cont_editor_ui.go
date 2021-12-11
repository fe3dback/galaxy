package di

import (
	"github.com/fe3dback/galaxy/scope/editor/ui"
)

func (c *Container) createEditorUILayerEntities() *ui.LayerEntities {
	return ui.NewLayerEntities()
}
