package entity

import "github.com/fe3dback/galaxy/galx"

func (e *Entity) IsRoot() bool {
	return e.parent == nil
}

func (e *Entity) IsLeaf() bool {
	return len(e.child) == 0
}

func (e *Entity) Child() []galx.GameObject {
	return e.child
}

func (e *Entity) AddChild(child galx.GameObject) {
	e.child = append(e.child, child)
}

func (e *Entity) RemoveChild(id UUID) {
	newChild := make([]galx.GameObject, 0)

	for _, object := range e.child {
		if object.Id() == id {
			continue
		}

		newChild = append(newChild, object)
	}

	e.child = newChild
}

func (e *Entity) SetParent(parent galx.GameObject) {
	e.parent = parent
}

func (e *Entity) garbageCollect() {
	child := make([]galx.GameObject, 0, len(e.child))

	for _, childEntity := range e.child {
		if childEntity.IsDestroyed() {
			continue
		}

		child = append(child, childEntity)
	}

	e.child = child
}
