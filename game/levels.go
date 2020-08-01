package game

import "github.com/fe3dback/galaxy/game/entities"

func NewLevel01() *World {
	world := NewWorld()
	world.AddEntity(entities.NewGrid())
	world.AddEntity(entities.NewPlayer())

	return world
}
