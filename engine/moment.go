package engine

import "time"

type Moment interface {
	FPS() int
	TargetFPS() int
	FrameDuration() time.Duration
	LimitDuration() time.Duration
	DeltaTime() float64
	SinceStart() time.Duration
}
