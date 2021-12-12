package entity

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

func (e *Entity) AddComponent(c galx.Component) {
	if _, exist := e.components[c.Id()]; exist {
		panic(fmt.Sprintf("can`t add component '%s' (%s) to entity `%T`, already exist", c.Title(), c.Id(), e))
	}

	e.components[c.Id()] = c
}

func (e *Entity) GetComponent(ref galx.Component) galx.Component {
	if c, exist := e.components[ref.Id()]; exist {
		return c
	}

	panic(fmt.Sprintf("can`t find component '%s' (%s) in `%T`", ref.Title(), ref.Id(), e))
}

func (e *Entity) Components() map[UUID]galx.Component {
	return e.components
}
