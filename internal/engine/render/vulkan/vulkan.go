package vulkan

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	Vk struct {
		instance              *vkInstance
		physicalDevice        *vkPhysicalDevice
		logicalDevice         *vkLogicalDevice
		swapChain             *vkSwapChain
		swapChainFactory      *vkSwapChainFactory
		swapChainFrameManager *vkSwapChainFrameManager
		shaderManager         *vkShaderManager
		pipeLine              *vkPipeline
		frameBuffers          *vkFrameBuffers
		commandPool           *vkCommandPool
	}
)

func NewVulkanApi(window *glfw.Window, cfg Config, closer *utils.Closer) *Vk {
	err := vulkan.Init()
	if err != nil {
		panic(fmt.Errorf("failed init vulkan: %w", err))
	}

	log.Printf("vk: lib initialized: [%#v]\n", cfg)

	// required ext
	requiredExt := window.GetRequiredInstanceExtensions()

	// init
	inst := createInstance(requiredExt, cfg.debug, closer)
	surface := inst.createSurfaceFromWindow(window, closer)

	// find GPU for usage
	physicalDeviceFinder := newPhysicalDeviceFinder(inst, surface)
	physicalDevice := physicalDeviceFinder.physicalDevicePick()
	logicalDevice := physicalDevice.createLogicalDevice(closer)

	// create render swapChain
	swapChainFactory := newSwapChainFactory(surface, physicalDevice, logicalDevice, createWindowSizeExtractor(window), cfg, closer)
	swapChain := swapChainFactory.createSwapChain()
	swapChainFrameManager := newSwapChainFrameManager(logicalDevice, closer)

	// create pipeline and render staff
	shaderManager := newShaderManager(logicalDevice, closer)
	pipeLineCfg := newPipeLineCfg(logicalDevice, swapChain, closer)
	renderPass := pipeLineCfg.renderPass

	// pipeline (todo: dynamics)
	inputShaders := []vulkan.PipelineShaderStageCreateInfo{
		shaderManager.shaderModule(shaderIDTriangleVert).stageInfo,
		shaderManager.shaderModule(shaderIDTriangleFrag).stageInfo,
	}
	// todo: inputRenderPass is shader params?
	pipeLine := createPipeline(pipeLineCfg, logicalDevice, swapChain, inputShaders, closer)
	frameBuffers := createFrameBuffers(swapChain, logicalDevice, renderPass, closer)
	commandPool := createCommandPool(physicalDevice, logicalDevice, frameBuffers, renderPass, swapChain, pipeLine, closer)

	return &Vk{
		instance:              inst,
		physicalDevice:        physicalDevice,
		logicalDevice:         logicalDevice,
		swapChain:             swapChain,
		swapChainFactory:      swapChainFactory,
		swapChainFrameManager: swapChainFrameManager,
		shaderManager:         shaderManager,
		pipeLine:              pipeLine,
		frameBuffers:          frameBuffers,
		commandPool:           commandPool,
	}
}

func createWindowSizeExtractor(window *glfw.Window) vkScreenSizeExtractor {
	return func() (w, h uint32) {
		width, height := window.GetFramebufferSize()
		return uint32(width), uint32(height)
	}
}
