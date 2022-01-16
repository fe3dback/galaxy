package vulkan

import (
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/render/vulkan_depr/shader/shaderm"
)

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
	vk.renderQueue = []shaderProgram{}
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

func (vk *Vk) DrawTriangle() {
	vk.renderQueue = append(vk.renderQueue, &shaderm.Triangle{
		Position: [3]galx.Vec2{
			{X: 0.0, Y: -0.5},
			{X: 0.5, Y: 0.5},
			{X: -0.5, Y: 0.5},
		},
		Color: [3]galx.Vec3{
			{R: 1, G: 0, B: 0},
			{R: 0, G: 1, B: 0},
			{R: 0, G: 0, B: 1},
		},
	})
}

func (vk *Vk) Draw() {
	if !vk.currentFrameAvailableForRender {
		return
	}

	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))

	for _, shaderProgram := range vk.renderQueue {
		instanceCount := uint32(1) // todo: batch
		pipeline := vk.pipelineManager.pipeline(shaderProgram.ID())

		vulkan.CmdBindPipeline(commandBuffer, vulkan.PipelineBindPointGraphics, pipeline)

		// todo: vertexBuffer
		// todo: bindingCount = 1
		// vulkan.CmdBindVertexBuffers(commandBuffer, 0, 0, []vulkan.Buffer{}, []vulkan.DeviceSize{0})
		vulkan.CmdDraw(commandBuffer, shaderProgram.VertexCount(), instanceCount, 0, 0)
	}
}

func (vk *Vk) GPUWait() {
	vulkan.DeviceWaitIdle(vk.ld.ref)
}
