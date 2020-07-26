package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/game"
)

const fpsLimit = 60

func main() {
	err := run(options{
		frames: framesOpt{
			targetFps: fpsLimit,
		},
		debug: debugOpt{
			frames: true,
		},
	})
	if err != nil {
		panic(err)
	}
}

func run(opt options) error {
	level := game.NewBasicLevel()
	var err error

	frames := NewFrames(opt.frames.targetFps)

	for {
		frames.Begin()

		err = level.OnUpdate(frames.DeltaTime())
		if err != nil {
			return fmt.Errorf("can`t update level: %v", err)
		}

		err = level.OnDraw()
		if err != nil {
			return fmt.Errorf("can`t draw level: %v", err)
		}

		frames.End()

		if opt.debug.frames {
			debugPrintFps(frames)
		}
	}
}

func debugPrintFps(f *frames) {
	fmt.Printf("--\n")
	fmt.Printf("           FPS: %d / %d\n", f.fps, f.limitFps)
	fmt.Printf("frame duration: %s\n", f.frameDuration)
	fmt.Printf("frame throttle: %s\n", f.frameThrottle)
	fmt.Printf("limit duration: %s\n", f.limitDuration)
	fmt.Printf("    delta time: %f\n", f.DeltaTime())
	fmt.Printf("       seconds: %s\n", f.Seconds())
}
