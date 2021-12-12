package ui

import (
	"github.com/inkyblackness/imgui-go/v4"

	"github.com/fe3dback/galaxy/galx"
)

type (
	LayerEntities struct {
		open bool
	}
)

func NewLayerEntities() *LayerEntities {
	return &LayerEntities{
		open: true,
	}
}

func (l *LayerEntities) OnUpdate(s galx.State) error {
	imgui.BeginV("Hierarchy", &l.open, 0)

	for _, gameObject := range s.Scene().Entities() {
		l.renderObject(gameObject)
	}

	imgui.End()
	return nil
}

func (l *LayerEntities) renderObject(gameObject galx.GameObject) {
	imgui.PushID(gameObject.Id())

	flags := imgui.TreeNodeFlagsNone

	if gameObject.IsLeaf() {
		flags |= imgui.TreeNodeFlagsLeaf
	}

	if imgui.TreeNodeV(gameObject.Name(), flags) {
		for _, child := range gameObject.Child() {
			l.renderObject(child)
		}

		imgui.TreePop()
	}

	imgui.PopID()
}

func (l *LayerEntities) OnDraw(_ galx.Renderer) (err error) {
	return nil
}
