package render

import "github.com/fe3dback/galaxy/engine"

func (r *Renderer) screenX(x float64) float64 {
	if r.renderMode == engine.RenderModeUI {
		return x
	}

	return x - r.camera.position.X
}

func (r *Renderer) screenY(y float64) float64 {
	if r.renderMode == engine.RenderModeUI {
		return y
	}

	return y - r.camera.position.Y
}
