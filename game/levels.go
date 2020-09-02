package game

import (
	"math/rand"

	"github.com/fe3dback/galaxy/game/gm/physics"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities"
	"github.com/fe3dback/galaxy/game/entities/components/weapon"
	"github.com/fe3dback/galaxy/generated"
)

func NewLevel01() WorldProviderFn {
	return func(creator engine.WorldCreator) *World {
		world := NewWorld(creator.Physics())

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
		player := entity.NewEntity(engine.Vec{}, engine.Angle0)
		player = playerFactory(player, creator)
		physShape := creator.Physics().CreateShapeBox(
			40,
			40,
		)

		physBody := creator.Physics().AddBodyDynamic(
			engine.Vec{X: 500, Y: 400},
			engine.Angle0,
			1,
			physShape,
			physics.LayerPlayer.Category(),
			physics.LayerPlayer.Mask(),
		)
		player.AttachPhysicsBody(physBody)
		world.AddEntity(player)

		// create walls
		xRand := rand.Intn(4) + rand.Intn(16)
		yRand := rand.Intn(4) + rand.Intn(16)

		for xx := 0; xx < xRand; xx++ {
			for yy := 0; yy < yRand; yy++ {
				boxWidth := int32(20 + rand.Intn(100))
				posX := float64(rand.Intn(100) * xx)
				posY := float64(rand.Intn(100) * yy)

				shapeBox := creator.Physics().CreateShapeBox(
					boxWidth,
					boxWidth,
				)

				staticBody := creator.Physics().AddBodyStatic(
					engine.Vec{X: posX, Y: posY},
					engine.Angle0,
					shapeBox,
					physics.LayerGround.Category(),
					physics.LayerGround.Mask(),
				)

				wallFactory := entities.NewStaticFactory(entities.StaticFactoryParams{})

				wall := entity.NewEntity(engine.Vec{}, engine.Angle0)
				wall = wallFactory(wall, creator)
				wall.AttachPhysicsBody(staticBody)
				world.AddEntity(wall)
			}
		}

		// return world
		return world
	}
}
