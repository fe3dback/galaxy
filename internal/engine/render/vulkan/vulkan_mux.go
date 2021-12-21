package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

func vkCreateSemaphores(ld *vkLogicalDevice, opts vkCreateOptions) (imageAvailable, presentAvailable vulkan.Semaphore) {
	createInfo := &vulkan.SemaphoreCreateInfo{
		SType: vulkan.StructureTypeSemaphoreCreateInfo,
	}

	vkAssert(
		vulkan.CreateSemaphore(ld.ref, createInfo, nil, &imageAvailable),
		fmt.Errorf("failed create image available mux"),
	)
	vkAssert(
		vulkan.CreateSemaphore(ld.ref, createInfo, nil, &presentAvailable),
		fmt.Errorf("failed create present available mux"),
	)

	opts.closer.EnqueueFree(func() {
		vulkan.DestroySemaphore(ld.ref, imageAvailable, nil)
		vulkan.DestroySemaphore(ld.ref, presentAvailable, nil)
	})

	return imageAvailable, presentAvailable
}
