package entity

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/engine"
)

type components map[string]Component

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
		components: make(components),
		destroyed:  false,
	}
}

func (e *Entity) SetPosition(pos engine.Vector2D) {
	e.position = pos

	fmt.Printf("entity `%T` moved to %v\n", e, pos)
}

func (e *Entity) Position() engine.Vector2D {
	return e.position
}

func (e *Entity) SetRotation(rot engine.Angle) {
	e.rotation = rot
}

func (e *Entity) Rotation() engine.Angle {
	return e.rotation
}

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(s engine.State) error {
	for id, component := range e.components {
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update component `%s` from element `%T`: %v", id, e, err)
		}
	}

	return nil
}

func (e *Entity) OnDraw(r engine.Renderer) error {
	for id, component := range e.components {
		err := component.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw component `%s` from element `%T`: %v", id, e, err)
		}
	}

	return nil
}

func (e *Entity) AddComponent(c Component) {
	id := reflect.TypeOf(c).String()

	if _, ok := e.components[id]; ok {
		panic(fmt.Sprintf("can`t add component `%s` to element, already exist", id))
	}

	e.components[id] = c
}

func (e *Entity) GetComponent(ref Component) Component {
	id := reflect.TypeOf(ref).String()

	if c, ok := e.components[id]; ok {
		return c
	}

	panic(fmt.Sprintf("can`t find component by id: %v", id))
}
