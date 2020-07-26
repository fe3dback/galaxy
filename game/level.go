package game

import "github.com/fe3dback/galaxy/game/entities"

type Level struct {
	world *World
}

func (l *Level) OnUpdate(deltaTime float64) error {
	return l.world.OnUpdate(deltaTime)
}

func (l *Level) OnDraw() error {
	return l.world.OnDraw()
}

func NewBasicLevel() *Level {
	world := NewWorld()
	world.AddEntity(entities.NewPlayer())

	return &Level{
		world: world,
	}
}
