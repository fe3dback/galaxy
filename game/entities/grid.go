package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/debug"
)

type Grid = entity.Entity

func NewGrid() *Grid {
	g := entity.NewEntity(
		engine.Vec{X: 0, Y: 0},
		engine.NewAngle(0),
	)
	g.AddComponent(debug.NewGridDrawer(g))

	return g
}
