package player

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

const speed = 10
const speedMultiplier = 5

type Movement struct {
	entity *entity.Entity
	vec    engine.Vector2D
}

func NewMovement(entity *entity.Entity) *Movement {
	return &Movement{
		entity: entity,
	}
}

func (r *Movement) OnDraw(d engine.Renderer) error {
	text := fmt.Sprintf("%.4f, %.4f", r.vec.X, r.vec.Y)
	d.DrawText(generated.ResourcesFontsJetBrainsMonoRegular, engine.ColorCyan, text, engine.Point{
		X: r.entity.Position().RoundX() + 50,
		Y: r.entity.Position().RoundY() + 50,
	})

	return nil
}

func (r *Movement) OnUpdate(state engine.State) error {
	move := state.Movement().Vector().Mul(speed * state.Moment().DeltaTime())
	r.vec = state.Movement().Vector()

	if state.Movement().Shift() {
		move = move.Mul(speedMultiplier)
	}

	r.entity.AddPosition(move)

	return nil
}
