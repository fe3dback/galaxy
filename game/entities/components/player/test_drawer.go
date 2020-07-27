package player

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/render"
)

type TestDrawer struct {
	entity *engine.Entity
}

func NewTestDrawer(entity *engine.Entity) *TestDrawer {
	return &TestDrawer{
		entity: entity,
	}
}

func (td *TestDrawer) Id() engine.ComponentId {
	return "player_test_drawer"
}

func (td *TestDrawer) OnDraw(r *render.Renderer) error {
	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			r.DrawSprite(
				generated.ResourcesImgGfxAnimTestScheet,
				int32(int(td.entity.Position().X)-i*300),
				int32(int(td.entity.Position().Y)-j*300),
			)
		}
	}

	return nil
}

func (td *TestDrawer) OnUpdate(_ float64) error {
	return nil
}
