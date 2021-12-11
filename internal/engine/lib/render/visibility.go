package render

import (
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine"
)

func (r *Renderer) isLineInsideCamera(line galx.Line) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return r.isRectInsideCamera(galx.Rect{
		Min: line.A,
		Max: line.B.Sub(line.A),
	}.Screen())
}

func (r *Renderer) isRectInsideCamera(rect galx.Rect) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	xB := rect.Min.X > r.camera.position.X+(float64(r.camera.width)/r.camera.zoom)
	xL := rect.Min.X+rect.Max.X < r.camera.position.X

	yB := rect.Min.Y > r.camera.position.Y+(float64(r.camera.height)/r.camera.zoom)
	yL := rect.Min.Y+rect.Max.Y < r.camera.position.Y

	return !(xB || xL || yB || yL)
}

func (r *Renderer) isPointInsideCamera(vec galx.Vec) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return vec.X >= r.camera.position.X &&
		vec.Y >= r.camera.position.Y &&
		vec.X <= r.camera.position.X+(float64(r.camera.width)/r.camera.zoom) &&
		vec.Y <= r.camera.position.Y+(float64(r.camera.height)/r.camera.zoom)
}
