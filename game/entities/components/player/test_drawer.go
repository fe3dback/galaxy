package player

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type TestDrawer struct {
	entity *entity.Entity
}

func NewTestDrawer(entity *entity.Entity) *TestDrawer {
	return &TestDrawer{
		entity: entity,
	}
}

func (td *TestDrawer) OnDraw(r engine.Renderer) error {
	lx := td.entity.Position().X - 10
	rx := td.entity.Position().X + 10
	ty := td.entity.Position().Y - 10
	by := td.entity.Position().Y + 10

	r.DrawLine(
		engine.ColorCyan,
		engine.Point{X: int(lx), Y: int(ty)},
		engine.Point{X: int(rx), Y: int(by)},
	)

	r.DrawLine(
		engine.ColorCyan,
		engine.Point{X: int(lx), Y: int(by)},
		engine.Point{X: int(rx), Y: int(ty)},
	)

	return nil
}

func (td *TestDrawer) OnUpdate(_ engine.Moment) error {
	return nil
}
