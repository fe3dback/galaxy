package vulkan_depr

import (
	"github.com/vulkan-go/vulkan"
)

func (vk *Vk) onWindowResized(_, _ int) {
	vk.rebuildGraphicsPipeline()
}

func (vk *Vk) rebuildGraphicsPipeline() {
	if vk.isResizing {
		// already in rebuildGraphicsPipeline mode
		// wait for end
		return
	}

	vk.isResizing = true

	// wait for all current render is done
	vulkan.DeviceWaitIdle(vk.logicalDevice.ref)

	// free all pipeline staff
	vk.frameBuffers.free()
	vk.commandPool.free()
	vk.pipeLine.free()
	vk.swapChain.free()

	wWidth, wHeight := vk.windowSizeExtractor()
	if wWidth == 0 || wHeight == 0 {
		// window is minimized now, just wait for next resize for
		// swapChain recreate

		vk.isResizing = false
		vk.isDrawAvailable = false
		return
	}

	// create new pipeline
	swapChain, pipeline, frameBuffers, commandPool :=
		vk.swapChainFactory.createAllPipeline(vk.physicalDevice, vk.logicalDevice, vk.shaderManager, vk.closer)

	vk.frameBuffers = frameBuffers
	vk.commandPool = commandPool
	vk.pipeLine = pipeline
	vk.swapChain = swapChain
	vk.swapChainFrameManager.setSwapChain(swapChain)
	vk.swapChainFrameManager.setCommandPool(commandPool)

	// finalize
	vk.isResizing = false
	vk.isDrawAvailable = true
}
