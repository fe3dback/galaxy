package vulkan

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"
)

func newSurfaceFromWindow(inst *vkInstance, window *glfw.Window) *vkSurface {
	surfacePtr, err := window.CreateWindowSurface(inst.ref, nil)
	if err != nil {
		panic(fmt.Errorf("failed create vulkan windows surface: %w", err))
	}

	return &vkSurface{
		inst: inst,
		ref:  vulkan.SurfaceFromPointer(surfacePtr),
	}
}

func (surf *vkSurface) free() {
	vulkan.DestroySurface(surf.inst.ref, surf.ref, nil)
}
