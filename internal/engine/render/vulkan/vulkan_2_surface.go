package vulkan

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkSurface struct {
		ref vulkan.Surface
	}
)

func (inst *vkInstance) createSurfaceFromWindow(window *glfw.Window, closer *utils.Closer) *vkSurface {
	surfacePtr, err := window.CreateWindowSurface(inst.ref, nil)
	if err != nil {
		panic(fmt.Errorf("failed create vulkan windows surface: %w", err))
	}

	surface := vulkan.SurfaceFromPointer(surfacePtr)
	closer.EnqueueFree(func() {
		vulkan.DestroySurface(inst.ref, surface, nil)
	})

	return &vkSurface{
		ref: surface,
	}
}
