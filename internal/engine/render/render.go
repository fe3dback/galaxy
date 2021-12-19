package render

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

type Render struct {
	vkAPI *vk
}

func NewRender(closer *utils.Closer, window *glfw.Window, debugVulkan bool) *Render {
	vkApi := newVulkanApi(vkCreateOptions{
		closer:      closer,
		window:      window,
		debugVulkan: debugVulkan,
	})

	return &Render{
		vkAPI: vkApi,
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
	// TODO implement me
	panic("implement me")
}
