package factory

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/game"
	"github.com/fe3dback/galaxy/game/entities/components/movement"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/trail"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
)

func BulletFactoryFn(scheme schemefactory.Bullet) entity.FactoryFn {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		e.AddComponent(movement.NewVelocity(
			e,
			scheme.StartAccelerationVec,
			scheme.StartVelocityVec,
			scheme.MaxVelocityVec,
		))

		// add LifeTime component
		e.AddComponent(game.NewLifeTime(e, scheme.LifeTime))

		// add trail
		if scheme.HasTrail {
			e.AddComponent(trail.NewTrail(e, scheme.TrailColor))
		}

		return e
	}
}
