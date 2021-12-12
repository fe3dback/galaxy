package node

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

func (n *Node) Destroy() {
	n.destroyed = true
}

func (n *Node) IsDestroyed() bool {
	return n.destroyed
}

func (n *Node) OnUpdate(s galx.State) error {
	// update self
	for _, component := range n.components {
		updatableComponent, ok := component.(galx.ComponentCycleUpdated)
		if !ok {
			continue
		}

		err := updatableComponent.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component '%s' (%s): %w", n, component.Title(), component.Id(), err)
		}
	}

	// update child
	needGc := false
	for _, child := range n.child {
		if child.IsDestroyed() {
			needGc = true
			continue
		}

		err := child.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update child entity %s from %s: %w", child, n, err)
		}
	}

	if needGc {
		n.garbageCollect()
	}

	return nil
}

func (n *Node) OnDraw(r galx.Renderer) error {
	// draw self
	for _, component := range n.components {
		drawer, ok := component.(galx.ComponentCycleDrawer)
		if !ok {
			continue
		}

		err := drawer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw entity `%T` component '%s' (%s): %w", n, component.Title(), component.Id(), err)
		}
	}

	if r.Gizmos().Primary() {
		r.DrawCrossLines(n.gizmosColor(), 5, n.AbsPosition())
	}

	if r.Gizmos().Secondary() {
		r.DrawVector(n.gizmosColor(), 10, n.AbsPosition(), n.Rotation())
		r.DrawSquare(n.gizmosColor(), n.BoundingBox(4).MaxToSize())
	}

	// draw child
	for _, child := range n.child {
		if n.IsDestroyed() {
			continue
		}

		err := child.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw child entity %s from %s: %w", child, n, err)
		}
	}

	return nil
}

func (n *Node) gizmosColor() galx.Color {
	if n.IsSelected() {
		return galx.ColorOrange
	}

	return galx.ColorForeground
}
