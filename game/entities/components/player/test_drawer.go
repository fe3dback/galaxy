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
	r.DrawCrossLines(engine.ColorCyan, 10, engine.Point{
		X: int(td.entity.Position().X),
		Y: int(td.entity.Position().Y),
	})

	return nil
}

func (td *TestDrawer) OnUpdate(_ engine.State) error {
	return nil
}
