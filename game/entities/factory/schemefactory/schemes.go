package schemefactory

import (
	"github.com/fe3dback/galaxy/engine/entity"
)

const (
	// add factory method to world.go - schemeFactoryMap
	// todo: code generation (+world.go schemeFactoryMap generation)
	SchemeGameBullet entity.SchemeID = "game.bullet"
	SchemeGameCar    entity.SchemeID = "game.car"
	SchemeGameUnit   entity.SchemeID = "game.unit"
	SchemeGameWall   entity.SchemeID = "game.wall"
)
