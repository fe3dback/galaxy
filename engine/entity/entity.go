package entity

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/engine"
)

type (
	components []Component

	Entity struct {
		id         int64
		body       engine.PhysicsBody
		position   engine.Vec
		rotation   engine.Angle
		components components
		destroyed  bool
	}
)

var lastId int64 = 0

func NewEntity(pos engine.Vec, rot engine.Angle) *Entity {
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

func (e *Entity) AttachPhysicsBody(body engine.PhysicsBody) {
	if e.body != nil {
		panic(fmt.Sprintf("Entity `%d` already have physics body", e.Id()))
	}

	e.body = body
	e.updatePhysicsState()
}

func (e *Entity) Position() engine.Vec {
	return e.position
}

func (e *Entity) SetPosition(pos engine.Vec) {
	// todo: set phys pos
	if e.body != nil {
		// managed with physics
		return
	}

	e.position = pos
}

func (e *Entity) AddPosition(v engine.Vec) {
	e.SetPosition(e.Position().Add(v))
}

func (e *Entity) Rotation() engine.Angle {
	return e.rotation
}

func (e *Entity) SetRotation(rot engine.Angle) {
	// todo: set phys angle
	if e.body != nil {
		// managed with physics
		return
	}

	e.rotation = rot
}

func (e *Entity) AddRotation(rot engine.Angle) {
	e.SetRotation(e.Rotation().Add(rot))
}

func (e *Entity) ApplyForceFrom(force engine.Vec, relativePosition engine.Vec) {
	if e.body != nil {
		e.body.ApplyForce(force, e.position.Add(relativePosition))
	}
}

func (e *Entity) ApplyForce(force engine.Vec) {
	if e.body != nil {
		e.body.ApplyForce(force, e.position)
	}
}

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(s engine.State) error {
	if e.body != nil {
		e.updatePhysicsState()
	}

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

	if r.Gizmos().Primary() {
		r.DrawPoint(engine.ColorForeground, e.position)
		r.DrawVector(engine.ColorForeground, 10, e.position, e.rotation)
	}

	if r.Gizmos().Debug() {
		if e.body != nil {
			e.body.DebugDraw(r)
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

func (e *Entity) updatePhysicsState() {
	e.position = e.body.Position()
	e.rotation = e.body.Rotation()
}
