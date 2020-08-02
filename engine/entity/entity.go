package entity

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/engine"
)

type components []Component

type Entity struct {
	position   engine.Vector2D
	rotation   engine.Angle
	components components
	destroyed  bool
}

func NewEntity(pos engine.Vector2D, rot engine.Angle) *Entity {
	return &Entity{
		position:   pos,
		rotation:   rot,
		components: make(components, 0),
		destroyed:  false,
	}
}

func (e *Entity) Position() engine.Vector2D {
	return e.position
}

func (e *Entity) SetPosition(pos engine.Vector2D) {
	e.position = pos
}

func (e *Entity) AddPosition(v engine.Vector2D) {
	e.SetPosition(e.Position().Add(v))
}

func (e *Entity) Rotation() engine.Angle {
	return e.rotation
}

func (e *Entity) SetRotation(rot engine.Angle) {
	e.rotation = rot
}

func (e *Entity) AddRotation(rot engine.Angle) {
	e.SetRotation(e.Rotation().Add(rot))
}

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(s engine.State) error {
	for _, component := range e.components {
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component `%T`: %v", e, component, err)
		}
	}

	return nil
}

func (e *Entity) OnDraw(r engine.Renderer) error {
	for _, component := range e.components {
		err := component.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw entity `%T` component `%s`: %v", e, component, err)
		}
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
