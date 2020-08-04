package car

import (
	"github.com/fe3dback/galaxy/engine"
)

func (phys *Physics) OnDraw(r engine.Renderer) error {
	if r.Gizmos().Secondary() {
		phys.debugDrawBoundingBox(r)
	}

	// debug weights
	if r.Gizmos().Debug() {
		phys.debugDrawWeights(r)
	}

	// draw wheels (movement component)
	phys.movements.draw(r)

	// ok
	return nil
}

func (phys *Physics) OnUpdate(s engine.State) error {
	newPos, newDirection := phys.movements.update(s)

	phys.entity.SetPosition(newPos)
	phys.entity.SetRotation(newDirection)

	return nil
}

func (phys *Physics) debugDrawBoundingBox(r engine.Renderer) {
	carSize := phys.spec.model.size
	carPos := phys.entity.Position()

	// draw bounding box
	// 90 spin because car width is not image width, is car width (and car angle direct to right)
	r.DrawSquareEx(engine.ColorSelection, phys.entity.Rotation(), engine.RectScreen(
		int(carPos.X)-carSize.width/2,
		int(carPos.Y)-carSize.height/2,
		carSize.width,
		carSize.height,
	))
}

func (phys *Physics) debugDrawWeights(r engine.Renderer) {
	var col engine.Color
	for _, weightTarget := range phys.spec.weights {
		weight := weightTarget.weight

		if weight >= 0.8 {
			col = engine.ColorGreen
		} else if weight >= 0.6 {
			col = engine.ColorOrange
		} else if weight >= 0.4 {
			col = engine.ColorYellow
		} else if weight >= 0.2 {
			col = engine.ColorRed
		} else {
			col = engine.ColorPink
		}

		vector := phys.entity.Position().
			Add(engine.Vec{
				X: float64(weightTarget.posRelative.x),
				Y: float64(weightTarget.posRelative.y),
			}).
			RotateAround(phys.entity.Position(), phys.entity.Rotation())

		r.DrawCrossLines(col, 3, vector)
	}
}
