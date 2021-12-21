package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

const swapChainMaxFrames = 2

type (
	frameID = uint8

	vkSwapChainFrameManager struct {
		maxFrames    uint8
		currentFrame frameID

		muxRenderAvailable  map[frameID]vulkan.Semaphore
		muxPresentAvailable map[frameID]vulkan.Semaphore
	}
)

func newSwapChainFrameManager(ld *vkLogicalDevice, closer *utils.Closer) *vkSwapChainFrameManager {
	frameManager := &vkSwapChainFrameManager{
		maxFrames:           swapChainMaxFrames,
		currentFrame:        0,
		muxRenderAvailable:  make(map[frameID]vulkan.Semaphore),
		muxPresentAvailable: make(map[frameID]vulkan.Semaphore),
	}

	createInfo := &vulkan.SemaphoreCreateInfo{
		SType: vulkan.StructureTypeSemaphoreCreateInfo,
	}

	for i := frameID(0); i < swapChainMaxFrames; i++ {
		var muxRender vulkan.Semaphore
		var muxPresent vulkan.Semaphore

		vkAssert(
			vulkan.CreateSemaphore(ld.ref, createInfo, nil, &muxRender),
			fmt.Errorf("failed create image available mux"),
		)
		vkAssert(
			vulkan.CreateSemaphore(ld.ref, createInfo, nil, &muxPresent),
			fmt.Errorf("failed create present available mux"),
		)

		frameManager.muxRenderAvailable[i] = muxRender
		frameManager.muxPresentAvailable[i] = muxPresent
	}

	closer.EnqueueFree(func() {
		for i := frameID(0); i < frameManager.maxFrames; i++ {
			vulkan.DestroySemaphore(ld.ref, frameManager.muxPresentAvailable[i], nil)
			vulkan.DestroySemaphore(ld.ref, frameManager.muxRenderAvailable[i], nil)
		}
	})

	return frameManager
}
