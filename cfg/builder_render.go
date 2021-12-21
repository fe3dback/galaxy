package cfg

func WithTargetFPS(fps int) Modifier {
	return func(flags *InitFlags) {
		flags.targetFPS = fps
	}
}

func WithScreen(fullscreen bool, width int, height int) Modifier {
	return func(flags *InitFlags) {
		flags.isFullScreen = fullscreen
		flags.screenWidth = width
		flags.screenHeight = height
	}
}

// WithGraphicsVulkanDebug will print vulkan validation errors
// on stdout. It require installed Vulkan SDK to work
func WithGraphicsVulkanDebug(enabled bool) Modifier {
	return func(flags *InitFlags) {
		flags.vulkanOpt.Debug = enabled
	}
}

// WithGraphicsVSync will use FIFO rendering
// true - vsync, good for mobile (small power consumption)
// false - low latency, high power consumption
func WithGraphicsVSync(enabled bool) Modifier {
	return func(flags *InitFlags) {
		flags.vulkanOpt.VSync = enabled
	}
}
