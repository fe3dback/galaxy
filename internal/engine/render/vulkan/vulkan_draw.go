package vulkan

import (
	"fmt"
	"log"
	"time"

	"github.com/vulkan-go/vulkan"
)

func (vk *Vk) Draw() {
	const timeout = time.Second * 10
	imageIndex := uint32(0)
	vkAssert(
		vulkan.AcquireNextImage(vk.logicalDevice.ref, vk.swapChain.ref, uint64(timeout.Nanoseconds()), vk.muxImageAvailable, nil, &imageIndex),
		fmt.Errorf("failed acquire next image for rendering"),
	)

	submitInfo := vulkan.SubmitInfo{
		SType:                vulkan.StructureTypeSubmitInfo,
		WaitSemaphoreCount:   1,
		PWaitSemaphores:      []vulkan.Semaphore{vk.muxImageAvailable},
		PWaitDstStageMask:    []vulkan.PipelineStageFlags{vulkan.PipelineStageFlags(vulkan.PipelineStageColorAttachmentOutputBit)},
		CommandBufferCount:   1,
		PCommandBuffers:      []vulkan.CommandBuffer{vk.commandPool.buffers[imageIndex]},
		SignalSemaphoreCount: 1,
		PSignalSemaphores:    []vulkan.Semaphore{vk.muxPresentAvailable},
	}

	vkAssert(
		vulkan.QueueSubmit(vk.logicalDevice.graphicsQueue, 1, []vulkan.SubmitInfo{submitInfo}, nil),
		fmt.Errorf("failed submit graphics queue"),
	)

	presentInfo := &vulkan.PresentInfo{
		SType:              vulkan.StructureTypePresentInfo,
		WaitSemaphoreCount: 1,
		PWaitSemaphores:    []vulkan.Semaphore{vk.muxPresentAvailable},
		SwapchainCount:     1,
		PSwapchains:        []vulkan.Swapchain{vk.swapChain.ref},
		PImageIndices:      []uint32{imageIndex},
	}

	result := vulkan.QueuePresent(vk.logicalDevice.presetQueue, presentInfo)
	if result != vulkan.Success {
		// todo: error handling?
		log.Println("failed present image from queue")
	}

	vulkan.QueueWaitIdle(vk.logicalDevice.presetQueue)
}

func (vk *Vk) Wait() {
	vulkan.DeviceWaitIdle(vk.logicalDevice.ref)
}
