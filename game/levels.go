package game

import (
	"math/rand"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/factory"
	"github.com/fe3dback/galaxy/game/entities/factory/schemefactory"
	"github.com/fe3dback/galaxy/game/loader/weaponloader"
	"github.com/fe3dback/galaxy/generated"
)

func NewLevel01() WorldProviderFn {
	return func(creator engine.WorldCreator) *World {
		world := NewWorld(creator)

		// create one entity process
		//carFactory := entities.CarFactoryFn(entities.CarParams{
		//	TextureRes: generated.ResourcesCarsTaxiTexture,
		//	PhysicsRes: generated.ResourcesCarsTaxiTaxi,
		//})
		//carEntity := entity.NewEntity(engine.Vec{X: 500, Y: 400}, engine.Angle45)
		//car := carFactory(carEntity, creator)
		//world.AddEntity(car)

		// create player
		playerFactory := factory.UnitFactoryFn(schemefactory.Unit{
			TextureRes:    generated.ResourcesImgCharDefaultCharSheet,
			WeaponsLoader: weaponloader.NewLoader(creator.Loader()),
			EntitySpawner: world,
		})
		player := entity.NewEntity(engine.Vec{X: 500, Y: 400}, engine.Angle0)
		player = playerFactory(player, creator)

		world.AddEntity(player)

		// create walls
		xRand := rand.Intn(4) + rand.Intn(16)
		yRand := rand.Intn(4) + rand.Intn(16)

		for xx := 0; xx < xRand; xx++ {
			for yy := 0; yy < yRand; yy++ {
				boxWidth := int32(20 + rand.Intn(100))
				posX := float64(rand.Intn(100) * xx)
				posY := float64(rand.Intn(100) * yy)

				wallFactory := factory.WallFactoryFn(schemefactory.Wall{
					BoxWidth: boxWidth,
				})

				wall := entity.NewEntity(engine.Vec{X: posX, Y: posY}, engine.Angle0)
				wall = wallFactory(wall, creator)
				world.AddEntity(wall)
			}
		}

		// return world
		return world
	}
}
