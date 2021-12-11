package entity

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/galx"
)

type (
	components []Component

	Entity struct {
		id         int64
		position   galx.Vec
		rotation   galx.Angle
		components components
		destroyed  bool
	}
)

var lastId int64 = 0

func NewEntity(pos galx.Vec, rot galx.Angle) *Entity {
	lastId++
	return &Entity{
		id:         lastId,
		position:   pos,
		rotation:   rot,
		components: make(components, 0),
		destroyed:  false,
	}
}

func (e *Entity) Id() int64 {
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
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component `%T`: %w", e, component, err)
		}
	}

	return nil
}

func (e *Entity) OnDraw(r galx.Renderer) error {
	for _, component := range e.components {
		err := component.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw entity `%T` component `%s`: %w", e, component, err)
		}
	}

	if r.Gizmos().Primary() {
		r.DrawPoint(galx.ColorForeground, e.position)
		r.DrawVector(galx.ColorForeground, 10, e.position, e.rotation)
	}

	return nil
}

func (e *Entity) AddComponent(c Component) {
	id := reflect.TypeOf(c).String()

	for _, component := range e.components {
		newId := reflect.TypeOf(component).String()
		if id == newId {
			panic(fmt.Sprintf("can`t add component `%s` to entity `%T`, already exist", id, e))
		}
	}

	e.components = append(e.components, c)
}

func (e *Entity) GetComponent(ref Component) Component {
	id := reflect.TypeOf(ref).String()

	for _, component := range e.components {
		newId := reflect.TypeOf(component).String()
		if id == newId {
			return component
		}
	}

	panic(fmt.Sprintf("can`t find component `%s` in `%T`", id, e))
}
