package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkCommandPool struct {
		pool    vulkan.CommandPool
		buffers []vulkan.CommandBuffer
	}
)

func (vk *Vk) vkCreateCommandPool(opts vkCreateOptions) *vkCommandPool {
	// pool
	poolInfo := &vulkan.CommandPoolCreateInfo{
		SType:            vulkan.StructureTypeCommandPoolCreateInfo,
		QueueFamilyIndex: vk.physicalDevice.family.graphicsFamilyId,
	}

	var pool vulkan.CommandPool
	vkAssert(
		vulkan.CreateCommandPool(vk.logicalDevice.ref, poolInfo, nil, &pool),
		fmt.Errorf("failed create command pool"),
	)

	opts.closer.EnqueueFree(func() {
		vulkan.DestroyCommandPool(vk.logicalDevice.ref, pool, nil)
	})

	// command buffers
	allocInfo := &vulkan.CommandBufferAllocateInfo{
		SType:              vulkan.StructureTypeCommandBufferAllocateInfo,
		CommandPool:        pool,
		Level:              vulkan.CommandBufferLevelPrimary,
		CommandBufferCount: uint32(len(vk.frameBuffers.buffers)),
	}

	commandBuffers := make([]vulkan.CommandBuffer, len(vk.frameBuffers.buffers))

	vkAssert(
		vulkan.AllocateCommandBuffers(vk.logicalDevice.ref, allocInfo, commandBuffers),
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
			RenderPass:  vk.renderPass,
			Framebuffer: vk.frameBuffers.buffers[ind],
			RenderArea: vulkan.Rect2D{
				Offset: vulkan.Offset2D{
					X: 0,
					Y: 0,
				},
				Extent: vulkan.Extent2D{
					Width:  vk.swapChain.info.bufferSize.Width,
					Height: vk.swapChain.info.bufferSize.Height,
				},
			},
			ClearValueCount: 1,
			PClearValues:    []vulkan.ClearValue{{0.0, 0.0, 0.0, 1.0}},
		}

		vulkan.CmdBeginRenderPass(buffer, renderPassBeginInfo, vulkan.SubpassContentsInline)
		vulkan.CmdBindPipeline(buffer, vulkan.PipelineBindPointGraphics, vk.pipeLine.ref)
		vulkan.CmdDraw(buffer, 3, 1, 0, 0)
		vulkan.CmdEndRenderPass(buffer)

		vkAssert(
			vulkan.EndCommandBuffer(buffer),
			fmt.Errorf("failed end command buffer"),
		)
	}

	return &vkCommandPool{
		pool:    pool,
		buffers: commandBuffers,
	}
}
