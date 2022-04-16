package render

import (
	"github.com/fe3dback/galaxy/galx"
)

type Render struct {
	renderer renderer

	renderMode galx.RenderMode
	camera     *Camera
}

func NewRender(renderer renderer, camera *Camera) *Render {
	return &Render{
		renderer: renderer,
		camera:   camera,
	}
}

func (r *Render) Gizmos() galx.Gizmos {
	// TODO implement me
	panic("implement me")
}

func (r *Render) SetRenderTarget(id galx.RenderTarget) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) StartEngineFrame(color galx.Color) {
	r.renderMode = galx.RenderModeWorld
	r.renderer.Clear(color)
	r.renderer.FrameStart()
}

func (r *Render) EndEngineFrame() {
	r.renderer.Draw()
	r.renderer.FrameEnd()
}

// DrawTemporary todo: remove tmp
func (r *Render) DrawTemporary() {
	// r.renderer.DrawTmpTriangle()

	for x := 50; x < 1000; x += 50 {
		for y := 50; y < 720; y += 50 {
			r.DrawSquare(galx.ColorCyan, galx.Rect{
				Min: galx.Vec{
					X: float64(x),
					Y: float64(y),
				},
				Max: galx.Vec{
					X: float64(x + 25),
					Y: float64(y + 40),
				},
			})
		}
	}

	// r.renderer.DrawTmpTriangle()
}

func (r *Render) StartGUIFrame(color galx.Color) {
	r.renderMode = galx.RenderModeUI
	// TODO implement me
	panic("implement me")
}

func (r *Render) EndGUIFrame() {
	// do nothing here
}

func (r *Render) WaitGPUOperationsBeforeQuit() {
	r.renderer.GPUWait()
}
