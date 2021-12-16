package gui

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type Hierarchy struct {
	enabled bool
}

func (g Hierarchy) Id() string {
	return "52a64902-baac-4802-978b-8fbeb2832503"
}

func (g *Hierarchy) OnUpdate(s galx.State) error {
	imgui.BeginV("Hierarchy", &g.enabled, 0)

	for _, gameObject := range s.Scene().Entities() {
		g.renderObject(gameObject)
	}

	imgui.End()
	return nil
}

func (g *Hierarchy) renderObject(gameObject galx.GameObject) {
	imgui.PushID(gameObject.Id())

	flags := imgui.TreeNodeFlagsNone

	if gameObject.IsLeaf() {
		flags |= imgui.TreeNodeFlagsLeaf
	}

	if imgui.TreeNodeV(gameObject.Name(), flags) {
		for _, child := range gameObject.Child() {
			g.renderObject(child)
		}

		imgui.TreePop()
	}

	imgui.PopID()
}
