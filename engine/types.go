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
		Position() Vec
		Width() int
		Height() int
		MoveTo(p Vec)
		CenterOn(p Vec)
		Resize(width, height int)
	}

	Drawer interface {
		OnDraw(Renderer) error
	}

	RenderMode = uint8

	Renderer interface {
		// base
		SetDrawColor(Color)
		DrawSquare(color Color, rect Rect)
		DrawSquareEx(color Color, angle Angle, rect Rect)
		DrawLine(color Color, line Line)
		DrawVector(color Color, dist float64, vec Vec, angle Angle)
		DrawCrossLines(color Color, size int, vec Vec)
		DrawPoint(color Color, vec Vec)

		// camera
		Camera() Camera

		// sprite
		TextureQuery(res generated.ResourcePath) TextureInfo
		DrawSprite(res generated.ResourcePath, vec Vec)
		DrawSpriteAngle(res generated.ResourcePath, vec Vec, angle Angle)
		DrawSpriteEx(res generated.ResourcePath, src, dest Rect, angle Angle)

		// text
		DrawText(fontId generated.ResourcePath, color Color, text string, vec Vec)

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
		MouseCoords() Vec
	}

	Movement interface {
		Vector() Vec
		Shift() bool
	}

	// Game State

	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
		Movement() Movement
	}
)
