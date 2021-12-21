package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkFrameBuffers struct {
		buffers []vulkan.Framebuffer
	}
)

func (vk *Vk) vkCreateFrameBuffers(opts vkCreateOptions) *vkFrameBuffers {
	buffers := make([]vulkan.Framebuffer, 0, len(vk.swapChain.imagesView))

	for _, view := range vk.swapChain.imagesView {
		createInfo := &vulkan.FramebufferCreateInfo{
			SType:           vulkan.StructureTypeFramebufferCreateInfo,
			RenderPass:      vk.renderPass,
			AttachmentCount: 1,
			PAttachments: []vulkan.ImageView{
				view,
			},
			Width:  vk.swapChain.info.bufferSize.Width,
			Height: vk.swapChain.info.bufferSize.Height,
			Layers: 1,
		}

		var buffer vulkan.Framebuffer
		vkAssert(
			vulkan.CreateFramebuffer(vk.logicalDevice.ref, createInfo, nil, &buffer),
			fmt.Errorf("failed create frame buffer"),
		)

		opts.closer.EnqueueFree(func() {
			vulkan.DestroyFramebuffer(vk.logicalDevice.ref, buffer, nil)
		})

		buffers = append(buffers, buffer)
	}

	return &vkFrameBuffers{
		buffers: buffers,
	}
}
