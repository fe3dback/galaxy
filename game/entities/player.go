package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/player"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/animator"
	"github.com/fe3dback/galaxy/generated"
)

type Player = entity.Entity

func NewPlayer() *Player {
	p := entity.NewEntity(
		engine.Vector2D{X: 300, Y: 300},
		engine.Anglef(0),
	)

	// prepare components
	anim := animator.NewAnimator(p)
	anim.AddSequence("idle", animSequenceIdle())
	anim.AddSequence("explode", animSequenceExplode())

	// register components
	p.AddComponent(anim)
	p.AddComponent(player.NewRandomMover(p))
	p.AddComponent(player.NewTestDrawer(p))

	return p
}

func animSequenceIdle() *animator.Sequence {
	texId := generated.ResourcesImgGfxAnimTestSheet2
	slice := animator.SequenceSlice{
		FrameWidth:  32,
		FrameHeight: 32,
		FirstX:      0,
		FirstY:      0,
		FramesCount: 8,
		SliceType:   animator.SequenceSliceDirectionToRightToBottom,
	}

	return animator.NewSequence(texId, slice)
}

func animSequenceExplode() *animator.Sequence {
	texId := generated.ResourcesImgGfxAnimTestScheet
	slice := animator.SequenceSlice{
		FrameWidth:  512,
		FrameHeight: 512,
		FirstX:      0,
		FirstY:      0,
		FramesCount: 64,
		SliceType:   animator.SequenceSliceDirectionToRightToBottom,
	}

	return animator.NewSequence(texId, slice,
		animator.WithFps(64),
		animator.WithCustomPlayback(true, true, animator.SequenceDirectionBackward),
	)
}
