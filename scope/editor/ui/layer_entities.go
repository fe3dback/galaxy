package ui

import (
	"fmt"
	"strconv"

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
	imgui.ShowDemoWindow(&l.open)
	imgui.BeginV("Entities", &l.open, 0)

	for _, gameObject := range s.Scene().Entities() {
		// game object
		imgui.PushID(strconv.FormatInt(gameObject.Id(), 10))

		// properties tree
		if imgui.TreeNode(fmt.Sprintf("ID %d", gameObject.Id())) {
			// pos
			imgui.PushID("pos")
			pos := [2]float32{float32(gameObject.Position().X), float32(gameObject.Position().Y)}
			if imgui.DragFloat2("pos", &pos) {
				gameObject.SetPosition(galx.Vec{
					X: float64(pos[0]),
					Y: float64(pos[1]),
				})
			}
			imgui.PopID()
			// end pos

			imgui.TreePop()
		}

		imgui.PopID()
	}

	imgui.End()
	return nil
}

func (l *LayerEntities) OnDraw(_ galx.Renderer) (err error) {
	return nil
}
