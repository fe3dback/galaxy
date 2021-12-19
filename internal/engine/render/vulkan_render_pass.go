package render

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

func (vk *vk) vkCreateRenderPass(opts vkCreateOptions) vulkan.RenderPass {
	colorAttachment := vulkan.AttachmentDescription{
		Format:         vk.swapChain.info.imageFormat,
		Samples:        vulkan.SampleCount1Bit,
		LoadOp:         vulkan.AttachmentLoadOpClear,
		StoreOp:        vulkan.AttachmentStoreOpStore,
		StencilLoadOp:  vulkan.AttachmentLoadOpDontCare,
		StencilStoreOp: vulkan.AttachmentStoreOpDontCare,
		InitialLayout:  vulkan.ImageLayoutUndefined,
		FinalLayout:    vulkan.ImageLayoutPresentSrc,
	}

	colorAttachmentRef := vulkan.AttachmentReference{
		Attachment: 0,
		Layout:     vulkan.ImageLayoutColorAttachmentOptimal,
	}

	subpass := vulkan.SubpassDescription{
		PipelineBindPoint:       vulkan.PipelineBindPointGraphics,
		InputAttachmentCount:    0,
		PInputAttachments:       nil,
		ColorAttachmentCount:    1,
		PColorAttachments:       []vulkan.AttachmentReference{colorAttachmentRef},
		PResolveAttachments:     nil,
		PDepthStencilAttachment: nil,
		PreserveAttachmentCount: 0,
		PPreserveAttachments:    nil,
	}

	renderPassInfo := &vulkan.RenderPassCreateInfo{
		SType:           vulkan.StructureTypeRenderPassCreateInfo,
		AttachmentCount: 1,
		PAttachments:    []vulkan.AttachmentDescription{colorAttachment},
		SubpassCount:    1,
		PSubpasses:      []vulkan.SubpassDescription{subpass},
	}

	var renderPass vulkan.RenderPass
	vkAssert(
		vulkan.CreateRenderPass(vk.logicalDevice.ref, renderPassInfo, nil, &renderPass),
		fmt.Errorf("failed create render pass"),
	)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroyRenderPass(vk.logicalDevice.ref, renderPass, nil)
	})

	return renderPass
}
