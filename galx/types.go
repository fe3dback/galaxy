package galx

import (
	"fmt"
	"time"

	"github.com/fe3dback/galaxy/consts"
)

// --------------------------------------------
// Engine
// --------------------------------------------

type (
	RenderMode = uint8

	EngineState interface {
		InEditorMode() bool
	}

	Updater interface {
		OnUpdate(State) error
	}

	Drawer interface {
		OnDraw(Renderer) error
	}
)

// --------------------------------------------
// ECS
// --------------------------------------------

type (
	SceneManager interface {
		Switch(nextID string)
		Current() Scene
	}

	Scene interface {
		OnUpdate(s State) error
		OnDraw(r Renderer) error
		Entities() []GameObject
	}

	GameObject interface {
		Drawer
		Updater
		fmt.Stringer
		Destroy()
		IsDestroyed() bool
		Id() string
		Name() string
		SetName(name string)

		AbsPosition() Vec
		SetPosition(pos Vec)
		AddPosition(pos Vec)
		Rotation() Angle
		SetRotation(rot Angle)
		AddRotation(rot Angle)

		Components() map[string]Component
		GetComponent(ref Component) Component
		AddComponent(c Component)

		IsRoot() bool
		IsLeaf() bool
		Child() []GameObject
		AddChild(child GameObject)
		RemoveChild(id string)
		SetParent(parent GameObject)
	}

	Component interface {
		Id() string
		Title() string
		Description() string
	}

	ComponentCycleCreated interface {
		OnCreated(entity GameObject)
	}

	ComponentCycleUpdated interface {
		Updater
	}

	ComponentCycleDrawer interface {
		Drawer
	}
)

// --------------------------------------------
// State
// --------------------------------------------

type (
	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
		Movement() Movement
		EngineState() EngineState
		SoundMixer() SoundMixer
		Scene() Scene
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

	Gizmos interface {
		System() bool
		Primary() bool
		Secondary() bool
		Debug() bool
		Spam() bool
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
)

// --------------------------------------------
// Control
// --------------------------------------------

type (
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
)

// --------------------------------------------
// APIs
// --------------------------------------------

type (
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

	SoundMixer interface {
		Play(res consts.AssetsPath)
	}
)
