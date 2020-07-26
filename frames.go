package main

import (
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
	}

	go func() {
		for range time.Tick(time.Second) {
			// count fps
			f.fps = f.count
			f.count = 0
		}
	}()

	return f
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

func (f *frames) DeltaTime() float64 {
	return f.deltaTime.Seconds()
}

func (f *frames) Seconds() time.Duration {
	return time.Since(f.gameStart)
}

func (f *frames) throttleGame() {
	if f.frameThrottle <= 0 {
		return
	}

	time.Sleep(f.frameThrottle)
}
