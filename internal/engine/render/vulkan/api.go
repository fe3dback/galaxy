package vulkan

import (
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/engine/render/vulkan/shader/shaderm"
)

func (vk *Vk) appendToRenderQueue(sp shaderProgram) {
	if queue, exist := vk.renderQueue[sp.ID()]; exist {
		vk.renderQueue[sp.ID()] = append(queue, sp)
		return
	}

	vk.renderQueue[sp.ID()] = []shaderProgram{sp}
}

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

func (vk *Vk) DrawTriangle() {
	for i := float32(-1); i < 1; i += 0.01 {
		vk.appendToRenderQueue(&shaderm.Triangle{
			Position: [3]galx.Vec2{
				{X: i, Y: -0.5},
				{X: 0.5, Y: 0.5},
				{X: -0.5, Y: 0.5},
			},
			Color: [3]galx.Vec3{
				{R: (i + 1) / 2, G: 0, B: 0},
				{R: 0, G: 1, B: 0},
				{R: 0, G: 0, B: 1},
			},
		})
	}
}

func (vk *Vk) Draw() {
	if !vk.currentFrameAvailableForRender {
		return
	}

	// reset
	drawCalls := 0
	commandBuffer := vk.commandPool.commandBuffer(int(vk.currentFrameImageID))
	vertexBuffersStats := vk.dataBuffersManager.resetVertexBuffers()

	// todo:  uint32(1) - what is bindingCount?
	vulkan.CmdBindVertexBuffers(commandBuffer, 0, uint32(1), vertexBuffersStats.buffers, vertexBuffersStats.offsets)

	// draw all shaders
	for pipelineID, instances := range vk.renderQueue {
		// bind pipeline for this group of shaders
		pipeline := vk.pipelineManager.pipeline(pipelineID)
		vulkan.CmdBindPipeline(commandBuffer, vulkan.PipelineBindPointGraphics, pipeline)

		// copy vertex data to buffers
		for i, instance := range instances {
			vk.dataBuffersManager.writeToVertexBuffers(instance)

			// todo: split draw calls (for test only):
			vulkan.CmdDraw(commandBuffer, instances[0].VertexCount(), uint32(1), uint32(i)*3, 0)
			drawCalls++
		}

		// draw all instances of this shader
		// todo: vulkan.CmdDraw
		// drawCalls++ // todo
	}

	// flush vertex buffers
	vk.dataBuffersManager.flushVertexBuffers()

	log.Printf("draw calls: %d\n", drawCalls)
}

func (vk *Vk) GPUWait() {
	vulkan.DeviceWaitIdle(vk.ld.ref)
}
