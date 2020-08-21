package entities

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/collider"
	"github.com/fe3dback/galaxy/engine/entity"
)

type StaticFactoryParams struct {
	Collider *collider.Collider
}

func NewStaticFactory(params StaticFactoryParams) entity.Factory {
	return func(e *entity.Entity, creator engine.WorldCreator) *entity.Entity {
		e.AddCollider(params.Collider)

		return e
	}
}
