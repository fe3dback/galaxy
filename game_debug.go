package main

import (
	"fmt"
	"runtime"

	"github.com/fe3dback/galaxy/game"
)

func debug(params *gameParams) {
	if params.options.debug.system {
		debugSystem()
	}

	if params.options.debug.frames {
		debugPrintFps(params.frames)
	}

	if params.options.debug.world {
		debugPrintWorld(params.world)
	}
}

func debugSystem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("-- system:\n")
	fmt.Printf("      mem: %s\n", formatBytes(m.Alloc))
	fmt.Printf("total mem: %s\n", formatBytes(m.TotalAlloc))
	fmt.Printf("  sys mem: %s\n", formatBytes(m.Sys))
}

func debugPrintFps(f *frames) {
	fmt.Printf("-- frames:\n")
	fmt.Printf("           FPS: %d / %d\n", f.fps, f.limitFps)
	fmt.Printf("frame duration: %s\n", f.frameDuration)
	fmt.Printf("frame throttle: %s\n", f.frameThrottle)
	fmt.Printf("limit duration: %s\n", f.limitDuration)
	fmt.Printf("    delta time: %f\n", f.DeltaTime())
	fmt.Printf("       seconds: %s\n", f.Seconds())
}

func debugPrintWorld(w *game.World) {
	fmt.Printf("-- world:\n")
	fmt.Printf("      entities: %d\n", len(w.Entities()))
}
