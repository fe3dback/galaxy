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

	const rSizeX = 24
	const rSizeY = 32
	const rOffset = 10
	const rStartX = rSizeX
	const rEndX = 1280 - rSizeX - rOffset
	const rStartY = rSizeY
	const rEndY = 720 - rSizeY - rOffset

	for x := rStartX; x < rEndX; x += rSizeX + rOffset {
		for y := rStartY; y < rEndY; y += rSizeY + rOffset {
			r.DrawSquare(galx.ColorCyan, galx.Rect{
				TL: galx.Vec{
					X: float64(x),
					Y: float64(y),
				},
				BR: galx.Vec{
					X: float64(x + rSizeX),
					Y: float64(y + rSizeY),
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
