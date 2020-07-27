package player

import (
	"math/rand"
	"time"

	"github.com/fe3dback/galaxy/render"

	"github.com/fe3dback/galaxy/engine"
)

type RandomMover struct {
	entity *engine.Entity
}

func NewRandomMover(entity *engine.Entity) *RandomMover {
	mv := &RandomMover{
		entity: entity,
	}

	time.AfterFunc(time.Second*15, func() {
		mv.entity.Destroy()
	})

	return mv
}

func (r *RandomMover) Id() engine.ComponentId {
	return "player_random_mover"
}

func (r *RandomMover) OnDraw(_ *render.Renderer) error {
	return nil
}

func (r *RandomMover) OnUpdate(deltaTime float64) error {
	pos := r.entity.Position()
	pos.X = pos.X + (5 * deltaTime)
	pos.Y = pos.Y - (5 * deltaTime)

	r.entity.SetPosition(pos)
	r.entity.SetRotation(engine.Anglef(rand.Float64() * 360))

	return nil
}
