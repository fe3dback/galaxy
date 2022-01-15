package vulkan

import "github.com/vulkan-go/vulkan"

type (
	Vk struct {
		inst         *vkInstance
		surface      *vkSurface
		pd           *vkPhysicalDevice
		ld           *vkLogicalDevice
		commandPool  *vkCommandPool
		frameManager *vkFrameManager

		// render variables
		currentFrameImageID            uint32
		currentFrameAvailableForRender bool
	}

	// --------------------------------------
	// SYSTEM
	// --------------------------------------

	vkInstance struct {
		ref vulkan.Instance
	}

	vkSurface struct {
		inst *vkInstance
		ref  vulkan.Surface
	}

	// --------------------------------------
	// DEVICES
	// --------------------------------------

	vkPhysicalDeviceFinder struct {
		surface *vkSurface
		inst    *vkInstance
	}

	vkPhysicalDevice struct {
		ref          vulkan.PhysicalDevice
		props        vulkan.PhysicalDeviceProperties
		features     vulkan.PhysicalDeviceFeatures
		extensions   []vulkan.ExtensionProperties
		families     vkPhysicalDeviceFamilies
		surfaceProps vkPhysicalDeviceSurfaceProps
	}

	vkPhysicalDeviceFamilies struct {
		graphicsFamilyId uint32
		presentFamilyId  uint32
		supportGraphics  bool
		supportPresent   bool
	}

	vkPhysicalDeviceSurfaceProps struct {
		capabilities   vulkan.SurfaceCapabilities
		formats        []vulkan.SurfaceFormat
		presentModes   []vulkan.PresentMode
		surfaceSupport bool
	}

	vkLogicalDevice struct {
		ref           vulkan.Device
		queueGraphics vulkan.Queue
		queuePresent  vulkan.Queue
	}

	// --------------------------------------
	// Render Requirements
	// --------------------------------------

	vkCommandPool struct {
		ref     vulkan.CommandPool
		buffers []vulkan.CommandBuffer
		ld      *vkLogicalDevice
	}

	frameID = uint8
	imageID = uint32

	vkFrameManager struct {
		maxFrames         uint8
		currentFrameID    frameID
		currentImageID    imageID
		presentFailsCount int

		muxRenderAvailable  map[frameID]vulkan.Semaphore
		muxPresentAvailable map[frameID]vulkan.Semaphore
		fence               map[frameID]vulkan.Fence
		imagesInFlight      map[imageID]vulkan.Fence
		onSwapOutOfDate     func()

		ld *vkLogicalDevice
	}

	// --------------------------------------
	// Graphics
	// --------------------------------------

	vkSwapChain struct {
		ref          vulkan.Swapchain
		surfaceProps vkPhysicalDeviceSurfaceProps
		images       []vulkan.Image
		imagesView   []vulkan.ImageView
		info         vkSwapChainInfo
	}

	vkSwapChainInfo struct {
		imageFormat     vulkan.Format
		imageColorSpace vulkan.ColorSpace
		bufferSize      vulkan.Extent2D
		presentMode     vulkan.PresentMode
	}
)
