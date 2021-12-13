package control

import (
	"github.com/fe3dback/galaxy/galx"
)

type Select struct {
	objectQueries galx.ObjectQueries
}

func NewSelect(objectQueries galx.ObjectQueries) *Select {
	return &Select{
		objectQueries: objectQueries,
	}
}

func (c *Select) OnUpdate(state galx.State) error {
	if !state.Mouse().IsButtonsAvailable(galx.MousePropagationPriorityEditorSelect) {
		return nil
	}

	if !state.Mouse().LeftReleased() {
		return nil
	}

	// mouse just clicked
	clickWorldPos := state.Camera().Screen2World(state.Mouse().MouseCoords())
	foundObject := c.objectAt(clickWorldPos)

	// switch object state && apply to select group
	if state.Movement().Shift() {
		state.Mouse().StopPropagation(galx.MousePropagationPriorityEditorSelect)
		if foundObject == nil {
			return nil
		}

		if foundObject.IsSelected() {
			foundObject.Unselect()
		} else {
			foundObject.Select()
		}
		return nil
	}

	// set as selected, reset all another
	for _, anyObject := range state.ObjectQueries().All() {
		anyObject.Unselect()
	}

	if foundObject != nil {
		state.Mouse().StopPropagation(galx.MousePropagationPriorityEditorSelect)
		foundObject.Select()
	}

	return nil
}

func (c *Select) objectAt(clickPosition galx.Vec) galx.GameObject {
	const selectPrecision = 4

	var current galx.GameObject
	minLevel := uint8(255)

	for _, screenObject := range c.objectQueries.AllIn(galx.QueryFlagOnlyOnScreen | galx.QueryFlagExcludeLocked) {
		bbox := screenObject.BoundingBox(selectPrecision)
		if bbox.Contains(clickPosition) && screenObject.HierarchyLevel() < minLevel {
			current = screenObject
			minLevel = screenObject.HierarchyLevel()
		}
	}

	return current
}

func (c *Select) OnDraw(_ galx.Renderer) error {
	return nil
}
