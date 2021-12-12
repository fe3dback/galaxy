package node

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

func (n *Node) AddComponent(c galx.Component) {
	if _, exist := n.components[c.Id()]; exist {
		panic(fmt.Sprintf("can`t add component '%s' (%s) to entity `%T`, already exist", c.Title(), c.Id(), n))
	}

	n.components[c.Id()] = c
}

func (n *Node) GetComponent(ref galx.Component) galx.Component {
	if c, exist := n.components[ref.Id()]; exist {
		return c
	}

	panic(fmt.Sprintf("can`t find component '%s' (%s) in `%T`", ref.Title(), ref.Id(), n))
}

func (n *Node) Components() map[UUID]galx.Component {
	return n.components
}
