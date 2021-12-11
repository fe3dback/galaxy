package galaxy

import (
	"log"
	"runtime"

	"github.com/fe3dback/galaxy/internal/di"
	"github.com/fe3dback/galaxy/internal/frames"
	"github.com/fe3dback/galaxy/internal/utils"
)

func debug(c *di.Container) {
	debug := c.Flags().DebugOpts()

	if debug.Memory {
		debugMemory()
	}

	if debug.Frames {
		debugPrintFps(c.ProvideFrames())
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

func debugPrintFps(f *frames.Frames) {
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
