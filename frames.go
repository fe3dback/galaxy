package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"
)

type frames struct {
	limitFps      int
	limitDuration time.Duration

	count         int
	fps           int
	gameStart     time.Time
	frameStart    time.Time
	frameDuration time.Duration
	frameThrottle time.Duration
	deltaTime     time.Duration
	isInterrupted bool
}

func NewFrames(targetFps int) *frames {
	f := &frames{
		limitFps:      targetFps,
		limitDuration: time.Second / time.Duration(targetFps),
		count:         0,
		fps:           0,
		gameStart:     time.Now(),
		frameStart:    time.Now(),
		frameDuration: 0,
		frameThrottle: 0,
		isInterrupted: false,
	}

	f.listenTimer()
	f.listenOsSignals()

	return f
}

func (f *frames) listenTimer() {
	go func() {
		for range time.Tick(time.Second) {
			// count fps
			f.fps = f.count
			f.count = 0
		}
	}()
}

func (f *frames) listenOsSignals() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		sig := <-c
		f.isInterrupted = true

		fmt.Printf("got os signal %v\n", sig)
	}()
}

func (f *frames) Begin() {
	f.frameStart = time.Now()
}

func (f *frames) End() {
	f.count++
	f.frameDuration = time.Since(f.frameStart)

	if f.frameDuration < f.limitDuration {
		f.frameThrottle = f.limitDuration - f.frameDuration
	} else {
		f.frameThrottle = 0
	}

	// do additional logic
	f.afterFrame()

	// update state
	f.deltaTime = time.Since(f.frameStart)
}

func (f *frames) FPS() int {
	return f.fps
}

func (f *frames) TargetFPS() int {
	return f.limitFps
}

func (f *frames) FrameDuration() time.Duration {
	return f.frameDuration
}

func (f *frames) LimitDuration() time.Duration {
	return f.limitDuration
}

func (f *frames) DeltaTime() float64 {
	return f.deltaTime.Seconds()
}

func (f *frames) SinceStart() time.Duration {
	return time.Since(f.gameStart)
}

func (f *frames) Ready() bool {
	return !f.isInterrupted
}

func (f *frames) Interrupt() {
	f.isInterrupted = true
}

func (f *frames) afterFrame() {
	if f.frameThrottle <= 0 {
		return
	}

	// force run GC, because we have free time in this frame
	gcStart := time.Now()
	runtime.GC()
	runtime.Gosched()

	time.Sleep(f.frameThrottle - time.Since(gcStart))
}
