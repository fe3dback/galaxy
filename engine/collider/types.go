package collider

import "github.com/fe3dback/galaxy/engine"

type (
	Entity interface {
		Position() engine.Vec
	}
)
