package vulkan_depr

import (
	"github.com/vulkan-go/vulkan"
)

func (vk *Vk) Draw() {
	if vk.isResizing || !vk.isDrawAvailable {
		return
	}

	vk.swapChainFrameManager.drawFrame()
}

func (vk *Vk) Wait() {
	vulkan.DeviceWaitIdle(vk.logicalDevice.ref)
}
