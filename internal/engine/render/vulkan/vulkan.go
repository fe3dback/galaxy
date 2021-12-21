package vulkan

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	Vk struct {
		instance            *vkInstance
		physicalDevice      *vkPhysicalDevice
		logicalDevice       *vkLogicalDevice
		swapChain           *vkSwapChain
		renderPass          vulkan.RenderPass
		pipeLine            *vkPipeline
		frameBuffers        *vkFrameBuffers
		commandPool         *vkCommandPool
		muxImageAvailable   vulkan.Semaphore
		muxPresentAvailable vulkan.Semaphore
	}

	vkCreateOptions struct {
		closer      *utils.Closer
		window      *glfw.Window
		debugVulkan bool
	}
)

func NewVulkanApi(closer *utils.Closer, window *glfw.Window, debugVulkan bool) *Vk {
	err := vulkan.Init()
	if err != nil {
		panic(fmt.Errorf("failed init vulkan: %w", err))
	}

	opts := vkCreateOptions{
		closer:      closer,
		window:      window,
		debugVulkan: debugVulkan,
	}

	inst := vkCreateInstance(opts)
	inst.surface = inst.vkCreateSurface(opts)

	physicalDevice := inst.vkPickPhysicalDevice()
	logicalDevice := physicalDevice.createLogicalDevice(opts)
	swapChain := vkCreateSwapChain(inst, physicalDevice, logicalDevice, opts)

	vk := &Vk{
		instance:       inst,
		physicalDevice: physicalDevice,
		logicalDevice:  logicalDevice,
		swapChain:      swapChain,
	}

	vk.muxImageAvailable, vk.muxPresentAvailable = vkCreateSemaphores(vk.logicalDevice, opts)
	vk.renderPass = vk.vkCreateRenderPass(opts)
	vk.pipeLine = vk.vkCreatePipeline(opts)
	vk.frameBuffers = vk.vkCreateFrameBuffers(opts)
	vk.commandPool = vk.vkCreateCommandPool(opts)

	return vk
}
