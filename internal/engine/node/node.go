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
		relativePosition galx.Vec2d
		relativeRotation galx.Angle
		relativeScale    float64
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
		id:               id,
		name:             "",
		relativePosition: galx.Vec2d{},
		relativeRotation: galx.Angle0,
		relativeScale:    1,
		components:       make(components, 0),
		destroyed:        false,
		locked:           false,
		selected:         false,
		hierarchyLevel:   0,
		parent:           nil,
		child:            nil,
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
