package windows

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine"
	"github.com/fe3dback/galaxy/internal/engine/event"
	"github.com/fe3dback/galaxy/internal/utils"
)

type Manager struct {
	window *glfw.Window
}

func NewManager(closer *utils.Closer, dispatcher *event.Dispatcher, tech engine.RenderTech, defaultWidth, defaultHeight int, fullscreen bool) *Manager {
	var window *glfw.Window

	switch tech {
	case engine.RenderTechVulkan:
		window = newVulkanWindow(closer, dispatcher, defaultWidth, defaultHeight, fullscreen)
	default:
		panic(fmt.Errorf("failed create window: not supported render tech: %s", tech))
	}

	return &Manager{
		window: window,
	}
}

func (m *Manager) Window() *glfw.Window {
	return m.window
}

func newVulkanWindow(closer *utils.Closer, dispatcher *event.Dispatcher, defaultWidth, defaultHeight int, fullscreen bool) *glfw.Window {
	// set vulkan address
	procAddr := glfw.GetVulkanGetInstanceProcAddress()
	if procAddr == nil {
		panic(fmt.Errorf("failed get vulkan proc address"))
	}
	vulkan.SetGetInstanceProcAddr(procAddr)

	// init
	err := glfw.Init()
	if err != nil {
		panic(fmt.Errorf("failed init glfw library: %w", err))
	}
	closer.EnqueueFree(glfw.Terminate)

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	// create window
	var monitor *glfw.Monitor
	if fullscreen {
		monitor = glfw.GetPrimaryMonitor()
	}

	window, err := glfw.CreateWindow(defaultWidth, defaultHeight, "Galaxy", monitor, nil)
	if err != nil {
		panic(fmt.Errorf("failed create glfw window: %w", err))
	}
	closer.EnqueueFree(window.Destroy)

	window.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		dispatcher.PublishEventWindowResized(event.WindowResizedEvent{
			NewWidth:  width,
			NewHeight: height,
		})
	})

	// return
	return window
}
