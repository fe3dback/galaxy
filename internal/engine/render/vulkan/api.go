package vulkan

import "github.com/vulkan-go/vulkan"

func (vk *Vk) Clear(color uint32) {
	// todo: implement
	// todo: save color to ctx
}

func (vk *Vk) FrameStart() {
	if vk.inResizing || vk.isMinimized {
		vk.currentFrameAvailableForRender = false
		return
	}

	// 1. start frame
	vk.currentFrameImageID, vk.currentFrameAvailableForRender = vk.frameManager.frameStart(vk.swapChain)
	if !vk.currentFrameAvailableForRender {
		return
	}

	// 2. start command buffer
	vk.commandPool.commandBufferStart(int(vk.currentFrameImageID))

	// 3. start render pass
	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))
	renderPass := vk.container.defaultRenderPass()
	vk.frameBuffers.renderPassStart(int(vk.currentFrameImageID), commandBuffer, renderPass)
}

func (vk *Vk) FrameEnd() {
	if !vk.currentFrameAvailableForRender {
		return
	}

	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))

	// 3. end render pass
	vk.frameBuffers.renderPassEnd(commandBuffer)

	// 2. end command buffer
	vk.commandPool.commandBufferEnd(int(vk.currentFrameImageID))

	// 1. end frame
	vk.frameManager.frameEnd(vk.swapChain, commandBuffer)
}

func (vk *Vk) Draw() {
	if !vk.currentFrameAvailableForRender {
		return
	}

	// todo: implement
	// todo: bind CmdBindPipeline
	// todo: bind CmdBindVertexBuffers
	// todo: bind CmdDraw
}

func (vk *Vk) GPUWait() {
	vulkan.DeviceWaitIdle(vk.ld.ref)
}
