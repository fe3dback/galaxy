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
		Position() Vector2D
		Width() int
		Height() int
		MoveTo(p Vector2D)
		CenterOn(p Vector2D)
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
		DrawVector(color Color, dist float64, vec Vector2D, angle Angle)
		DrawCrossLines(color Color, size int, p Point)
		DrawPoint(color Color, p Point)

		// camera
		Camera() Camera

		// sprite
		TextureQuery(res generated.ResourcePath) TextureInfo
		DrawSprite(res generated.ResourcePath, p Point)
		DrawSpriteAngle(res generated.ResourcePath, p Point, angle Angle)
		DrawSpriteEx(res generated.ResourcePath, src, dest Rect, angle Angle)

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

	Movement interface {
		Vector() Vector2D
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
