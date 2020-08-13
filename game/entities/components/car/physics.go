package car

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/entity"
)

type Vec = engine.Vec
type Angle = engine.Angle

type Physics struct {
	entity    *entity.Entity
	spec      spec
	movements *movements
}

func NewPhysics(entity *entity.Entity, yamlSpec YamlSpec) *Physics {
	phys := &Physics{
		entity: entity,
	}

	phys.spec = phys.assembleSpec(yamlSpec)
	phys.movements = newMovements(entity.Position(), entity.Rotation(), phys.spec)

	return phys
}
