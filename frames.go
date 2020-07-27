package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type frames struct {
	limitFps      int64
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

func NewFrames(targetFps int64) *frames {
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
	f.throttleGame()

	// update state
	f.deltaTime = time.Since(f.frameStart)
}

func (f *frames) FPS() int {
	return f.fps
}

func (f *frames) TotalFPS() int {
	//      max 16ms
	// throttle 15ms
	//    frame 1ms
	// framesSkippedAvg = 15
	// avgFPS = 60
	// totalFPS = 60 * 15 ~= 900

	framesSkippedAvg := f.frameThrottle.Seconds() / f.frameDuration.Seconds()
	if framesSkippedAvg <= 1 {
		framesSkippedAvg = 0
	}

	skippedPerSecond := framesSkippedAvg * float64(f.fps)
	totalFPS := f.FPS() + int(skippedPerSecond)

	// round total fps
	return totalFPS / 100 * 100
}

func (f *frames) DeltaTime() float64 {
	return f.deltaTime.Seconds()
}

func (f *frames) Seconds() time.Duration {
	return time.Since(f.gameStart)
}

func (f *frames) Ready() bool {
	return !f.isInterrupted
}

func (f *frames) Interrupt() {
	f.isInterrupted = true
}

func (f *frames) throttleGame() {
	if f.frameThrottle <= 0 {
		return
	}

	time.Sleep(f.frameThrottle)
}
