package vulkan_depr

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkCommandPool struct {
		ref     vulkan.CommandPool
		buffers []vulkan.CommandBuffer

		_free   bool
		_freeLd *vkLogicalDevice
	}
)

func (pool *vkCommandPool) free() {
	if pool._free {
		return
	}

	pool._free = true
	vulkan.FreeCommandBuffers(pool._freeLd.ref, pool.ref, uint32(len(pool.buffers)), pool.buffers)
	vulkan.DestroyCommandPool(pool._freeLd.ref, pool.ref, nil)
}

func createCommandPool(
	pd *vkPhysicalDevice,
	ld *vkLogicalDevice,
	fb *vkFrameBuffers,
	renderPass vulkan.RenderPass,
	swapChain *vkSwapChain,
	pipeLine *vkPipeline,
	vertexBuffer vulkan.Buffer,
	closer *utils.Closer,
) *vkCommandPool {
	commandPool := &vkCommandPool{
		_freeLd: ld,
	}
	closer.EnqueueFree(commandPool.free)

	// pool
	poolInfo := &vulkan.CommandPoolCreateInfo{
		SType:            vulkan.StructureTypeCommandPoolCreateInfo,
		QueueFamilyIndex: pd.families.graphicsFamilyId,
	}

	var pool vulkan.CommandPool
	vkAssert(
		vulkan.CreateCommandPool(ld.ref, poolInfo, nil, &pool),
		fmt.Errorf("failed create command pool"),
	)
	commandPool.ref = pool

	// command buffers
	allocInfo := &vulkan.CommandBufferAllocateInfo{
		SType:              vulkan.StructureTypeCommandBufferAllocateInfo,
		CommandPool:        pool,
		Level:              vulkan.CommandBufferLevelPrimary,
		CommandBufferCount: uint32(len(fb.buffers)),
	}

	commandBuffers := make([]vulkan.CommandBuffer, len(fb.buffers))

	vkAssert(
		vulkan.AllocateCommandBuffers(ld.ref, allocInfo, commandBuffers),
		fmt.Errorf("failed allocate command buffers"),
	)

	for ind, buffer := range commandBuffers {
		beginInfo := &vulkan.CommandBufferBeginInfo{
			SType: vulkan.StructureTypeCommandBufferBeginInfo,
		}

		vkAssert(
			vulkan.BeginCommandBuffer(buffer, beginInfo),
			fmt.Errorf("failed begin command buffer"),
		)

		renderPassBeginInfo := &vulkan.RenderPassBeginInfo{
			SType:       vulkan.StructureTypeRenderPassBeginInfo,
			RenderPass:  renderPass,
			Framebuffer: fb.buffers[ind],
			RenderArea: vulkan.Rect2D{
				Offset: vulkan.Offset2D{
					X: 0,
					Y: 0,
				},
				Extent: vulkan.Extent2D{
					Width:  swapChain.info.bufferSize.Width,
					Height: swapChain.info.bufferSize.Height,
				},
			},
			ClearValueCount: 1,
			PClearValues:    []vulkan.ClearValue{{0.0, 0.0, 0.0, 1.0}},
		}

		vulkan.CmdBeginRenderPass(buffer, renderPassBeginInfo, vulkan.SubpassContentsInline)

		vulkan.CmdBindPipeline(buffer, vulkan.PipelineBindPointGraphics, pipeLine.ref)
		vulkan.CmdBindVertexBuffers(buffer, 0, 1, []vulkan.Buffer{vertexBuffer}, []vulkan.DeviceSize{0})
		vulkan.CmdDraw(buffer, 3, 1, 0, 0)

		vulkan.CmdEndRenderPass(buffer)

		vkAssert(
			vulkan.EndCommandBuffer(buffer),
			fmt.Errorf("failed end command buffer"),
		)
	}

	commandPool.buffers = commandBuffers
	return commandPool
}