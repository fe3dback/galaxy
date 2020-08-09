package main

import (
	"fmt"
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

	fmt.Printf("-- system:\n")
	fmt.Printf("      mem: %s\n", utils.FormatBytes(m.Alloc))
	fmt.Printf("total mem: %s\n", utils.FormatBytes(m.TotalAlloc))
	fmt.Printf("  sys mem: %s\n", utils.FormatBytes(m.Sys))
}

func debugPrintFps(f *system.Frames) {
	fmt.Printf("-- frames:\n")
	fmt.Printf("           FPS: %d / %d\n", f.FPS(), f.TargetFPS())
	fmt.Printf("frame duration: %s\n", f.FrameDuration())
	fmt.Printf("frame throttle: %s\n", f.LimitDuration())
	fmt.Printf("limit duration: %s\n", f.FrameThrottle())
	fmt.Printf("    delta time: %f\n", f.DeltaTime())
	fmt.Printf("       seconds: %s\n", f.SinceStart())
}

func debugPrintWorld(w *game.World) {
	fmt.Printf("-- world:\n")
	fmt.Printf("      entities: %d\n", len(w.Entities()))
}
