package frames

import (
	"log"
	"math"
	"os"
	"os/signal"
	"runtime"
	"time"
)

const gameSpeedMultiplier = 1

type Frames struct {
	limitFps      int
	limitDuration time.Duration

	total         int
	count         int
	fps           int
	gameStart     time.Time
	frameStart    time.Time
	frameDuration time.Duration
	frameThrottle time.Duration
	deltaTime     time.Duration
	isInterrupted bool
}

func NewFrames(targetFps int) *Frames {
	f := &Frames{
		limitFps:      targetFps,
		limitDuration: time.Second / time.Duration(targetFps),
		total:         0,
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

func (f *Frames) listenTimer() {
	go func() {
		for range time.Tick(time.Second) {
			// count fps
			f.fps = f.count
			f.count = 0
		}
	}()
}

func (f *Frames) listenOsSignals() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		sig := <-c
		f.isInterrupted = true

		log.Printf("got os signal: %v", sig)
	}()
}

func (f *Frames) Begin() {
	f.frameStart = time.Now()
}

func (f *Frames) End() {
	if f.total == math.MaxInt16 {
		f.total = 0
	} else {
		f.total++
	}

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

func (f *Frames) FPS() int {
	return f.fps
}

func (f *Frames) TargetFPS() int {
	return f.limitFps
}

func (f *Frames) FrameId() int {
	return f.total
}

func (f *Frames) FrameDuration() time.Duration {
	return f.frameDuration
}

func (f *Frames) LimitDuration() time.Duration {
	return f.limitDuration
}

func (f *Frames) FrameThrottle() time.Duration {
	return f.frameThrottle
}

func (f *Frames) DeltaTime() float64 {
	return f.deltaTime.Seconds() * gameSpeedMultiplier
}

func (f *Frames) SinceStart() time.Duration {
	return time.Since(f.gameStart)
}

func (f *Frames) Ready() bool {
	return !f.isInterrupted
}

func (f *Frames) Interrupt() {
	f.isInterrupted = true
}

func (f *Frames) afterFrame() {
	if f.frameThrottle <= 0 {
		return
	}

	// force run GC, because we have free time in this frame
	gcStart := time.Now()
	runtime.GC()
	runtime.Gosched()

	time.Sleep(f.frameThrottle - time.Since(gcStart))
}
