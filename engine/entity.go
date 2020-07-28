package engine

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/render"
)

type components map[string]Component

type Entity struct {
	position   Vector2D
	rotation   Angle
	components components
	destroyed  bool
}

func NewEntity(pos Vector2D, rot Angle) *Entity {
	return &Entity{
		position:   pos,
		rotation:   rot,
		components: make(components),
		destroyed:  false,
	}
}

func (e *Entity) SetPosition(pos Vector2D) {
	e.position = pos

	fmt.Printf("entity `%T` moved to %v\n", e, pos)
}

func (e *Entity) Position() Vector2D {
	return e.position
}

func (e *Entity) SetRotation(rot Angle) {
	e.rotation = rot
}

func (e *Entity) Rotation() Angle {
	return e.rotation
}

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(deltaTime float64) error {
	for id, component := range e.components {
		err := component.OnUpdate(deltaTime)
		if err != nil {
			return fmt.Errorf("can`t update component `%s` from element `%T`: %v", id, e, err)
		}
	}

	return nil
}

func (e *Entity) OnDraw(r *render.Renderer) error {
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
