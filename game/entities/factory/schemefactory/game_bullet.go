package schemefactory

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type Bullet struct {
	StartAccelerationVec engine.Vec
	StartVelocityVec     engine.Vec
	MaxVelocityVec       engine.Vec
	LifeTime             time.Duration
	HasTrail             bool
	TrailColor           engine.Color
}

func (b Bullet) SchemeID() entity.SchemeID {
	return SchemeGameBullet
}
