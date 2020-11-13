package factory

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/car"
	"github.com/fe3dback/galaxy/game/entities/components/debug"
	"github.com/fe3dback/galaxy/game/entities/components/game"
	"github.com/fe3dback/galaxy/game/entities/components/sprite/sprite2d"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
)

func CarFactoryFn(scheme schemefactory.Car) entity.FactoryFn {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		// common
		e.AddComponent(debug.NewGridDrawer(e))
		e.AddComponent(sprite2d.NewSprite2D(e, scheme.TextureRes))

		// physics
		physicsSpec := car.YamlSpec{}
		creator.Loader().LoadYaml(scheme.PhysicsRes, &physicsSpec)
		e.AddComponent(car.NewPhysics(e, physicsSpec))

		// camera
		e.AddComponent(game.NewCameraFollower(e))

		return e
	}
}
