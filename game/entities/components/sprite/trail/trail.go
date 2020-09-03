package trail

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/utils"
)

type Trail struct {
	entity           *entity.Entity
	color            engine.Color
	previousPosition engine.Vec
	nextPosition     engine.Vec
}

func NewTrail(entity *entity.Entity, color engine.Color) *Trail {
	return &Trail{
		entity:           entity,
		color:            color,
		previousPosition: entity.Position(),
		nextPosition:     entity.Position(),
	}
}

func (t *Trail) OnDraw(r engine.Renderer) error {
	r.SetRenderTarget(utils.RenderTargetTrails)
	r.DrawLine(t.color, engine.Line{
		A: t.previousPosition,
		B: t.nextPosition,
	})
	r.SetRenderTarget(utils.RenderTargetPrimary)

	return nil
}

func (t *Trail) OnUpdate(s engine.State) error {
	t.previousPosition = t.nextPosition
	t.nextPosition = t.entity.Position()

	return nil
}
