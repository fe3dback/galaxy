package entity

import "github.com/fe3dback/galaxy/engine"

type Spawner interface {
	SpawnEntity(pos engine.Vec, angle engine.Angle, factory FactoryFn)
}
