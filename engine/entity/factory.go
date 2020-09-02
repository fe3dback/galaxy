package entity

import "github.com/fe3dback/galaxy/engine"

type FactoryFn = func(*Entity, engine.WorldCreator) *Entity
