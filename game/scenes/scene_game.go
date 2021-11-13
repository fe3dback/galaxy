package scenes

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/movement"
)

type SceneGame struct {
}

func (s SceneGame) CreateEntities() []engine.GameObject {
	testEntity := entity.NewEntity(
		engine.Vec{X: 5, Y: 10},
		engine.Angle0,
	)
	testEntity.AddComponent(movement.NewVelocity(testEntity, engine.VectorForward(0.5), engine.VectorForward(0), engine.Vec{X: 2, Y: 2}))

	return []engine.GameObject{
		testEntity,
	}
}
