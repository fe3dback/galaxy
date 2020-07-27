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
	r.DrawSprite(generated.ResourcesImgGfxAnimTestScheet, -250, -250)

	return nil
}

func (td *TestDrawer) OnUpdate(_ float64) error {
	return nil
}
