package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkSurface struct {
		ref vulkan.Surface
	}
)

func (inst *vkInstance) vkCreateSurface(opts vkCreateOptions) *vkSurface {
	surfacePtr, err := opts.window.CreateWindowSurface(inst.ref, nil)
	if err != nil {
		panic(fmt.Errorf("failed create vulkan windows surface: %w", err))
	}

	surface := vulkan.SurfaceFromPointer(surfacePtr)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroySurface(inst.ref, surface, nil)
	})

	return &vkSurface{
		ref: surface,
	}
}
