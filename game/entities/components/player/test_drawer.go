package player

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/render"
	"github.com/veandco/go-sdl2/sdl"
)

type TestDrawer struct {
	entity *engine.Entity
}

func NewTestDrawer(entity *engine.Entity) *TestDrawer {
	return &TestDrawer{
		entity: entity,
	}
}

func (td *TestDrawer) OnDraw(r *render.Renderer) error {
	lx := td.entity.Position().X - 10
	rx := td.entity.Position().X + 10
	ty := td.entity.Position().Y - 10
	by := td.entity.Position().Y + 10

	r.DrawLine(
		engine.ColorBlue,
		sdl.Point{X: int32(lx), Y: int32(ty)},
		sdl.Point{X: int32(rx), Y: int32(by)},
	)

	r.DrawLine(
		engine.ColorBlue,
		sdl.Point{X: int32(lx), Y: int32(by)},
		sdl.Point{X: int32(rx), Y: int32(ty)},
	)

	return nil
}

func (td *TestDrawer) OnUpdate(_ engine.Moment) error {
	return nil
}
