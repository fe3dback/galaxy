package node

import "github.com/fe3dback/galaxy/galx"

func (n *Node) IsRoot() bool {
	return n.parent == nil
}

func (n *Node) IsLeaf() bool {
	return len(n.child) == 0
}

func (n *Node) HasChild() bool {
	return len(n.child) > 0
}

func (n *Node) HasParent() bool {
	return n.parent != nil
}

func (n *Node) Child() []galx.GameObject {
	return n.child
}

func (n *Node) AddChild(child galx.GameObject) {
	n.child = append(n.child, child)
}

func (n *Node) RemoveChild(id UUID) {
	newChild := make([]galx.GameObject, 0)

	for _, object := range n.child {
		if object.Id() == id {
			continue
		}

		newChild = append(newChild, object)
	}

	n.child = newChild
}

func (n *Node) SetParent(parent galx.GameObject) {
	n.parent = parent
	n.hierarchyLevel = parent.HierarchyLevel() + 1
}

func (n *Node) HierarchyLevel() uint8 {
	return n.hierarchyLevel
}

func (n *Node) garbageCollect() {
	child := make([]galx.GameObject, 0, len(n.child))

	for _, childEntity := range n.child {
		if childEntity.IsDestroyed() {
			continue
		}

		child = append(child, childEntity)
	}

	n.child = child
}
