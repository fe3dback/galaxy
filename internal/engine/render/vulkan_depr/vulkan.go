package vulkan_depr

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine/event"
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

		closer              *utils.Closer
		isResizing          bool
		isDrawAvailable     bool
		windowSizeExtractor windowSizeExtractor

		// frame
		frameQueue []shaderProgram
	}

	windowSizeExtractor = func() (width, height uint32)
)

func NewVulkanApi(window *glfw.Window, dispatcher *event.Dispatcher, cfg Config, closer *utils.Closer) *Vk {
	err := vulkan.Init()
	if err != nil {
		panic(fmt.Errorf("failed init vulkan: %w", err))
	}

	log.Printf("vk: lib initialized: [%#v]\n", cfg)

	// required ext
	requiredExt := window.GetRequiredInstanceExtensions()

	// todo: debug callbacks

	// init
	vk := &Vk{}
	inst := createInstance(requiredExt, cfg.debug, closer)
	surface := inst.createSurfaceFromWindow(window, closer)

	// find GPU for usage
	physicalDeviceFinder := newPhysicalDeviceFinder(inst, surface)
	physicalDevice := physicalDeviceFinder.physicalDevicePick()
	logicalDevice := physicalDevice.createLogicalDevice(closer)
	shaderManager := newShaderManager(logicalDevice, closer)

	// create render swapChain
	swapChainFactory := newSwapChainFactory(surface, physicalDevice, logicalDevice, createWindowSizeExtractor(window), cfg, closer)
	swapChain, pipeLine, frameBuffers, commandPool :=
		swapChainFactory.createAllPipeline(physicalDevice, logicalDevice, shaderManager, closer)

	// render
	onOutOfDate := func() {
		vk.rebuildGraphicsPipeline()
	}
	swapChainFrameManager := newSwapChainFrameManager(logicalDevice, onOutOfDate, closer)
	swapChainFrameManager.setSwapChain(swapChain)
	swapChainFrameManager.setCommandPool(commandPool)

	// assemble
	vk.instance = inst
	vk.physicalDevice = physicalDevice
	vk.logicalDevice = logicalDevice
	vk.swapChain = swapChain
	vk.swapChainFactory = swapChainFactory
	vk.swapChainFrameManager = swapChainFrameManager
	vk.shaderManager = shaderManager
	vk.pipeLine = pipeLine
	vk.frameBuffers = frameBuffers
	vk.commandPool = commandPool
	vk.closer = closer
	vk.windowSizeExtractor = createWindowSizeExtractor(window)
	vk.isDrawAvailable = true

	// subscribe to system events
	dispatcher.OnWindowResized(func(windowResizedEvent event.WindowResizedEvent) error {
		vk.onWindowResized(windowResizedEvent.NewWidth, windowResizedEvent.NewHeight)
		return nil
	})

	return vk
}

func createWindowSizeExtractor(window *glfw.Window) windowSizeExtractor {
	return func() (w, h uint32) {
		width, height := window.GetFramebufferSize()
		return uint32(width), uint32(height)
	}
}
