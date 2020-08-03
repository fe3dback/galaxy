package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/game/entities/components/debug"
	"github.com/fe3dback/galaxy/game/entities/components/game"
)

type AngleOverlay = entity.Entity

func NewAngleOverlay() *AngleOverlay {
	g := entity.NewEntity(
		engine.Vec{X: 0, Y: 0},
		engine.NewAngle(0),
	)
	g.AddComponent(debug.NewAngleOverlay(g))
	g.AddComponent(game.NewCameraFollower(g))

	return g
}
