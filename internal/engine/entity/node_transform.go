package entity

import "github.com/fe3dback/galaxy/galx"

func (e *Entity) RelativePosition() galx.Vec {
	return e.relativePosition
}

func (e *Entity) AbsPosition() galx.Vec {
	if !e.IsRoot() {
		return e.parent.AbsPosition().Add(e.RelativePosition())
	}

	// is root node, so its abs position on world
	return e.RelativePosition()
}

func (e *Entity) SetPosition(relativePosition galx.Vec) {
	e.relativePosition = relativePosition
}

func (e *Entity) AddPosition(v galx.Vec) {
	e.SetPosition(e.AbsPosition().Add(v))
}

func (e *Entity) Rotation() galx.Angle {
	return e.relativeRotation
}

func (e *Entity) SetRotation(rot galx.Angle) {
	e.relativeRotation = rot
}

func (e *Entity) AddRotation(rot galx.Angle) {
	e.SetRotation(e.Rotation().Add(rot))
}
