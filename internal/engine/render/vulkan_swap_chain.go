package render

import (
	"fmt"
	"math"

	"github.com/vulkan-go/vulkan"
)

type (
	vkSwapChain struct {
		ref        vulkan.Swapchain
		surface    *vkPhysicalDeviceSurface
		images     []vulkan.Image
		imagesView []vulkan.ImageView
	}
)

func vkCreateSwapChain(inst *vkInstance, pd *vkPhysicalDevice, ld *vkLogicalDevice, opts vkCreateOptions) *vkSwapChain {
	format := pd.surface.richColorSpaceFormat()
	if format == nil {
		panic(fmt.Errorf("device '%s' not support rich color space format", pd.props.DeviceName))
	}

	families := []uint32{queueFamilyGraphics, queueFamilyPresent}
	sharingMode := vulkan.SharingModeExclusive
	if len(families) > 1 {
		sharingMode = vulkan.SharingModeConcurrent
	}

	swapChainCreateInfo := &vulkan.SwapchainCreateInfo{
		SType:                 vulkan.StructureTypeSwapchainCreateInfo,
		Surface:               inst.surface.ref,
		MinImageCount:         pd.surface.imageBuffersCount(),
		ImageFormat:           (*format).Format,
		ImageColorSpace:       (*format).ColorSpace,
		ImageExtent:           pd.surface.chooseSwapExtent(opts),
		ImageArrayLayers:      1,
		ImageUsage:            vulkan.ImageUsageFlags(vulkan.ImageUsageColorAttachmentBit),
		ImageSharingMode:      sharingMode,
		QueueFamilyIndexCount: uint32(len(families)),
		PQueueFamilyIndices:   families,
		PreTransform:          pd.surface.capabilities.CurrentTransform,
		CompositeAlpha:        vulkan.CompositeAlphaOpaqueBit,
		PresentMode:           pd.surface.bestPresentMode(),
		Clipped:               vulkan.True,
	}

	var swapChain vulkan.Swapchain
	vkAssert(
		vulkan.CreateSwapchain(ld.ref, swapChainCreateInfo, nil, &swapChain),
		fmt.Errorf("failed create swapChain"),
	)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroySwapchain(ld.ref, swapChain, nil)
	})

	imagesCount := uint32(0)
	vkAssert(
		vulkan.GetSwapchainImages(ld.ref, swapChain, &imagesCount, nil),
		fmt.Errorf("failed fetch swapChain images"),
	)
	if imagesCount == 0 {
		panic(fmt.Errorf("swapchain should have at least 1 image buffer"))
	}

	images := make([]vulkan.Image, imagesCount)
	vkAssert(
		vulkan.GetSwapchainImages(ld.ref, swapChain, &imagesCount, images),
		fmt.Errorf("failed fetch swapChain images"),
	)

	return &vkSwapChain{
		ref:        swapChain,
		surface:    pd.surface,
		images:     images,
		imagesView: vkCreateSwapChainImagesView(images, pd, ld, opts),
	}
}

func vkCreateSwapChainImagesView(images []vulkan.Image, pd *vkPhysicalDevice, ld *vkLogicalDevice, opts vkCreateOptions) []vulkan.ImageView {
	views := make([]vulkan.ImageView, 0, len(images))

	for _, image := range images {
		views = append(views, vkCreateSwapChainImageView(image, pd, ld, opts))
	}

	return views
}

