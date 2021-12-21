package vulkan

import (
	"github.com/vulkan-go/vulkan"
)

type (
	vkSwapChain struct {
		ref          vulkan.Swapchain
		surfaceProps vkPhysicalDeviceSurfaceProps
		images       []vulkan.Image
		imagesView   []vulkan.ImageView
		info         vkSwapChainInfo

		_free   bool
		_freeLd *vkLogicalDevice
	}

	vkSwapChainInfo struct {
		imageFormat     vulkan.Format
		imageColorSpace vulkan.ColorSpace
		bufferSize      vulkan.Extent2D
		presentMode     vulkan.PresentMode
	}
)

func (sc *vkSwapChain) free() {
	if sc._free {
		return
	}
	sc._free = true

	for _, view := range sc.imagesView {
		vulkan.DestroyImageView(sc._freeLd.ref, view, nil)
	}

	vulkan.DestroySwapchain(sc._freeLd.ref, sc.ref, nil)
}

func (sc *vkSwapChain) viewport() vulkan.Viewport {
	return vulkan.Viewport{
		X:        0,
		Y:        0,
		Width:    float32(sc.info.bufferSize.Width),
		Height:   float32(sc.info.bufferSize.Height),
		MinDepth: 0.0,
		MaxDepth: 1.0,
	}
}

func (sc *vkSwapChain) scissor() vulkan.Rect2D {
	return vulkan.Rect2D{
		Offset: vulkan.Offset2D{
			X: 0,
			Y: 0,
		},
		Extent: sc.info.bufferSize,
	}
}
