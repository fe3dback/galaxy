package entities

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/car"
	"github.com/fe3dback/galaxy/game/entities/components/debug"
	"github.com/fe3dback/galaxy/game/entities/components/game"
	"github.com/fe3dback/galaxy/generated"
)

type Car = entity.Entity

func NewCar(texture, physics generated.ResourcePath) *Car {
	e := entity.NewEntity(
		engine.Vec{X: 300, Y: 250},
		engine.NewAngle(80),
	)

	go func() {
		tick := time.Tick(time.Millisecond * 100)
		for range tick {
			e.AddRotation(engine.NewAngle(1))
		}
	}()

	//e.AddComponent(sprite2d.NewSprite2D(e, texture))
	e.AddComponent(car.NewPhysics(e, physics))
	e.AddComponent(game.NewCameraFollower(e))
	e.AddComponent(debug.NewGridDrawer(e))

	return e
}
