package vulkan

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkSwapChainFactory struct {
		surface               *vkSurface
		pd                    *vkPhysicalDevice
		ld                    *vkLogicalDevice
		vkScreenSizeExtractor vkScreenSizeExtractor
		closer                *utils.Closer
		cfg                   Config
	}

	vkScreenSizeExtractor = func() (width, height uint32)
)

func newSwapChainFactory(
	surface *vkSurface,
	pd *vkPhysicalDevice,
	ld *vkLogicalDevice,
	vkScreenSizeExtractor vkScreenSizeExtractor,
	cfg Config,
	closer *utils.Closer,
) *vkSwapChainFactory {
	return &vkSwapChainFactory{
		surface:               surface,
		pd:                    pd,
		ld:                    ld,
		vkScreenSizeExtractor: vkScreenSizeExtractor,
		cfg:                   cfg,
		closer:                closer,
	}
}

func (f *vkSwapChainFactory) createSwapChain() *vkSwapChain {
	vkSwapChain := &vkSwapChain{
		_freeLd:      f.ld,
		surfaceProps: f.pd.surfaceProps,
	}
	f.closer.EnqueueFree(vkSwapChain.free)

	// assemble info
	wWidth, wHeight := f.vkScreenSizeExtractor()
	vkSwapChain.info = vkSwapChainInfo{
		imageFormat:     (*f.pd.surfaceProps.richColorSpaceFormat()).Format,
		imageColorSpace: (*f.pd.surfaceProps.richColorSpaceFormat()).ColorSpace,
		bufferSize:      f.pd.surfaceProps.chooseSwapExtent(wWidth, wHeight),
		presentMode:     f.pd.surfaceProps.bestPresentMode(f.cfg.vSync),
	}

	uniqFamilies := f.pd.families.uniqueIDs()
	sharingMode := vulkan.SharingModeExclusive
	if len(uniqFamilies) > 1 {
		sharingMode = vulkan.SharingModeConcurrent
	}

	// assemble create request
	swapChainCreateInfo := &vulkan.SwapchainCreateInfo{
		SType:                 vulkan.StructureTypeSwapchainCreateInfo,
		Surface:               f.surface.ref,
		MinImageCount:         f.pd.surfaceProps.imageBuffersCount(),
		ImageFormat:           vkSwapChain.info.imageFormat,
		ImageColorSpace:       vkSwapChain.info.imageColorSpace,
		ImageExtent:           vkSwapChain.info.bufferSize,
		ImageArrayLayers:      1,
		ImageUsage:            vulkan.ImageUsageFlags(vulkan.ImageUsageColorAttachmentBit),
		ImageSharingMode:      sharingMode,
		QueueFamilyIndexCount: uint32(len(uniqFamilies)),
		PQueueFamilyIndices:   uniqFamilies,
		PreTransform:          f.pd.surfaceProps.capabilities.CurrentTransform,
		CompositeAlpha:        vulkan.CompositeAlphaOpaqueBit,
		PresentMode:           vkSwapChain.info.presentMode,
		Clipped:               vulkan.True,
	}

	// allocate swapChain
	var swapChain vulkan.Swapchain
	vkAssert(
		vulkan.CreateSwapchain(f.ld.ref, swapChainCreateInfo, nil, &swapChain),
		fmt.Errorf("failed create swapChain"),
	)
	vkSwapChain.ref = swapChain

	// allocate swap images
	vkSwapChain.images = f.createSwapChainImages(vkSwapChain.ref)
	vkSwapChain.imagesView = f.createSwapChainImageViews(vkSwapChain.images)

	log.Printf("VK: swapchain created, images=%d, info=(%s)\n", len(vkSwapChain.images), vkSwapChain.info.String())
	return vkSwapChain
}

func (f *vkSwapChainFactory) createSwapChainImages(swapChain vulkan.Swapchain) []vulkan.Image {
	imagesCount := uint32(0)
	vkAssert(
		vulkan.GetSwapchainImages(f.ld.ref, swapChain, &imagesCount, nil),
		fmt.Errorf("failed fetch swapChain images"),
	)
	if imagesCount == 0 {
		panic(fmt.Errorf("swapchain should have at least 1 image buffer"))
	}

	images := make([]vulkan.Image, imagesCount)
	vkAssert(
		vulkan.GetSwapchainImages(f.ld.ref, swapChain, &imagesCount, images),
		fmt.Errorf("failed fetch swapChain images"),
	)

	return images
}

func (f *vkSwapChainFactory) createSwapChainImageViews(images []vulkan.Image) []vulkan.ImageView {
	views := make([]vulkan.ImageView, 0, len(images))

	for _, image := range images {
		views = append(views, f.createSwapChainImageView(image))
	}

	return views
}

func (f *vkSwapChainFactory) createSwapChainImageView(image vulkan.Image) vulkan.ImageView {
	createInfo := &vulkan.ImageViewCreateInfo{
		SType:    vulkan.StructureTypeImageViewCreateInfo,
		Image:    image,
		ViewType: vulkan.ImageViewType2d,
		Format:   (*f.pd.surfaceProps.richColorSpaceFormat()).Format,
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
		vulkan.CreateImageView(f.ld.ref, createInfo, nil, &imageView),
		fmt.Errorf("failed create image view"),
	)

	return imageView
}
