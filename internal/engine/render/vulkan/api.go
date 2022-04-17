package vulkan

import (
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
)

func (vk *Vk) appendToRenderQueue(sp shaderProgram) {
	if queue, exist := vk.renderQueue[sp.ID()]; exist {
		vk.renderQueue[sp.ID()] = append(queue, sp)
		return
	}

	vk.renderQueue[sp.ID()] = []shaderProgram{sp}
}

func (vk *Vk) Clear(color galx.Color) {
	vk.frameBuffers.setClearColor(color)
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
	vk.renderQueue = make(map[string][]shaderProgram)
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

	// todo: uniform buffers for per instance data (transform, rotation, scale)
	// todo: global uniform buffer for mat4 (projection, view) - single frame buffer for all objects

	// reset
	drawCalls := 0
	drawInstances := 0
	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))

	// draw all shaders
	for pipelineID, instances := range vk.renderQueue {
		// bind pipeline for this group of shaders
		pipeline := vk.pipelineManager.pipeline(pipelineID)
		vulkan.CmdBindPipeline(commandBuffer, vulkan.PipelineBindPointGraphics, pipeline)

		// clean buffers
		vk.dataBuffersManager.resetBuffers()

		// copy data to buffers
		for _, instance := range instances {
			vk.dataBuffersManager.writeToBuffers(instance)
			drawInstances++
		}

		// flush buffers to GPU memory
		result := vk.dataBuffersManager.flushBuffers()

		// bind indexes
		vulkan.CmdBindIndexBuffer(commandBuffer, result.indexBuffer, 0, vulkan.IndexTypeUint16)
		indexesCount := uint32(len(instances[0].Indexes()))

		// draw commands
		for _, batch := range result.vertexChunks {
			vulkan.CmdBindVertexBuffers(commandBuffer, 0, uint32(1), batch.buffers, batch.offsets)
			vulkan.CmdDrawIndexed(commandBuffer, indexesCount*batch.instanceCount, 1, 0, 0, 0)
			drawCalls++
		}
	}

	log.Printf("draw calls=%d, instances=%d\n", drawCalls, drawInstances)
}

func (vk *Vk) GPUWait() {
	vulkan.DeviceWaitIdle(vk.ld.ref)
}
