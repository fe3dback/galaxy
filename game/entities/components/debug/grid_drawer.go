package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/gm"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type GridDrawer struct {
	entity *entity.Entity
}

func NewGridDrawer(entity *entity.Entity) *GridDrawer {
	comp := &GridDrawer{
		entity: entity,
	}

	return comp
}

func (td *GridDrawer) OnDraw(r engine.Renderer) error {
	if !r.Gizmos().Secondary() {
		return nil
	}

	px := td.entity.Position().X
	py := td.entity.Position().Y

	worldX := gm.Meter(int(px/gm.DistanceMeter) * int(gm.DistanceMeter))
	worldY := gm.Meter(int(py/gm.DistanceMeter) * int(gm.DistanceMeter))

	startX := worldX - gm.DistanceMeter*5
	startY := worldY - gm.DistanceMeter*5
	endX := worldX + gm.DistanceMeter*5
	endY := worldY + gm.DistanceMeter*5

	for x := startX; x < endX; x += gm.DistanceMeter {
		for y := startY; y < endY; y += gm.DistanceMeter {

			r.DrawSquare(engine.ColorSelection, engine.RectScreen(
				int(x),
				int(y),
				int(gm.DistanceMeter),
				int(gm.DistanceMeter),
			))

			if r.Gizmos().Debug() {
				r.DrawText(
					generated.ResourcesFontsJetBrainsMonoRegular,
					engine.ColorSelection,
					fmt.Sprintf("%.0f, %.0f", x/gm.DistanceMeter, y/gm.DistanceMeter),
					engine.Vec{
						X: x + 3,
						Y: y + 3,
					},
				)
			}
		}
	}

	return nil
}

func (td *GridDrawer) OnUpdate(_ engine.State) error {
	return nil
}
