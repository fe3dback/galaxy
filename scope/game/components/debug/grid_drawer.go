package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
)

type GridDrawer struct {
	entity galx.GameObject
}

func NewGridDrawer(entity galx.GameObject) *GridDrawer {
	comp := &GridDrawer{
		entity: entity,
	}

	return comp
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
					consts.DefaultFont,
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
