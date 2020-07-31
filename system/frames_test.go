package system

import (
	"testing"
	"time"
)

func TestNewFrames(t *testing.T) {
	type args struct {
		targetFps int
	}
	tests := []struct {
		name string
		args args
		want *Frames
	}{
		{
			name: "fps 60",
			args: args{targetFps: 60},
			want: &Frames{
				limitFps:      60,
				limitDuration: time.Second / 60,
			},
		},
		{
			name: "fps 30",
			args: args{targetFps: 30},
			want: &Frames{
				limitFps:      30,
				limitDuration: time.Second / 30,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFrames(tt.args.targetFps)

			if tt.want.limitDuration != got.limitDuration {
				t.Errorf("limit duration, want %v, got %v", tt.want.limitDuration, got.limitDuration)
			}

			if tt.want.limitFps != got.limitFps {
				t.Errorf("limit fps, want %v, got %v", tt.want.limitFps, got.limitFps)
			}
		})
	}
}

func TestFrames(t *testing.T) {

	frames := NewFrames(60)
	frame := 0
	maxFrames := 10

	for {
		frames.Begin()
		time.Sleep(time.Millisecond * 10)
		frames.End()

		frame++
		if frame >= maxFrames {
			break
		}
	}

	if frames.count != maxFrames {
		t.Errorf("frames should by %v", maxFrames)
	}

	// expect 16ms frame in 60fps
	if !(frames.frameDuration > time.Millisecond*10 && frames.frameThrottle < time.Millisecond*7) {
		t.Errorf("frames calculated incorrect: dur: %v, throttle: %v", frames.frameDuration, frames.frameThrottle)
	}
}
