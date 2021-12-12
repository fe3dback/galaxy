package control

import "github.com/fe3dback/galaxy/galx"

type Transform struct {
	centroid    galx.Vec
	hasSelected bool
}

func NewTransform() *Transform {
	return &Transform{}
}

func (c *Transform) OnUpdate(state galx.State) error {
	state.Mouse().SetPriority(galx.MousePropagationPriorityEditorGizmos)
	defer state.Mouse().ResetPriority()

	selectedObjects := state.ObjectQueries().AllIn(galx.QueryFlagOnlySelected)
	c.hasSelected = len(selectedObjects) > 0
	if !c.hasSelected {
		return nil
	}

	c.centroid = c.calculateCentroid(selectedObjects)

	return nil
}

func (c *Transform) calculateCentroid(list []galx.GameObject) galx.Vec {
	bbox := make([]galx.Rect, 0, len(list))

	for _, object := range list {
		bbox = append(bbox, object.BoundingBox(0))
	}

	return galx.SurroundRect(bbox...).Center()
}

func (c *Transform) OnDraw(r galx.Renderer) error {
	const arrowSize = 2
	const arrowLength = 50
	const circleRadius = 12

	if !c.hasSelected {
		return nil
	}

	// todo: mouse propagation on axis touch
	// todo: select/unselect axis to move/lock object
	// todo: move object on mouse move
	// todo: rotation square with polar coordinate line towards center

	// axis X
	axisX := galx.Vec{
		X: c.centroid.X + arrowLength,
		Y: c.centroid.Y,
	}

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
	r.DrawSquare(galx.ColorRed, galx.Rect{
		Min: axisX.Minus(circleRadius),
		Max: axisX.Plus(circleRadius),
	})

	// axis Y
	axisY := galx.Vec{
		X: c.centroid.X,
		Y: c.centroid.Y - arrowLength,
	}

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

	r.DrawSquare(galx.ColorGreen, galx.Rect{
		Min: axisY.Minus(circleRadius),
		Max: axisY.Plus(circleRadius),
	})

	// axis XY
	axisXY := galx.Vec{
		X: c.centroid.X + (circleRadius / 2),
		Y: c.centroid.Y - (circleRadius / 2),
	}

	r.DrawSquare(galx.ColorYellow, galx.Rect{
		Min: axisXY.Minus(circleRadius),
		Max: axisXY.Plus(circleRadius),
	})

	return nil
}
