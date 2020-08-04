package render

import "github.com/fe3dback/galaxy/engine"

func (r *Renderer) isLineInsideCamera(line engine.Line) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return r.isRectInsideCamera(engine.Rect{
		Min: line.A,
		Max: line.B.Sub(line.A),
	}.Screen())
}

func (r *Renderer) isRectInsideCamera(rect engine.Rect) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	xB := rect.Min.X > r.camera.position.X+float64(r.camera.width)
	xL := rect.Min.X+rect.Max.X < r.camera.position.X

	yB := rect.Min.Y > r.camera.position.Y+float64(r.camera.height)
	yL := rect.Min.Y+rect.Max.Y < r.camera.position.Y

	return !(xB || xL || yB || yL)
}

func (r *Renderer) isPointInsideCamera(vec engine.Vec) bool {
	if r.renderMode == engine.RenderModeUI {
		return true
	}

	return vec.X >= r.camera.position.X &&
		vec.Y >= r.camera.position.Y &&
		vec.X <= r.camera.position.X+float64(r.camera.width) &&
		vec.Y <= r.camera.position.Y+float64(r.camera.height)
}
