package main

import (
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/di"
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/system"
	"github.com/fe3dback/galaxy/utils"
)

func debug(c *di.Container) {
	debug := c.ProvideGameOptions().Debug

	if debug.Memory {
		debugMemory()
	}

	if debug.Frames {
		debugPrintFps(c.ProvideFrames())
	}

	if debug.World {
		debugPrintWorld(c.ProvideEngineScenesManager().Current())
	}
}

func debugMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("memory: [alloc: %s, total: %s, sys: %s]",
		utils.FormatBytes(m.Alloc),
		utils.FormatBytes(m.TotalAlloc),
		utils.FormatBytes(m.Sys),
	)
}

func debugPrintFps(f *system.Frames) {
	log.Printf("frames: [fps: %d / %d, duration: %s / %s, throttle: %s]",
		f.FPS(), f.TargetFPS(),
		f.FrameDuration(),
		f.LimitDuration(),
		f.FrameThrottle(),
	)
	log.Printf("frames: [dt: %.4f, sec: %s]",
		f.DeltaTime(),
		f.SinceStart(),
	)
}

func debugPrintWorld(s engine.Scene) {
	log.Printf("scene: [entities: %d]",
		len(s.Entities()),
	)
}
