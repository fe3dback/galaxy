package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/entities/components/player"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/animator"
	"github.com/fe3dback/galaxy/generated"
)

type Player = engine.Entity

func NewPlayer() *Player {
	p := engine.NewEntity(
		engine.Vector2D{X: 0, Y: 0},
		engine.Anglef(0),
	)
	p.AddComponent(player.NewRandomMover(p))

	anim := animator.NewAnimator(p)
	anim.AddSequence("idle", animSequenceIdle())

	p.AddComponent(anim)

	return p
}

func animSequenceIdle() *animator.Sequence {
	texId := generated.ResourcesImgGfxAnimTestScheet
	slice := animator.SequenceSlice{
		FrameWidth:  512,
		FrameHeight: 512,
		FirstX:      0,
		FirstY:      0,
		FramesCount: 64,
		SliceType:   animator.SequenceSliceDirectionToRightToBottom,
	}

	return animator.NewSequence(texId, slice, animator.WithFps(5))
}
