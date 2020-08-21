package entity

import (
	"fmt"
	"reflect"

	"github.com/fe3dback/galaxy/engine"
)

type (
	components []Component
	colliders  []*Collider

	Entity struct {
		id         int64
		position   engine.Vec
		rotation   engine.Angle
		components components
		colliders  colliders
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
		colliders:  make(colliders, 0),
		destroyed:  false,
	}
}

func (e *Entity) Id() int64 {
	return e.id
}

func (e *Entity) Position() engine.Vec {
	return e.position
}

func (e *Entity) SetPosition(pos engine.Vec) {
	e.position = pos
}

func (e *Entity) AddPosition(v engine.Vec) {
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

func (e *Entity) IsCollideWith(another *Entity) bool {
	if e.id == another.id {
		return false
	}

	// todo: quadratic complexity, rewrite to optimized
	// two step checking

	// todo: collision mask's
	for _, myCollider := range e.colliders {
		for _, anotherCollider := range another.colliders {
			if myCollider.IsCollideWith(anotherCollider) {
				return true
			}
		}
	}

	return false
}

func (e *Entity) OnUpdate(s engine.State) error {
	for _, component := range e.components {
		err := component.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component `%T`: %v", e, component, err)
		}
	}

	for _, collider := range e.colliders {
		collider.Update(e)
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
		for _, collider := range e.colliders {
			collider.OnDraw(r)
		}
	}

	return nil
}

func (e *Entity) OnCollide(target engine.Entity, _ uint8) {
	// todo: collision masks (layers)
	for _, component := range e.components {
		if resolver, ok := component.(engine.CollisionResolver); ok {
			resolver.OnCollide(target, 0)
		}
	}
}

func (e *Entity) AddCollider(c *Collider) {
	e.colliders = append(e.colliders, c)
}

func (e *Entity) Colliders() []*Collider {
	return e.colliders
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
