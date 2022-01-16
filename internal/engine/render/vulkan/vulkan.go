package vulkan

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/utils"
)

func NewVulkanApi(window *glfw.Window, dispatcher *event.Dispatcher, cfg *Config, closer *utils.Closer) *Vk {
	cont := newContainer(window, dispatcher, cfg, closer)
	renderer := cont.renderer()
	closer.EnqueueFree(renderer.free)

	// subscribe to system events
	dispatcher.OnWindowResized(func(_ event.WindowResizedEvent) error {
		renderer.rebuildGraphicsPipeline()
		return nil
	})

	return renderer
}

func (vk *Vk) free() {
	if vk.pipelineLayout != nil {
		vulkan.DestroyPipelineLayout(vk.ld.ref, vk.pipelineLayout, nil)
		vk.pipelineLayout = nil
	}

	if vk.pipelineManager != nil {
		vk.pipelineManager.free()
		vk.pipelineManager = nil
	}

	if vk.shaderManager != nil {
		vk.shaderManager.free()
		vk.shaderManager = nil
	}

	if vk.frameBuffers != nil {
		vk.frameBuffers.free()
		vk.frameBuffers = nil
	}

	if vk.swapChain != nil {
		vk.swapChain.free()
		vk.swapChain = nil
	}

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
	// resize mode
	// ----------------------------

	if vk.inResizing {
		// already in rebuildGraphicsPipeline mode
		// wait for end
		return
	}

	vk.inResizing = true

	// wait for all current render is done
	// ----------------------------
	vulkan.DeviceWaitIdle(vk.ld.ref)

	// free all pipeline staff
	// ----------------------------

	if vk.pipelineManager != nil {
		vk.pipelineManager.free()
		vk.pipelineManager = nil
		vk.container.vkPipelineManager = nil
	}

	vk.container.vkRenderPassHandlesLazyCache = make(map[renderPassType]vulkan.RenderPass)

	if vk.frameBuffers != nil {
		vk.frameBuffers.free()
		vk.frameBuffers = nil
		vk.container.vkFrameBuffers = nil
	}

	if vk.swapChain != nil {
		vk.swapChain.free()
		vk.swapChain = nil
		vk.container.vkSwapChain = nil
	}

	if vk.commandPool != nil {
		vk.commandPool.free()
		vk.commandPool = nil
		vk.container.vkCommandPool = nil
	}

	// minimization handle
	// ----------------------------

	vk.isMinimized = false
	wWidth, wHeight := vk.container.window.GetFramebufferSize()
	if wWidth == 0 || wHeight == 0 {
		// window is minimized now, just wait for next resize for
		// swapChain recreate

		vk.inResizing = false
		vk.isMinimized = true
		return
	}

	// recreate vk objects
	// ----------------------------

	vk.commandPool = vk.container.provideVkCommandPool()
	vk.swapChain = vk.container.provideSwapChain()
	vk.frameBuffers = vk.container.provideFrameBuffers()
	vk.pipelineManager = vk.container.providePipelineManager()

	for shader, pipelineFactory := range buildInShaders {
		vk.pipelineManager.preloadPipelineFor(shader, pipelineFactory(vk.container, shader))
	}

	// finalize
	// ----------------------------

	vk.inResizing = false
}
