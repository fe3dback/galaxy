package game

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/collider"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities"
	"github.com/fe3dback/galaxy/game/entities/components/weapon"
	"github.com/fe3dback/galaxy/generated"
)

func NewLevel01() WorldProviderFn {
	return func(creator engine.WorldCreator) *World {
		world := NewWorld()

		// basic
		weaponsLoader := weapon.NewLoader(creator)

		// create one entity process
		//carFactory := entities.NewCarFactory(entities.CarFactoryParams{
		//	TextureRes: generated.ResourcesCarsTaxiTexture,
		//	PhysicsRes: generated.ResourcesCarsTaxiTaxi,
		//})
		//carEntity := entity.NewEntity(engine.Vec{X: 500, Y: 400}, engine.Angle45)
		//car := carFactory(carEntity, creator)
		//world.AddEntity(car)

		// create player
		playerFactory := entities.NewUnitFactory(entities.UnitFactoryParams{
			TextureRes:    generated.ResourcesImgCharDefaultCharSheet,
			WeaponsLoader: weaponsLoader,
		})
		player := entity.NewEntity(engine.Vec{X: 500, Y: 400}, engine.Angle45)
		player = playerFactory(player, creator)
		world.AddEntity(player)

		// create walls
		wallFactory := entities.NewStaticFactory(entities.StaticFactoryParams{
			Collider: collider.NewRectCollider(engine.Rect{
				Min: engine.Vec{X: -25, Y: -25},
				Max: engine.Vec{X: 25, Y: 25},
			}),
		})
		wall := entity.NewEntity(engine.Vec{X: 500, Y: 300}, engine.Angle0)
		wall = wallFactory(wall, creator)
		world.AddEntity(wall)

		// return world
		return world
	}
}
