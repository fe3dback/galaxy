package render

import (
	"github.com/fe3dback/galaxy/galx"
)

type Render struct {
	renderer renderer
}

func NewRender(renderer renderer) *Render {
	return &Render{
		renderer: renderer,
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

func (r *Render) SetRenderMode(mode galx.RenderMode) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) StartEngineFrame(color galx.Color) {
	r.renderer.FrameStart()
	r.renderer.Clear(uint32(color))
}

func (r *Render) EndEngineFrame() {
	r.renderer.FrameEnd()
}

func (r *Render) StartGUIFrame(color galx.Color) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) EndGUIFrame() {
	// do nothing here
}

func (r *Render) UpdateGPU() {
	r.renderer.Draw()
}

func (r *Render) WaitGPUOperationsBeforeQuit() {
	r.renderer.GPUWait()
}