func vkCreateSwapChainImageView(image vulkan.Image, pd *vkPhysicalDevice, ld *vkLogicalDevice, opts vkCreateOptions) vulkan.ImageView {
	createInfo := &vulkan.ImageViewCreateInfo{
		SType:    vulkan.StructureTypeImageViewCreateInfo,
		Image:    image,
		ViewType: vulkan.ImageViewType2d,
		Format:   (*pd.surface.richColorSpaceFormat()).Format,
		Components: vulkan.ComponentMapping{
			R: vulkan.ComponentSwizzleIdentity,
			G: vulkan.ComponentSwizzleIdentity,
			B: vulkan.ComponentSwizzleIdentity,
			A: vulkan.ComponentSwizzleIdentity,
		},
		SubresourceRange: vulkan.ImageSubresourceRange{
			AspectMask:     vulkan.ImageAspectFlags(vulkan.ImageAspectColorBit),
			BaseMipLevel:   0,
			LevelCount:     1,
			BaseArrayLayer: 0,
			LayerCount:     1,
		},
	}

	var imageView vulkan.ImageView
	vkAssert(
		vulkan.CreateImageView(ld.ref, createInfo, nil, &imageView),
		fmt.Errorf("failed create image view"),
	)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroyImageView(ld.ref, imageView, nil)
	})

	return imageView
}

func (ds *vkPhysicalDeviceSurface) richColorSpaceFormat() *vulkan.SurfaceFormat {
	for _, surfaceFormat := range ds.formats {
		if surfaceFormat.Format != vulkan.FormatB8g8r8a8Srgb {
			continue
		}

		if surfaceFormat.ColorSpace != vulkan.ColorSpaceSrgbNonlinear {
			continue
		}

		return &surfaceFormat
	}

	return nil
}

func (ds *vkPhysicalDeviceSurface) bestPresentMode() vulkan.PresentMode {
	for _, mode := range ds.presentModes {
		// rank 1:  Instead of blocking the application when the queue is full,
		// the images that are already queued are simply replaced with the newer ones.
		// This mode can be used to render frames as fast as possible while still avoiding tearing,
		// resulting in fewer latency issues than standard vertical sync.
		// This is commonly known as "triple buffering", although the existence of three buffers alone
		// does not necessarily mean that the framerate is unlocked.
		if mode == vulkan.PresentModeMailbox {
			return mode
		}

		// rank 2: The swap chain is a queue where the display takes an image from the
		// front of the queue when the display is refreshed and the program inserts
		// rendered images at the back of the queue. If the queue is full then the program has to wait.
		// This is most similar to vertical sync as found in modern games.
		// The moment that the display is refreshed is known as "vertical blank".
		if mode == vulkan.PresentModeFifo {
			return mode
		}

		// rank 3: any available (IMMEDIATE_KHR or FIFO_RELAXED_KHR)
		return mode
	}

	panic(fmt.Errorf("GPU not support any present mode"))
}

func (ds *vkPhysicalDeviceSurface) chooseSwapExtent(opts vkCreateOptions) vulkan.Extent2D {
	calculatedMax := uint32((math.MaxInt32 * 2) + 1)

	// Vulkan tells us to match the resolution of the window by setting the width and height in the currentExtent member.
	// However, some window managers do allow us to differ here and this is indicated by setting the width and height
	// in currentExtent to a special value: the maximum value of uint32_t.
	//
	// In that case we'll pick the resolution that best matches the window within the minImageExtent and maxImageExtent bounds.
	// But we must specify the resolution in the correct unit.

	if ds.capabilities.CurrentExtent.Width != calculatedMax {
		return ds.capabilities.CurrentExtent
	}

	width, height := opts.window.GetFramebufferSize()
	actualExtent := vulkan.Extent2D{
		Width:  uint32(width),
		Height: uint32(height),
	}

	actualExtent.Width = vkClampUint(
		actualExtent.Width,
		ds.capabilities.MinImageExtent.Width,
		ds.capabilities.MaxImageExtent.Width,
	)
	actualExtent.Height = vkClampUint(
		actualExtent.Height,
		ds.capabilities.MinImageExtent.Height,
		ds.capabilities.MaxImageExtent.Height,
	)

	return actualExtent
}

func (ds *vkPhysicalDeviceSurface) imageBuffersCount() uint32 {
	optimalCount := ds.capabilities.MinImageCount + 1

	if ds.capabilities.MaxImageCount == 0 {
		// no maximum
		return optimalCount
	}

	if optimalCount > ds.capabilities.MaxImageCount {
		// clamp to max
		optimalCount = ds.capabilities.MaxImageCount
	}

	return optimalCount
}
