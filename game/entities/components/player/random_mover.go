package player

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/animator"
)

type RandomMover struct {
	entity *entity.Entity
}

func NewRandomMover(entity *entity.Entity) *RandomMover {
	// todo component codegen with typecast
	anim := entity.GetComponent(&animator.Animator{}).(*animator.Animator)

	mv := &RandomMover{
		entity: entity,
	}

	time.AfterFunc(time.Second*1, func() {
		anim.PlaySequence("explode")
	})

	//time.AfterFunc(time.Second*10, func() {
	//	anim.PlaySequence("idle")
	//})

	return mv
}

func (r *RandomMover) OnDraw(_ engine.Renderer) error {
	return nil
}

func (r *RandomMover) OnUpdate(s engine.State) error {
	pos := r.entity.Position()
	//pos.X = pos.X + (15 * moment.DeltaTime())
	//pos.Y = pos.Y - (15 * moment.DeltaTime())

	r.entity.SetPosition(pos)
	r.entity.SetRotation(r.entity.Rotation() + engine.Anglef(90*s.Moment().DeltaTime()))

	return nil
}
