package gui

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type Help struct {
	enabled bool
}

func (g Help) Id() string {
	return "2a80cf67-5775-4a94-9505-f92599661bd5"
}

func (g *Help) OnUpdate(_ galx.State) error {
	imgui.ShowDemoWindow(&g.enabled)
	return nil
}
