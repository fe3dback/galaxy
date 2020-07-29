package main

import (
	"fmt"
	"runtime"

	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/utils"
)

func debug(provider *provider) {
	debug := provider.registry.game.options.debug

	if debug.system {
		debugSystem()
	}

	if debug.frames {
		debugPrintFps(provider.registry.game.frames)
	}

	if debug.world {
		debugPrintWorld(provider.registry.game.world)
	}
}

func debugSystem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("-- system:\n")
	fmt.Printf("      mem: %s\n", utils.FormatBytes(m.Alloc))
	fmt.Printf("total mem: %s\n", utils.FormatBytes(m.TotalAlloc))
	fmt.Printf("  sys mem: %s\n", utils.FormatBytes(m.Sys))
}

func debugPrintFps(f *frames) {
	fmt.Printf("-- frames:\n")
	fmt.Printf("           FPS: %d / %d\n", f.fps, f.limitFps)
	fmt.Printf("frame duration: %s\n", f.frameDuration)
	fmt.Printf("frame throttle: %s\n", f.frameThrottle)
	fmt.Printf("limit duration: %s\n", f.limitDuration)
	fmt.Printf("    delta time: %f\n", f.DeltaTime())
	fmt.Printf("       seconds: %s\n", f.SinceStart())
}

func debugPrintWorld(w *game.World) {
	fmt.Printf("-- world:\n")
	fmt.Printf("      entities: %d\n", len(w.Entities()))
}
