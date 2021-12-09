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
		TextureQuery(res generated.ResourcePath) TextureInfo
		DrawSprite(res generated.ResourcePath, vec Vec)
		DrawSpriteAngle(res generated.ResourcePath, vec Vec, angle Angle)
		DrawSpriteEx(res generated.ResourcePath, src, dest Rect, angle Angle)

		// text
		DrawText(fontId generated.ResourcePath, color Color, text string, vec Vec)

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

	// Engine Assets

	LoaderYaml interface {
		LoadYaml(res generated.ResourcePath, data interface{})
	}

	LoaderSound interface {
		LoadSound(res generated.ResourcePath)
	}

	Loader interface {
		LoaderYaml
		LoaderSound
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
		Play(res generated.ResourcePath)
	}

	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
		Movement() Movement
		InEditorMode() bool
		SoundMixer() SoundMixer
		Scene() Scene
	}

	GameObject interface {
		Drawer
		Updater
		Destroy()
		IsDestroyed() bool
		Id() int64
		Position() Vec
		SetPosition(pos Vec)
		Rotation() Angle
		SetRotation(rot Angle)
	}

	Scene interface {
		OnUpdate(s State) error
		OnDraw(r Renderer) error
		Entities() []GameObject
	}

	SceneBlueprint interface {
		CreateEntities() []GameObject
	}

	SceneManager interface {
		AddBlueprint(ID string, blueprint SceneBlueprint)
		Switch(nextID string)
		Current() Scene
	}
)
