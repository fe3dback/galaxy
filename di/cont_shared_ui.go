package di

import "github.com/fe3dback/galaxy/shared/ui"

func (c *Container) createUILayerFPS() *ui.LayerFPS {
	return ui.NewLayerSharedFPS()
}
