package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/game/entities/components/player"
)

type Player = engine.Entity

func NewPlayer() *Player {
	p := engine.NewEntity(
		engine.Vector2D{X: 0, Y: 0},
		engine.Anglef(0),
	)
	p.AddComponent(player.NewRandomMover(p))
	p.AddComponent(player.NewTestDrawer(p))

	return p
}
