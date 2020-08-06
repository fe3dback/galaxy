package car

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
	"github.com/fe3dback/galaxy/generated"
)

type Vec = engine.Vec
type Angle = engine.Angle

type Physics struct {
	entity    *entity.Entity
	resource  generated.ResourcePath
	spec      spec
	movements *movements
}

func NewPhysics(entity *entity.Entity, resource generated.ResourcePath) *Physics {
	phys := &Physics{
		entity:   entity,
		resource: resource,
	}

	phys.assembleSpec()
	phys.movements = newMovements(entity.Position(), entity.Rotation(), phys.spec)

	return phys
}
