package engine

import (
	"time"

	"github.com/fe3dback/galaxy/generated"
)

const (
	RenderModeWorld RenderMode = iota
	RenderModeUI
)

type (
	Camera interface {
		Rect() Rect
		MoveTo(p Point)
		CenterOn(p Point)
	}

	Drawer interface {
		OnDraw(Renderer) error
	}

	RenderMode = uint8

	Renderer interface {
		// base
		SetDrawColor(Color)
		DrawSquare(color Color, rect Rect)
		DrawLine(color Color, line Line)
		DrawCrossLines(color Color, size int, p Point)
		DrawPoint(color Color, p Point)

		// camera
		Camera() Camera

		// sprite
		TextureQuery(res generated.ResourcePath) TextureInfo
		DrawSprite(res generated.ResourcePath, p Point)
		DrawSpriteEx(res generated.ResourcePath, src, dest Rect, angle float64)

		// text
		DrawText(fontId generated.ResourcePath, color Color, text string, p Point)

		// system
		SetRenderMode(RenderMode)
		FillRect(Rect)
		Clear(Color)
		Present()
	}

	Updater interface {
		OnUpdate(State) error
	}

	Moment interface {
		FPS() int
		TargetFPS() int
		FrameDuration() time.Duration
		LimitDuration() time.Duration
		DeltaTime() float64
		SinceStart() time.Duration
	}

	// Controls

	Mouse interface {
		MouseCoords() Point
	}

	// Game State

	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
	}
)
