package cfg

type (
	Modifier = func(*InitFlags)
)

func WithTargetFPS(fps int) Modifier {
	return func(flags *InitFlags) {
		flags.targetFPS = fps
	}
}

func WithProfiling(enabled bool, port int) Modifier {
	return func(flags *InitFlags) {
		flags.isProfiling = enabled
		flags.profilingPort = port
	}
}

func WithScreen(fullscreen bool, width int, height int) Modifier {
	return func(flags *InitFlags) {
		flags.isFullScreen = fullscreen
		flags.screenWidth = width
		flags.screenHeight = height
	}
}

func WithDebugOpts(system, memory, frames, world bool) Modifier {
	return func(flags *InitFlags) {
		flags.debugOpt = DebugOpt{
			System: system,
			Memory: memory,
			Frames: frames,
			World:  world,
		}
	}
}
