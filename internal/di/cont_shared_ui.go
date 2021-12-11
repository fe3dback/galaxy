package di

import (
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

func (c *Container) createUILayerFPS() *ui.LayerFPS {
	return ui.NewLayerSharedFPS()
}
