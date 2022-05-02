package render

import (
	"github.com/fe3dback/galaxy/galx"
)

func (r *Render) camX(x float64) float64 {
	if r.renderMode == galx.RenderModeUI {
		return x
	}

	return x - r.camera.position.X
}

func (r *Render) camY(y float64) float64 {
	if r.renderMode == galx.RenderModeUI {
		return y
	}

	return y - r.camera.position.Y
}

func (r *Render) cam(vec galx.Vec2d) galx.Vec2d {
	if r.renderMode == galx.RenderModeUI {
		return vec
	}

	return galx.Vec2d{
		X: r.camX(vec.X),
		Y: r.camY(vec.Y),
	}
}

func (r *Render) projectX(x float64) float64 {
	return (x/float64(r.camera.Width()))*2 - 1
}

func (r *Render) projectY(y float64) float64 {
	return (y/float64(r.camera.Height()))*2 - 1
}

func (r *Render) project(vec galx.Vec2d) galx.Vec2d {
	return galx.Vec2d{
		X: r.projectX(vec.X),
		Y: r.projectY(vec.Y),
	}
}
