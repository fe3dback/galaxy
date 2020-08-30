package game

import (
	"math/rand"

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
		xRand := rand.Intn(4) + rand.Intn(16)
		yRand := rand.Intn(4) + rand.Intn(16)

		for xx := 0; xx < xRand; xx++ {
			for yy := 0; yy < yRand; yy++ {
				boxWidth := float64(10 + rand.Intn(50))

				wallFactory := entities.NewStaticFactory(entities.StaticFactoryParams{
					Collider: collider.NewRectCollider(engine.Rect{
						Min: engine.Vec{X: -boxWidth, Y: -boxWidth},
						Max: engine.Vec{X: boxWidth, Y: boxWidth},
					}),
				})

				posX := float64(rand.Intn(100) * xx)
				posY := float64(rand.Intn(100) * yy)

				wall := entity.NewEntity(engine.Vec{X: posX, Y: posY}, engine.Angle0)
				wall = wallFactory(wall, creator)
				world.AddEntity(wall)
			}
		}

		// return world
		return world
	}
}
