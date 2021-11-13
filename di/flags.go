package di

import (
	"time"
)

type InitFlags struct {
	// process
	IsProfiling   bool
	ProfilingPort int

	// system
	Seed int64

	// render
	TargetFPS  int
	FullScreen bool
	Width      int
	Height     int
}

func NewInitFlags() *InitFlags {
	seed := time.Now().Unix()
	return &InitFlags{
		// process
		IsProfiling:   false,
		ProfilingPort: 0,

		// system
		Seed: seed,

		// render
		TargetFPS:  60,
		FullScreen: false,
		Width:      960,
		Height:     540,
	}
}
