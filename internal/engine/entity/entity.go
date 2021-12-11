package entity

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type (
	UUID       = string
	components []Component

	Entity struct {
		id         UUID
		position   galx.Vec
		rotation   galx.Angle
		components components
		destroyed  bool
	}
)

func NewEntity(id UUID, pos galx.Vec, rot galx.Angle) *Entity {
	return &Entity{
		id:         id,
		position:   pos,
		rotation:   rot,
		components: make(components, 0),
		destroyed:  false,
	}
}

func (e *Entity) Id() UUID {
	return e.id
}

func (e *Entity) Position() galx.Vec {
	return e.position
}

func (e *Entity) SetPosition(pos galx.Vec) {
	e.position = pos
}

func (e *Entity) AddPosition(v galx.Vec) {
	e.SetPosition(e.Position().Add(v))
}

func (e *Entity) Rotation() galx.Angle {
	return e.rotation
}

func (e *Entity) SetRotation(rot galx.Angle) {
	e.rotation = rot
}

func (e *Entity) AddRotation(rot galx.Angle) {
	e.SetRotation(e.Rotation().Add(rot))
}

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(s galx.State) error {
	for _, component := range e.components {
		updatableComponent, ok := component.(ComponentCycleUpdated)
		if !ok {
			continue
		}

		err := updatableComponent.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component '%s' (%s): %w", e, component.Title(), component.Id(), err)
		}
	}

	return nil
}

func (e *Entity) OnDraw(r galx.Renderer) error {
	for _, component := range e.components {
		drawer, ok := component.(ComponentCycleDrawer)
		if !ok {
			continue
		}

		err := drawer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw entity `%T` component '%s' (%s): %w", e, component.Title(), component.Id(), err)
		}
	}

	if r.Gizmos().Primary() {
		r.DrawPoint(galx.ColorForeground, e.position)
		r.DrawVector(galx.ColorForeground, 10, e.position, e.rotation)
	}

	return nil
}

func (e *Entity) AddComponent(c Component) {
	for _, component := range e.components {
		if component.Id() == c.Id() {
			panic(fmt.Sprintf("can`t add component '%s' (%s) to entity `%T`, already exist", component.Title(), component.Id(), e))
		}
	}

	e.components = append(e.components, c)
}

func (e *Entity) GetComponent(ref Component) Component {
	for _, component := range e.components {
		if component.Id() == ref.Id() {
			return component
		}
	}

	panic(fmt.Sprintf("can`t find component '%s' (%s) in `%T`", ref.Title(), ref.Id(), e))
}
