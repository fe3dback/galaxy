package node

import "github.com/fe3dback/galaxy/galx"

func (n *Node) RelativePosition() galx.Vec {
	return n.relativePosition
}

func (n *Node) AbsPosition() galx.Vec {
	if !n.IsRoot() {
		return n.parent.AbsPosition().Add(n.RelativePosition())
	}

	// is root node, so its abs position on world
	return n.RelativePosition()
}

func (n *Node) SetPosition(relativePosition galx.Vec) {
	n.relativePosition = relativePosition
}

func (n *Node) AddPosition(v galx.Vec) {
	n.SetPosition(n.AbsPosition().Add(v))
}

func (n *Node) Rotation() galx.Angle {
	return n.relativeRotation
}

func (n *Node) SetRotation(rot galx.Angle) {
	n.relativeRotation = rot
}

func (n *Node) AddRotation(rot galx.Angle) {
	n.SetRotation(n.Rotation().Add(rot))
}

func (n *Node) Scale() float64 {
	return n.relativeScale
}

func (n *Node) SetScale(scale float64) {
	n.relativeScale = scale
}
