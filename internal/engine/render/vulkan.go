package render

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vk struct {
		instance       *vkInstance
		surface        *vkSurface
		physicalDevice *vkPhysicalDevice
		logicalDevice  *vkLogicalDevice
		swapChain      *vkSwapChain
	}

	vkCreateOptions struct {
		closer      *utils.Closer
		window      *glfw.Window
		debugVulkan bool
	}
)

func newVulkanApi(opts vkCreateOptions) *vk {
	err := vulkan.Init()
	if err != nil {
		panic(fmt.Errorf("failed init vulkan: %w", err))
	}

	inst := vkCreateInstance(opts)
	inst.surface = inst.vkCreateSurface(opts)

	physicalDevice := inst.vkPickPhysicalDevice()
	logicalDevice := physicalDevice.createLogicalDevice(opts)
	swapChain := vkCreateSwapChain(inst, physicalDevice, logicalDevice, opts)

	vk := &vk{
		instance:       inst,
		surface:        inst.surface,
		physicalDevice: physicalDevice,
		logicalDevice:  logicalDevice,
		swapChain:      swapChain,
	}

	renderPass := vk.vkCreateRenderPass(opts)
	pipeline := vk.vkCreatePipeline(opts, renderPass)
	_ = pipeline

	return vk
}
