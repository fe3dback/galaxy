package factory

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/debug"
	"github.com/fe3dback/galaxy/game/entities/components/game"
	"github.com/fe3dback/galaxy/game/entities/components/player"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/animator"
	"github.com/fe3dback/galaxy/game/entities/components/weapon"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
	"github.com/fe3dback/galaxy/game/utils/physics"
	"github.com/fe3dback/galaxy/generated"
)

func UnitFactoryFn(scheme schemefactory.Unit) entity.FactoryFn {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		// debug
		e.AddComponent(debug.NewGridDrawer(e))

		// anim
		anim := animator.NewAnimator(e)
		anim.AddSequence("idle", animSeqIdle(scheme.TextureRes))
		e.AddComponent(anim)

		// weapon
		e.AddComponent(weapon.NewCharacterInventory(
			e,
			scheme.WeaponsLoader,
			creator.SoundMixer(),
			scheme.EntitySpawner,
		))

		// player
		e.AddComponent(game.NewCameraFollower(e))
		e.AddComponent(player.NewMovement(e, 1.6, 4.2))
		e.AddComponent(game.NewLookToMouse(e))

		// physics
		physShape := creator.Physics().CreateShapeBox(
			32,
			32,
		)

		physBody := creator.Physics().AddBodyDynamic(
			e.Position(),
			e.Rotation(),
			1, // todo
			physShape,
			physics.LayerPlayer.Category(),
			physics.LayerPlayer.Mask(),
		)
		e.AttachPhysicsBody(physBody)

		return e
	}
}

func animSeqIdle(tex generated.ResourcePath) *animator.Sequence {
	return animator.NewSequence(tex, animator.SequenceSlice{
		FrameWidth:  32,
		FrameHeight: 32,
		FirstX:      512,
		FirstY:      736,
		FramesCount: 1,
		SliceType:   animator.SequenceSliceDirectionToBottomToRight,
	})
}
