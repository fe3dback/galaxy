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
	px := td.entity.Position().X
	py := td.entity.Position().Y

	r.DrawCrossLines(engine.ColorSelection, 15, engine.Point{
		X: int(px),
		Y: int(py),
	})

	r.DrawSquare(engine.ColorPurple, engine.Rect{
		X: int(r.Camera().Position().X),
		Y: int(r.Camera().Position().Y),
		W: r.Camera().Width() - 1,
		H: r.Camera().Height() - 1,
	})

	for x := px - 1024; x < px+1024; x += 128 {
		for y := py - 1024; y < py+1024; y += 128 {
			r.DrawPoint(engine.ColorYellow, engine.Point{
				X: int(x),
				Y: int(y),
			})

			r.DrawText(
				generated.ResourcesFontsJetBrainsMonoRegular,
				engine.ColorSelection,
				fmt.Sprintf("%d, %d", int(x), int(y)),
				engine.Point{
					X: int(x + 3),
					Y: int(y + 3),
				},
			)
		}
	}

	return nil
}

func (td *GridDrawer) OnUpdate(_ engine.State) error {
	return nil
}
