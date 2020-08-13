package entity

import "github.com/fe3dback/galaxy/engine"

type Factory = func(*Entity, engine.WorldCreator) *Entity
