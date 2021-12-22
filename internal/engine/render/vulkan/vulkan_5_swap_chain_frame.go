package vulkan

import (
	"fmt"
	"time"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

const swapChainMaxFrames = 2
const swapChainTimeout = time.Second * 10 // GPU timeout for render. After this, app will be crashed
const maxPresetFails = 100                // How many frames can be failed continuously before crash

type (
	frameID = uint8
	imageID = uint8

	vkSwapChainFrameManager struct {
		maxFrames         uint8
		currentFrame      frameID
		presentFailsCount int

		muxRenderAvailable  map[frameID]vulkan.Semaphore
		muxPresentAvailable map[frameID]vulkan.Semaphore
		fence               map[frameID]vulkan.Fence
		imagesInFlight      map[imageID]vulkan.Fence

		ld              *vkLogicalDevice
		swapChain       *vkSwapChain
		commandPool     *vkCommandPool
		onSwapOutOfDate vkOnSwapOutOfDate
	}

	vkOnSwapOutOfDate = func()
)

func newSwapChainFrameManager(ld *vkLogicalDevice, onSwapOutOfDate vkOnSwapOutOfDate, closer *utils.Closer) *vkSwapChainFrameManager {
	frameManager := &vkSwapChainFrameManager{
		maxFrames:           swapChainMaxFrames,
		currentFrame:        0,
		muxRenderAvailable:  make(map[frameID]vulkan.Semaphore),
		muxPresentAvailable: make(map[frameID]vulkan.Semaphore),
		fence:               make(map[frameID]vulkan.Fence),
		imagesInFlight:      make(map[imageID]vulkan.Fence),
		ld:                  ld,
		onSwapOutOfDate:     onSwapOutOfDate,
	}

	muxCreateInfo := &vulkan.SemaphoreCreateInfo{
		SType: vulkan.StructureTypeSemaphoreCreateInfo,
	}
	fenceCreateInfo := &vulkan.FenceCreateInfo{
		SType: vulkan.StructureTypeFenceCreateInfo,
		Flags: vulkan.FenceCreateFlags(vulkan.FenceCreateSignaledBit),
	}

	for i := frameID(0); i < swapChainMaxFrames; i++ {
		var muxRender vulkan.Semaphore
		var muxPresent vulkan.Semaphore
		var fence vulkan.Fence

		vkAssert(
			vulkan.CreateSemaphore(ld.ref, muxCreateInfo, nil, &muxRender),
			fmt.Errorf("failed create image available mux"),
		)
		vkAssert(
			vulkan.CreateSemaphore(ld.ref, muxCreateInfo, nil, &muxPresent),
			fmt.Errorf("failed create present available mux"),
		)
		vkAssert(
			vulkan.CreateFence(ld.ref, fenceCreateInfo, nil, &fence),
			fmt.Errorf("failed create present available mux"),
		)

		frameManager.muxRenderAvailable[i] = muxRender
		frameManager.muxPresentAvailable[i] = muxPresent
		frameManager.fence[i] = fence
	}

	closer.EnqueueFree(func() {
		for i := frameID(0); i < frameManager.maxFrames; i++ {
			vulkan.DestroyFence(ld.ref, frameManager.fence[i], nil)
			vulkan.DestroySemaphore(ld.ref, frameManager.muxPresentAvailable[i], nil)
			vulkan.DestroySemaphore(ld.ref, frameManager.muxRenderAvailable[i], nil)
		}
	})

	return frameManager
}

func (m *vkSwapChainFrameManager) setSwapChain(swapChain *vkSwapChain) {
	m.swapChain = swapChain
}

func (m *vkSwapChainFrameManager) setCommandPool(commandPool *vkCommandPool) {
	m.commandPool = commandPool
}

func (m *vkSwapChainFrameManager) drawFrame() {
	timeout := uint64(swapChainTimeout.Nanoseconds())

	muxRender := m.muxRenderAvailable[m.currentFrame]
	muxPresent := m.muxPresentAvailable[m.currentFrame]
	fence := m.fence[m.currentFrame]

	vulkan.WaitForFences(m.ld.ref, 1, []vulkan.Fence{fence}, vulkan.True, timeout)

	imageIndex := uint32(0)
	result := vulkan.AcquireNextImage(m.ld.ref, m.swapChain.ref, timeout, muxRender, nil, &imageIndex)
	if result == vulkan.ErrorOutOfDate || result == vulkan.Suboptimal {
		// buffer size changes (window rebuildGraphicsPipeline, minimize, etc..)
		// and not more valid
		m.onSwapOutOfDate()
		return
	}

	vkAssert(result, fmt.Errorf("failed acquire next image for rendering"))

	if imageFence, inFlight := m.imagesInFlight[imageID(imageIndex)]; inFlight {
		vulkan.WaitForFences(m.ld.ref, 1, []vulkan.Fence{imageFence}, vulkan.True, timeout)
	}
	m.imagesInFlight[imageID(imageIndex)] = fence

	submitInfo := vulkan.SubmitInfo{
		SType:                vulkan.StructureTypeSubmitInfo,
		WaitSemaphoreCount:   1,
		PWaitSemaphores:      []vulkan.Semaphore{muxRender},
		PWaitDstStageMask:    []vulkan.PipelineStageFlags{vulkan.PipelineStageFlags(vulkan.PipelineStageColorAttachmentOutputBit)},
		CommandBufferCount:   1,
		PCommandBuffers:      []vulkan.CommandBuffer{m.commandPool.buffers[imageIndex]},
		SignalSemaphoreCount: 1,
		PSignalSemaphores:    []vulkan.Semaphore{muxPresent},
	}

	vulkan.ResetFences(m.ld.ref, 1, []vulkan.Fence{fence})

	vkAssert(
		vulkan.QueueSubmit(m.ld.queueGraphics, 1, []vulkan.SubmitInfo{submitInfo}, fence),
		fmt.Errorf("failed submit graphics queue"),
	)

	presentInfo := &vulkan.PresentInfo{
		SType:              vulkan.StructureTypePresentInfo,
		WaitSemaphoreCount: 1,
		PWaitSemaphores:    []vulkan.Semaphore{muxPresent},
		SwapchainCount:     1,
		PSwapchains:        []vulkan.Swapchain{m.swapChain.ref},
		PImageIndices:      []uint32{imageIndex},
	}

	result = vulkan.QueuePresent(m.ld.queuePresent, presentInfo)
	if result != vulkan.Success {
		m.presentFailsCount++
	} else {
		m.presentFailsCount = 0
	}

	if m.presentFailsCount > maxPresetFails {
		panic(fmt.Errorf("failed present rendered graphics into screen, more that %d times", maxPresetFails))
	}

	vulkan.QueueWaitIdle(m.ld.queuePresent)

	// done
	m.currentFrame = (m.currentFrame + 1) % m.maxFrames
}
