package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities"
	"github.com/fe3dback/galaxy/generated"
)

func NewLevel01() WorldProviderFn {
	return func(creator engine.WorldCreator) *World {
		world := NewWorld()

		// create one entity process
		carFactory := entities.NewCarFactory(entities.CarFactoryParams{
			TextureRes: generated.ResourcesCarsTaxiTexture,
			PhysicsRes: generated.ResourcesCarsTaxiTaxi,
		})
		carEntity := entity.NewEntity(engine.Vec{X: 500, Y: 400}, engine.Angle45)
		car := carFactory(carEntity, creator)
		world.AddEntity(car)

		// return world
		return world
	}
}
