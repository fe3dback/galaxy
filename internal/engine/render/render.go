package render

import (
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/render/vulkan"
)

type Render struct {
	vkAPI *vulkan.Vk
}

func NewRender(vk *vulkan.Vk) *Render {
	return &Render{
		vkAPI: vk,
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
	// TODO implement me
	panic("implement me")
}

func (r *Render) EndEngineFrame() {
	// TODO implement me
	panic("implement me")
}

func (r *Render) StartGUIFrame(color galx.Color) {
	// TODO implement me
	panic("implement me")
}

func (r *Render) EndGUIFrame() {
	// do nothing here
}

func (r *Render) UpdateGPU() {
	r.vkAPI.Draw()
}

func (r *Render) WaitGPUOperationsBeforeQuit() {
	r.vkAPI.Wait()
}
