package game

import (
	"github.com/fe3dback/galaxy/game/entities"
	"github.com/fe3dback/galaxy/generated"
)

func NewLevel01() *World {
	world := NewWorld()
	//world.AddEntity(entities.NewGrid())
	//world.AddEntity(entities.NewPlayer())
	world.AddEntity(entities.NewCar(
		generated.ResourcesCarsTaxiTexture,
		generated.ResourcesCarsTaxiTaxi,
	))

	return world
}
