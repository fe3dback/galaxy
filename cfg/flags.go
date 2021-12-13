package cfg

import (
	"time"

	"github.com/fe3dback/galaxy/galx"
)

type (
	InitFlags struct {
		// process
		isProfiling   bool
		profilingPort int

		// system
		includeEditor bool
		seed          int64

		// game
		defaultIsGameMode bool
		components        map[string]galx.Component

		// render
		targetFPS    int
		isFullScreen bool
		screenWidth  int
		screenHeight int

		debugOpt DebugOpt
	}

	DebugOpt struct {
		System bool
		Memory bool
		Frames bool
		World  bool
	}
)

func NewInitFlags(modifiers ...Modifier) *InitFlags {
	flags := &InitFlags{
		// process
		isProfiling:   false,
		profilingPort: 0,

		// system
		includeEditor:     true,
		defaultIsGameMode: false,
		seed:              time.Now().Unix(),

		// game
		components: map[string]galx.Component{},

		// render
		targetFPS:    60,
		isFullScreen: false,
		screenWidth:  960,
		screenHeight: 540,
	}

	for _, modifier := range modifiers {
		modifier(flags)
	}

	return flags
}

func (f *InitFlags) IsProfiling() bool {
	return f.isProfiling
}

func (f *InitFlags) ProfilingPort() int {
	return f.profilingPort
}

func (f *InitFlags) Seed() int64 {
	return f.seed
}

func (f *InitFlags) TargetFPS() int {
	return f.targetFPS
}

func (f *InitFlags) IsFullscreen() bool {
	return f.isFullScreen
}

func (f *InitFlags) ScreenWidth() int {
	return f.screenWidth
}

func (f *InitFlags) ScreenHeight() int {
	return f.screenHeight
}

func (f *InitFlags) DebugOpts() DebugOpt {
	return f.debugOpt
}

func (f *InitFlags) IsIncludeEditor() bool {
	return f.includeEditor
}

func (f *InitFlags) IsDefaultGameMode() bool {
	return f.defaultIsGameMode
}

func (f *InitFlags) Components() map[string]galx.Component {
	return f.components
}
