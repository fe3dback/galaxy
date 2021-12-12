package node

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

type (
	UUID       = string
	components map[UUID]galx.Component

	Node struct {
		id               UUID
		name             string
		relativePosition galx.Vec
		relativeRotation galx.Angle
		components       components
		destroyed        bool
		locked           bool
		selected         bool

		// hierarchy
		hierarchyLevel uint8
		parent         galx.GameObject
		child          []galx.GameObject
	}
)

func NewNode(id UUID) *Node {
	return &Node{
		id:         id,
		components: make(components, 0),
	}
}

func (n *Node) Id() UUID {
	return n.id
}

func (n *Node) Name() string {
	if n.name == "" {
		return "node"
	}

	return n.name
}

func (n *Node) SetName(name string) {
	n.name = name
}

func (n *Node) String() string {
	return fmt.Sprintf("'%s' (%s)", n.name, n.Id())
}
