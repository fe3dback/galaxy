package factory

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/game"
	"github.com/fe3dback/galaxy/game/entities/components/movement"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/trail"
)

type BulletParams struct {
	StartAccelerationVec engine.Vec
	StartVelocityVec     engine.Vec
	MaxVelocityVec       engine.Vec
	LifeTime             time.Duration
	HasTrail             bool
	TrailColor           engine.Color
}

func BulletFactoryFn(params BulletParams) entity.FactoryFn {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		e.AddComponent(movement.NewVelocity(
			e,
			params.StartAccelerationVec,
			params.StartVelocityVec,
			params.MaxVelocityVec,
		))

		// add LifeTime component
		e.AddComponent(game.NewLifeTime(e, params.LifeTime))

		// add trail
		if params.HasTrail {
			e.AddComponent(trail.NewTrail(e, params.TrailColor))
		}

		return e
	}
}
