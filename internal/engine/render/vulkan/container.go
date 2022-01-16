package vulkan

import (
	"fmt"
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	container struct {
		// dependencies
		window     *glfw.Window
		dispatcher *event.Dispatcher
		cfg        *Config
		closer     *utils.Closer

		// internal
		vkRenderPassHandlesLazyCache map[renderPassType]vulkan.RenderPass
		vkPipelineHandlesLazyCache   map[shaderModuleID]vulkan.Pipeline

		// vk handle wrappers
		vk                *Vk
		vkInstance        *vkInstance
		vkSurface         *vkSurface
		vkPhysicalDevice  *vkPhysicalDevice
		vkLogicalDevice   *vkLogicalDevice
		vkCommandPool     *vkCommandPool
		vkFrameManager    *vkFrameManager
		vkSwapChain       *vkSwapChain
		vkFrameBuffers    *vkFrameBuffers
		vkShaderManager   *vkShaderManager
		vkPipelineManager *vkPipelineManager
		vkPipelineLayout  vulkan.PipelineLayout
	}
)

func newContainer(window *glfw.Window, dispatcher *event.Dispatcher, cfg *Config, closer *utils.Closer) *container {
	return &container{
		window:     window,
		dispatcher: dispatcher,
		cfg:        cfg,
		closer:     closer,

		// internal
		vkRenderPassHandlesLazyCache: map[renderPassType]vulkan.RenderPass{},
		vkPipelineHandlesLazyCache:   map[shaderModuleID]vulkan.Pipeline{},
	}
}

func (c *container) renderer() *Vk {
	if c.vk != nil {
		return c.vk
	}

	err := vulkan.Init()
	if err != nil {
		panic(fmt.Errorf("failed init vulkan: %w", err))
	}

	log.Printf("Vk: lib initialized: [%#v]\n", c.cfg)

	// main
	c.vk = &Vk{}
	c.vk.container = c
	c.vk.inst = c.provideVkInstance()
	c.vk.surface = c.provideVkSurface()
	c.vk.pd = c.provideVkPhysicalDevice()
	c.vk.ld = c.provideVkLogicalDevice()

	// render utils
	c.vk.commandPool = c.provideVkCommandPool()
	c.vk.frameManager = c.provideFrameManager(c.vk.rebuildGraphicsPipeline)
	c.vk.swapChain = c.provideSwapChain()
	c.vk.frameBuffers = c.provideFrameBuffers()
	c.vk.shaderManager = c.provideShaderManager()
	c.vk.pipelineManager = c.providePipelineManager()
	c.vk.pipelineLayout = c.providePipelineLayout()

	for shader, pipelineFactory := range buildInShaders {
		c.vk.shaderManager.preloadShader(shader)
		c.vkPipelineManager.preloadPipelineFor(shader, pipelineFactory(c, shader))
	}

	// render

	return c.vk
}

func (c *container) provideVkInstance() *vkInstance {
	if c.vkInstance != nil {
		return c.vkInstance
	}

	// required ext
	requiredExt := c.window.GetRequiredInstanceExtensions()

	// todo: debug callbacks

	// init
	c.vkInstance = newVkInstance(requiredExt, c.cfg.debug)
	return c.vkInstance
}

func (c *container) provideVkSurface() *vkSurface {
	if c.vkSurface != nil {
		return c.vkSurface
	}

	c.vkSurface = newSurfaceFromWindow(
		c.provideVkInstance(),
		c.window,
	)
	return c.vkSurface
}

func (c *container) provideVkPhysicalDevice() *vkPhysicalDevice {
	if c.vkPhysicalDevice != nil {
		return c.vkPhysicalDevice
	}

	finder := newPhysicalDeviceFinder(
		c.provideVkInstance(),
		c.provideVkSurface(),
	)

	c.vkPhysicalDevice = finder.physicalDevicePick()
	return c.vkPhysicalDevice
}

func (c *container) provideVkLogicalDevice() *vkLogicalDevice {
	if c.vkLogicalDevice != nil {
		return c.vkLogicalDevice
	}

	c.vkLogicalDevice = newLogicalDevice(
		c.provideVkPhysicalDevice(),
	)
	return c.vkLogicalDevice
}

func (c *container) provideVkCommandPool() *vkCommandPool {
	if c.vkCommandPool != nil {
		return c.vkCommandPool
	}

	c.vkCommandPool = newCommandPool(
		c.provideVkPhysicalDevice(),
		c.provideVkLogicalDevice(),
	)
	return c.vkCommandPool
}

func (c *container) provideFrameManager(onSwapOutOfDate func()) *vkFrameManager {
	if c.vkFrameManager != nil {
		return c.vkFrameManager
	}

	c.vkFrameManager = newFrameManager(
		c.provideVkLogicalDevice(),
		c.provideVkPhysicalDevice(),
		onSwapOutOfDate,
	)
	return c.vkFrameManager
}

func (c *container) provideSwapChain() *vkSwapChain {
	if c.vkSwapChain != nil {
		return c.vkSwapChain
	}

	wWidth, wHeight := c.window.GetFramebufferSize()

	c.vkSwapChain = newSwapChain(
		uint32(wWidth), uint32(wHeight),
		c.provideVkPhysicalDevice(),
		c.provideVkLogicalDevice(),
		c.provideVkSurface(),
		c.cfg,
	)
	return c.vkSwapChain
}

func (c *container) provideFrameBuffers() *vkFrameBuffers {
	if c.vkFrameBuffers != nil {
		return c.vkFrameBuffers
	}

	c.vkFrameBuffers = newFrameBuffers(
		c.provideVkLogicalDevice(),
		c.provideSwapChain(),
		c.defaultRenderPass(),
	)
	return c.vkFrameBuffers
}

func (c *container) provideShaderManager() *vkShaderManager {
	if c.vkShaderManager != nil {
		return c.vkShaderManager
	}

	c.vkShaderManager = newShaderManager(
		c.provideVkLogicalDevice(),
	)
	return c.vkShaderManager
}

func (c *container) providePipelineManager() *vkPipelineManager {
	if c.vkPipelineManager != nil {
		return c.vkPipelineManager
	}

	c.vkPipelineManager = newPipelineManager(
		c.provideVkLogicalDevice(),
	)
	return c.vkPipelineManager
}

func (c *container) providePipelineLayout() vulkan.PipelineLayout {
	if c.vkPipelineLayout != nil {
		return c.vkPipelineLayout
	}

	c.vkPipelineLayout = newPipeLineLayout(
		c.provideVkLogicalDevice(),
	)
	return c.vkPipelineLayout
}
