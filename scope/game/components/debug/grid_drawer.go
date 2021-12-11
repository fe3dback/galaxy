package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type GridDrawer struct {
	entity galx.GameObject
}

func (td GridDrawer) Id() string {
	return "12173a53-9253-4709-bb3a-469433cab957"
}

func (td GridDrawer) Title() string {
	return "Debug.Grid Drawer"
}

func (td GridDrawer) Description() string {
	return "Draw grid over entity, used for debug purpose"
}

func (td *GridDrawer) OnCreated(entity galx.GameObject) {
	td.entity = entity
}

func (td *GridDrawer) OnDraw(r galx.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	px := td.entity.Position().X
	py := td.entity.Position().Y

	worldX := consts.Meter(int(px/consts.DistanceMeter) * int(consts.DistanceMeter))
	worldY := consts.Meter(int(py/consts.DistanceMeter) * int(consts.DistanceMeter))

	startX := worldX - consts.DistanceMeter*5
	startY := worldY - consts.DistanceMeter*5
	endX := worldX + consts.DistanceMeter*5
	endY := worldY + consts.DistanceMeter*5

	for x := startX; x < endX; x += consts.DistanceMeter {
		for y := startY; y < endY; y += consts.DistanceMeter {

			r.DrawSquare(galx.ColorSelection, galx.RectScreen(
				int(x),
				int(y),
				int(consts.DistanceMeter),
				int(consts.DistanceMeter),
			))

			if r.Gizmos().Debug() {
				r.DrawText(
					consts.AssetDefaultFont,
					galx.ColorSelection,
					fmt.Sprintf("%.0f, %.0f", x/consts.DistanceMeter, y/consts.DistanceMeter),
					galx.Vec{
						X: x + 3,
						Y: y + 3,
					},
				)
			}
		}
	}

	return nil
}

func (td *GridDrawer) OnUpdate(_ galx.State) error {
	return nil
}
