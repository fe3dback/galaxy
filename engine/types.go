package engine

import (
	"time"

	"github.com/fe3dback/galaxy/generated"
)

type (
	Drawer interface {
		OnDraw(Renderer) error
	}

	Renderer interface {
		// base
		SetDrawColor(Color)
		DrawSquare(color Color, x, y, w, h int)
		DrawLine(color Color, a, b Point)

		// sprite
		TextureQuery(res generated.ResourcePath) TextureInfo
		DrawSprite(res generated.ResourcePath, x, y int)
		DrawSpriteEx(res generated.ResourcePath, src, dest Rect, angle float64)

		// text
		DrawText(fontId generated.ResourcePath, color Color, text string, x, y int)

		// system
		FillRect(Rect)
		Clear(Color)
		Present()
	}

	Updater interface {
		OnUpdate(Moment) error
	}

	Moment interface {
		FPS() int
		TargetFPS() int
		FrameDuration() time.Duration
		LimitDuration() time.Duration
		DeltaTime() float64
		SinceStart() time.Duration
	}
)
