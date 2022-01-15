package vulkan

func (vk *Vk) Clear(color uint32) {
	// todo: implement
	// todo: save color to ctx
}

func (vk *Vk) FrameStart() {
	swapChain := &vkSwapChain{} // todo: ??

	// 1. start frame
	vk.currentFrameImageID, vk.currentFrameAvailableForRender = vk.frameManager.frameStart(swapChain)
	if !vk.currentFrameAvailableForRender {
		return
	}

	// 2. start command buffer
	vk.commandPool.commandBufferStart(int(vk.currentFrameImageID))

	// 3. start render pass
	// todo: render-pass begin
}

func (vk *Vk) FrameEnd() {
	if !vk.currentFrameAvailableForRender {
		return
	}

	// 3. end render pass
	// todo: render-pass end

	// 2. end command buffer
	vk.commandPool.commandBufferEnd(int(vk.currentFrameImageID))

	// 1. end frame
	swapChain := &vkSwapChain{} // todo: ??
	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))
	vk.frameManager.frameEnd(swapChain, commandBuffer)
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
	// todo: implement
}
