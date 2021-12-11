package di

import (
	"github.com/fe3dback/galaxy/scope/shared/ui"
)

func (c *Container) ProvideGameUI() *ui.UI {
	if c.memstate.game.ui != nil {
		return c.memstate.game.ui
	}

	c.memstate.game.ui = ui.NewUI(
		c.createUILayerFPS(),
	)
	return c.memstate.game.ui
}
