package debug

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/units"

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
	px := td.entity.Position().X
	py := td.entity.Position().Y

	worldX := units.Meters(int(px/units.DistanceMeter) * int(units.DistanceMeter))
	worldY := units.Meters(int(py/units.DistanceMeter) * int(units.DistanceMeter))

	startX := worldX - units.DistanceMeter*5
	startY := worldY - units.DistanceMeter*5
	endX := worldX + units.DistanceMeter*5
	endY := worldY + units.DistanceMeter*5

	for x := startX; x < endX; x += units.DistanceMeter {
		for y := startY; y < endY; y += units.DistanceMeter {

			r.DrawPoint(engine.ColorYellow, engine.Vec{
				X: x,
				Y: y,
			})

			r.DrawText(
				generated.ResourcesFontsJetBrainsMonoRegular,
				engine.ColorSelection,
				fmt.Sprintf("%.0f, %.0f", x/units.DistanceMeter, y/units.DistanceMeter),
				engine.Vec{
					X: x + 3,
					Y: y + 3,
				},
			)
		}
	}

	return nil
}

func (td *GridDrawer) OnUpdate(_ engine.State) error {
	return nil
}
