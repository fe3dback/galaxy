package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkFrameBuffers struct {
		buffers []vulkan.Framebuffer
	}
)

func createFrameBuffers(swapChain *vkSwapChain, ld *vkLogicalDevice, renderPass vulkan.RenderPass, closer *utils.Closer) *vkFrameBuffers {
	buffers := make([]vulkan.Framebuffer, 0, len(swapChain.imagesView))

	for _, view := range swapChain.imagesView {
		createInfo := &vulkan.FramebufferCreateInfo{
			SType:           vulkan.StructureTypeFramebufferCreateInfo,
			RenderPass:      renderPass,
			AttachmentCount: 1,
			PAttachments: []vulkan.ImageView{
				view,
			},
			Width:  swapChain.info.bufferSize.Width,
			Height: swapChain.info.bufferSize.Height,
			Layers: 1,
		}

		var buffer vulkan.Framebuffer
		vkAssert(
			vulkan.CreateFramebuffer(ld.ref, createInfo, nil, &buffer),
			fmt.Errorf("failed create frame buffer"),
		)

		closer.EnqueueFree(func() {
			vulkan.DestroyFramebuffer(ld.ref, buffer, nil)
		})

		buffers = append(buffers, buffer)
	}

	return &vkFrameBuffers{
		buffers: buffers,
	}
}
