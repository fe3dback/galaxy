package player

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

const speed = 300
const speedMultiplier = 3

type Movement struct {
	entity *entity.Entity
}

func NewMovement(entity *entity.Entity) *Movement {
	return &Movement{
		entity: entity,
	}
}

func (r *Movement) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *Movement) OnUpdate(state engine.State) error {
	move := state.Movement().Vector().Mul(speed * state.Moment().DeltaTime())

	if state.Movement().Shift() {
		move = move.Mul(speedMultiplier)
	}

	r.entity.AddPosition(move)

	return nil
}
