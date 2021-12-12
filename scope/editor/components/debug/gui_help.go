package debug

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type GuiHelp struct {
	enabled bool
}

func NewGuiHelp() *GuiHelp {
	return &GuiHelp{
		enabled: true,
	}
}

func (c *GuiHelp) OnUpdate(_ galx.State) error {
	imgui.ShowDemoWindow(&c.enabled)
	return nil
}

func (c *GuiHelp) OnDraw(_ galx.Renderer) error {
	return nil
}
