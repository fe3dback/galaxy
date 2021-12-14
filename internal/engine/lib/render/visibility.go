package render

import (
	"github.com/fe3dback/galaxy/internal/engine"
)

func (r *Renderer) isLineInsideCamera(line []Point) bool {
	// line always have only two points (enforced by public api)
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return r.isRectInsideCamera(Rect{
		X: line[0].X,
		Y: line[0].Y,
		W: line[1].X - line[0].X,
		H: line[1].Y - line[0].Y,
	})
}

func (r *Renderer) isRectInsideCamera(rect Rect) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	// [right edge] outside of camera [left edge]
	if float64(rect.X+rect.W) < 0 {
		return false
	}

	// [bottom edge] outside of camera [top edge]
	if float64(rect.Y+rect.H) < 0 {
		return false
	}

	// [left edge] outside of camera [right edge]
	if float64(rect.X) > float64(r.camera.width)/r.camera.zoom {
		return false
	}

	// [top edge] outside of camera [bottom edge]
	if float64(rect.Y) > float64(r.camera.height)/r.camera.zoom {
		return false
	}

	return true
}

func (r *Renderer) isPointInsideCamera(p Point) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	// outside left edge
	if p.X < 0 {
		return false
	}

	// outside top edge
	if p.Y < 0 {
		return false
	}

	// outside right edge
	if float64(p.X) > float64(r.camera.width)/r.camera.zoom {
		return false
	}

	// outside bottom edge
	if float64(p.Y) > float64(r.camera.height)/r.camera.zoom {
		return false
	}

	return true
}
