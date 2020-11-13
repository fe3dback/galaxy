package entity

import (
	"github.com/fe3dback/galaxy/engine"
)

type (
	SchemeID = string

	Scheme interface {
		SchemeID() SchemeID
	}

	Spawner interface {
		SpawnEntity(pos engine.Vec, angle engine.Angle, scheme Scheme)
	}
)
