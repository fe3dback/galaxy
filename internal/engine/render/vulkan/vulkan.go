package vulkan

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/utils"
)

func NewVulkanApi(window *glfw.Window, dispatcher *event.Dispatcher, cfg *Config, closer *utils.Closer) *Vk {
	cont := newContainer(window, dispatcher, cfg, closer)
	renderer := cont.renderer()
	closer.EnqueueFree(renderer.free)

	return renderer
}

func (vk *Vk) free() {
	if vk.frameManager != nil {
		vk.frameManager.free()
		vk.frameManager = nil
	}

	if vk.commandPool != nil {
		vk.commandPool.free()
		vk.commandPool = nil
	}

	if vk.ld != nil {
		vk.ld.free()
		vk.ld = nil
	}

	vk.pd = nil

	if vk.surface != nil {
		vk.surface.free()
		vk.surface = nil
	}

	if vk.inst != nil {
		vk.inst.free()
		vk.inst = nil
	}
}

func (vk *Vk) rebuildGraphicsPipeline() {
	// todo: implement
}
