package entity

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
)

func (e *Entity) Destroy() {
	e.destroyed = true
}

func (e *Entity) IsDestroyed() bool {
	return e.destroyed
}

func (e *Entity) OnUpdate(s galx.State) error {
	// update self
	for _, component := range e.components {
		updatableComponent, ok := component.(galx.ComponentCycleUpdated)
		if !ok {
			continue
		}

		err := updatableComponent.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update entity `%T` component '%s' (%s): %w", e, component.Title(), component.Id(), err)
		}
	}

	// update child
	needGc := false
	for _, child := range e.child {
		if child.IsDestroyed() {
			needGc = true
			continue
		}

		err := child.OnUpdate(s)
		if err != nil {
			return fmt.Errorf("can`t update child entity %s from %s: %w", child, e, err)
		}
	}

	if needGc {
		e.garbageCollect()
	}

	return nil
}

func (e *Entity) OnDraw(r galx.Renderer) error {
	// draw self
	for _, component := range e.components {
		drawer, ok := component.(galx.ComponentCycleDrawer)
		if !ok {
			continue
		}

		err := drawer.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw entity `%T` component '%s' (%s): %w", e, component.Title(), component.Id(), err)
		}
	}

	if r.Gizmos().Primary() {
		r.DrawPoint(galx.ColorForeground, e.AbsPosition())
		r.DrawVector(galx.ColorForeground, 25, e.AbsPosition(), e.Rotation())
	}

	// draw child
	for _, child := range e.child {
		if e.IsDestroyed() {
			continue
		}

		err := child.OnDraw(r)
		if err != nil {
			return fmt.Errorf("can`t draw child entity %s from %s: %w", child, e, err)
		}
	}

	return nil
}
