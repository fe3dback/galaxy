package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type StaticFactoryParams struct {
}

func NewStaticFactory(_ StaticFactoryParams) entity.Factory {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		return e
	}
}
