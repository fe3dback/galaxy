package galx

import (
	"time"

	"github.com/fe3dback/galaxy/consts"
)

type (
	EngineState interface {
		InEditorMode() bool
	}

	Camera interface {
		Position() Vec
		Width() int
		Height() int
		Zoom() float64
		MoveTo(p Vec)
		CenterOn(p Vec)
		Resize(width, height int)
		ZoomView(scale float64)
	}

	Drawer interface {
		OnDraw(Renderer) error
	}

	RenderMode = uint8

	Gizmos interface {
		System() bool
		Primary() bool
		Secondary() bool
		Debug() bool
		Spam() bool
	}

	Renderer interface {
		// base
		SetDrawColor(Color)
		DrawSquare(color Color, rect Rect)
		DrawCircle(color Color, circle Circle)
		DrawSquareEx(color Color, angle Angle, rect Rect)
		DrawLine(color Color, line Line)
		DrawVector(color Color, dist float64, vec Vec, angle Angle)
		DrawCrossLines(color Color, size int, vec Vec)
		DrawPoint(color Color, vec Vec)

		// camera
		Camera() Camera

		// sprite
		TextureQuery(res consts.AssetsPath) TextureInfo
		DrawSprite(res consts.AssetsPath, vec Vec)
		DrawSpriteAngle(res consts.AssetsPath, vec Vec, angle Angle)
		DrawSpriteEx(res consts.AssetsPath, src, dest Rect, angle Angle)

		// text
		DrawText(fontId consts.AssetsPath, color Color, text string, vec Vec)

		// system
		InEditorMode() bool
		Gizmos() Gizmos
		SetRenderTarget(id uint8)
		SetRenderMode(RenderMode)
		FillRect(Rect)
		Clear(Color)
		EndEngineFrame()
		UpdateGPU()

		// gui
		StartGUIFrame()
		EndGUIFrame()
	}

	Updater interface {
		OnUpdate(State) error
	}

	Moment interface {
		FPS() int
		TargetFPS() int
		FrameId() int
		FrameDuration() time.Duration
		LimitDuration() time.Duration
		DeltaTime() float64
		SinceStart() time.Duration
	}

	// Controls

	Mouse interface {
		MouseCoords() Vec
		ScrollPosition() float64
		ScrollLastOffset() float64
	}

	Movement interface {
		Vector() Vec
		Shift() bool
		Space() bool
	}

	// Game State

	SoundMixer interface {
		Play(res consts.AssetsPath)
	}

	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
		Movement() Movement
		EngineState() EngineState
		SoundMixer() SoundMixer
		Scene() Scene
	}

	GameObject interface {
		Drawer
		Updater
		Destroy()
		IsDestroyed() bool
		Id() string
		Position() Vec
		SetPosition(pos Vec)
		AddPosition(pos Vec)
		Rotation() Angle
		SetRotation(rot Angle)
		AddRotation(rot Angle)
	}

	Scene interface {
		OnUpdate(s State) error
		OnDraw(r Renderer) error
		Entities() []GameObject
	}

	SceneManager interface {
		Switch(nextID string)
		Current() Scene
	}
)
