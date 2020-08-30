package debug

import (
	"fmt"

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

	worldX := engine.Meter(int(px/engine.DistanceMeter) * int(engine.DistanceMeter))
	worldY := engine.Meter(int(py/engine.DistanceMeter) * int(engine.DistanceMeter))

	startX := worldX - engine.DistanceMeter*5
	startY := worldY - engine.DistanceMeter*5
	endX := worldX + engine.DistanceMeter*5
	endY := worldY + engine.DistanceMeter*5

	for x := startX; x < endX; x += engine.DistanceMeter {
		for y := startY; y < endY; y += engine.DistanceMeter {

			r.DrawSquare(engine.ColorSelection, engine.RectScreen(
				int(x),
				int(y),
				int(engine.DistanceMeter),
				int(engine.DistanceMeter),
			))

			if r.Gizmos().Debug() {
				r.DrawText(
					generated.ResourcesFontsJetBrainsMonoRegular,
					engine.ColorSelection,
					fmt.Sprintf("%.0f, %.0f", x/engine.DistanceMeter, y/engine.DistanceMeter),
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
