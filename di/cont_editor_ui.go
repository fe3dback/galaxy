package di

import (
	"github.com/fe3dback/galaxy/editor/ui"
)

func (c *Container) createEditorUILayerEntities() *ui.LayerEntities {
	return ui.NewLayerEntities()
}
