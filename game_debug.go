package main

import (
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/registry"
	"github.com/fe3dback/galaxy/system"

	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/utils"
)

func debug(provider *registry.Provider) {
	debug := provider.Registry.Game.Options.Debug

	if debug.Memory {
		debugMemory()
	}

	if debug.Frames {
		debugPrintFps(provider.Registry.Game.Frames)
	}

	if debug.World {
		debugPrintWorld(provider.Registry.Game.WorldManager.CurrentWorld())
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

func debugPrintWorld(w *game.World) {
	log.Printf("world: [entities: %d]",
		len(w.Entities()),
	)
}
