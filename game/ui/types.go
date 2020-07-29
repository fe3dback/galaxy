package ui

import (
	"time"

	"github.com/fe3dback/galaxy/engine"
)

type (
	Layer interface {
		engine.Drawer
		engine.Updater
	}

	FramesProvider interface {
		FPS() int
		FrameDuration() time.Duration
		LimitDuration() time.Duration
	}
)
