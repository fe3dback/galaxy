package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkFrameBuffers struct {
		buffers []vulkan.Framebuffer

		_free   bool
		_freeLd *vkLogicalDevice
	}
)

func (fb *vkFrameBuffers) free() {
	if fb._free {
		return
	}

	fb._free = true
	for _, buffer := range fb.buffers {
		vulkan.DestroyFramebuffer(fb._freeLd.ref, buffer, nil)
	}
}

func createFrameBuffers(swapChain *vkSwapChain, ld *vkLogicalDevice, renderPass vulkan.RenderPass, closer *utils.Closer) *vkFrameBuffers {
	vkFrameBuffers := &vkFrameBuffers{
		_freeLd: ld,
	}
	closer.EnqueueFree(vkFrameBuffers.free)

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

		buffers = append(buffers, buffer)
	}

	vkFrameBuffers.buffers = buffers
	return vkFrameBuffers
}
