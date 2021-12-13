package control

import "github.com/fe3dback/galaxy/galx"

const arrowSize = 1
const arrowLength = 40
const circleRadius = 9

type Transform struct {
	centroid             galx.Vec
	selectedObjects      []galx.GameObject
	axisX                galx.Rect
	axisY                galx.Rect
	axisXY               galx.Rect
	activeX              bool
	activeY              bool
	startMousePosition   galx.Vec
	startObjectsPosition []galx.Vec
}

func NewTransform() *Transform {
	return &Transform{}
}

func (c *Transform) OnUpdate(state galx.State) error {
	if !state.Mouse().IsButtonsAvailable(galx.MousePropagationPriorityEditorGizmos) {
		return nil
	}

	// disable move mode, when mouse released
	if state.Mouse().LeftReleased() {
		c.activeX = false
		c.activeY = false
	}

	mousePos := state.Mouse().MouseCoords()
	worldMouse := state.Camera().Screen2World(mousePos)

	// if move mode active
	// disable mouse actions
	// and move objects
	if (c.activeX || c.activeY) && state.Mouse().LeftDown() {
		state.Mouse().StopPropagation(galx.MousePropagationPriorityEditorGizmos)

		if state.Mouse().LeftDown() {
			c.move(worldMouse)
		}
		return nil
	}

	// find selected objects for draw gizmos
	c.selectedObjects = state.ObjectQueries().AllIn(galx.QueryFlagOnlySelected)
	if len(c.selectedObjects) == 0 {
		return nil
	}

	c.centroid = c.calculateCentroid()
	c.updateAxisX()
	c.updateAxisY()
	c.updateAxisXY()

	// if mouse not pressed, we not
	// touch any controls on this frame
	// so nothing to do next
	if !state.Mouse().LeftPressed() {
		return nil
	}

	// turn on move mode
	if c.axisXY.Contains(worldMouse) {
		c.activeX = true
		c.activeY = true
	} else {
		c.activeX = c.axisX.Contains(worldMouse)
		c.activeY = c.axisY.Contains(worldMouse)
	}

	// if turned, save objects state
	if c.activeX || c.activeY {
		c.startMousePosition = worldMouse
		c.startObjectsPosition = make([]galx.Vec, len(c.selectedObjects))
		for idx, object := range c.selectedObjects {
			c.startObjectsPosition[idx] = object.AbsPosition()
		}
	}

	return nil
}

func (c *Transform) move(worldMouse galx.Vec) {
	diff := worldMouse.Sub(c.startMousePosition)

	for idx, object := range c.selectedObjects {
		newPosition := c.startObjectsPosition[idx].Add(diff)

		// exclude X movement
		if !c.activeX {
			newPosition.X = c.startObjectsPosition[idx].X
		}

		// exclude Y movement
		if !c.activeY {
			newPosition.Y = c.startObjectsPosition[idx].Y
		}

		object.SetPosition(newPosition)
	}
}

func (c *Transform) updateAxisX() {
	axisX := galx.Vec{
		X: c.centroid.X + arrowLength,
		Y: c.centroid.Y,
	}
	c.axisX = galx.Rect{
		Min: axisX.Minus(circleRadius),
		Max: axisX.Plus(circleRadius),
	}
}

func (c *Transform) updateAxisY() {
	axisY := galx.Vec{
		X: c.centroid.X,
		Y: c.centroid.Y - arrowLength,
	}
	c.axisY = galx.Rect{
		Min: axisY.Minus(circleRadius),
		Max: axisY.Plus(circleRadius),
	}
}

func (c *Transform) updateAxisXY() {
	axisXY := galx.Vec{
		X: c.centroid.X + (circleRadius / 2),
		Y: c.centroid.Y - (circleRadius / 2),
	}
	c.axisXY = galx.Rect{
		Min: axisXY.Minus(circleRadius),
		Max: axisXY.Plus(circleRadius),
	}
}

func (c *Transform) calculateCentroid() galx.Vec {
	bbox := make([]galx.Rect, 0, len(c.selectedObjects))

	for _, object := range c.selectedObjects {
		bbox = append(bbox, object.BoundingBox(0))
	}

	return galx.SurroundRect(bbox...).Center()
}

func (c *Transform) OnDraw(r galx.Renderer) error {
	if len(c.selectedObjects) == 0 {
		return nil
	}

	// todo: select bbox (over all objects)
	// todo: bigger select box (with negative shape)
	// todo: do not draw controls on move
	// todo: draw axis when excluded
	// todo: keyboard X,Y toggle exclude mode
	// todo: ctrl - snap to grid
	// todo: better control UX
	// todo: rotation square with polar coordinate line towards center
	// todo: fix propagate ctl on mouse release

	// axis X
	r.DrawSquareFilled(galx.ColorRed, galx.Rect{
		Min: galx.Vec{
			X: c.centroid.X - arrowSize,
			Y: c.centroid.Y - arrowSize,
		},
		Max: galx.Vec{
			X: c.centroid.X + arrowLength,
			Y: c.centroid.Y + arrowSize,
		},
	})
	r.DrawSquare(galx.ColorRed, c.axisX)
	if c.activeX {
		r.DrawSquare(galx.ColorRed, c.axisX.Scale(1.05))
	}

	// axis Y
	r.DrawSquareFilled(galx.ColorGreen, galx.Rect{
		Min: galx.Vec{
			X: c.centroid.X - arrowSize,
			Y: c.centroid.Y + arrowSize,
		},
		Max: galx.Vec{
			X: c.centroid.X + arrowSize,
			Y: c.centroid.Y - arrowLength,
		},
	})

	r.DrawSquare(galx.ColorGreen, c.axisY)
	if c.activeY {
		r.DrawSquare(galx.ColorGreen, c.axisY.Scale(1.05))
	}

	return nil
}
