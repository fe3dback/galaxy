package player

import (
	"time"

	"github.com/fe3dback/galaxy/game/entities/components/sprite/animator"
	"github.com/fe3dback/galaxy/render"

	"github.com/fe3dback/galaxy/engine"
)

type RandomMover struct {
	entity *engine.Entity
}

func NewRandomMover(entity *engine.Entity) *RandomMover {
	// todo component codegen with typecast
	anim := entity.GetComponent(&animator.Animator{}).(*animator.Animator)

	mv := &RandomMover{
		entity: entity,
	}

	time.AfterFunc(time.Second*3, func() {
		anim.PlaySequence("explode")
	})

	time.AfterFunc(time.Second*10, func() {
		anim.PlaySequence("idle")
	})

	return mv
}

func (r *RandomMover) OnDraw(_ *render.Renderer) error {
	return nil
}

func (r *RandomMover) OnUpdate(moment engine.Moment) error {
	pos := r.entity.Position()
	//pos.X = pos.X + (15 * moment.DeltaTime())
	//pos.Y = pos.Y - (15 * moment.DeltaTime())

	r.entity.SetPosition(pos)
	r.entity.SetRotation(r.entity.Rotation() + engine.Anglef(90*moment.DeltaTime()))

	return nil
}
