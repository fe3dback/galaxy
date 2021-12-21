package cfg

import (
	"github.com/fe3dback/galaxy/galx"
)

type (
	Modifier = func(*InitFlags)
)

func WithProfiling(enabled bool, port int) Modifier {
	return func(flags *InitFlags) {
		flags.isProfiling = enabled
		flags.profilingPort = port
	}
}

func WithDebugOpts(system, memory, frames, world, vulkan bool) Modifier {
	return func(flags *InitFlags) {
		flags.debugOpt = DebugOpt{
			System: system,
			Memory: memory,
			Frames: frames,
			World:  world,
			Vulkan: vulkan,
		}
	}
}

func WithIncludeEditor(include bool, defaultIsGameMode bool) Modifier {
	return func(flags *InitFlags) {
		flags.includeEditor = include
		flags.defaultIsGameMode = defaultIsGameMode

		if !flags.includeEditor {
			// only game mode is possible, because app not bundle edit mode
			flags.defaultIsGameMode = true
		}
	}
}

func WithComponent(component galx.Component) Modifier {
	return func(flags *InitFlags) {
		flags.components[component.Id()] = component
	}
}
