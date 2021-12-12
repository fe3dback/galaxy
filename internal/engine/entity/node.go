package entity

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type (
	UUID       = string
	components map[UUID]galx.Component

	Entity struct {
		id               UUID
		name             string
		relativePosition galx.Vec
		relativeRotation galx.Angle
		components       components
		destroyed        bool

		// hierarchy
		parent galx.GameObject
		child  []galx.GameObject
	}
)

func NewEntity(id UUID, relPosition galx.Vec, relRotation galx.Angle) *Entity {
	return &Entity{
		id:               id,
		name:             "",
		relativePosition: relPosition,
		relativeRotation: relRotation,
		components:       make(components, 0),
		destroyed:        false,
	}
}

func (e *Entity) Id() UUID {
	return e.id
}

func (e *Entity) Name() string {
	if e.name == "" {
		return "node"
	}

	return e.name
}

func (e *Entity) SetName(name string) {
	e.name = name
}

func (e *Entity) String() string {
	return fmt.Sprintf("'%s' (%s)", e.name, e.Id())
}
