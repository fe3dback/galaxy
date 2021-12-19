package gui

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type Scene struct {
	enabled bool
}

func (g Scene) Id() string {
	return "67e10e3e-8e89-47a8-8b8e-984a072ed44e"
}

func (g *Scene) OnAfterDraw(r galx.Renderer) error {
	sceneSize := imgui.Vec2{
		X: 320,
		Y: 240,
	}

	imgui.SetNextWindowSize(sceneSize)
	imgui.BeginV("Scene", &g.enabled, 0)

	imgui.End()
	return nil
}
