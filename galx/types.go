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
		StateToGameMode()
		StateToEditorMode()
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

		Lock()
		Unlock()
		IsLocked() bool
		Select()
		Unselect()
		IsSelected() bool

		AbsPosition() Vec2d
		RelativePosition() Vec2d
		SetPosition(pos Vec2d)
		AddPosition(pos Vec2d)
		Rotation() Angle
		Scale() float64
		SetScale(scale float64)
		SetRotation(rot Angle)
		AddRotation(rot Angle)

		Components() map[string]Component
		GetComponent(ref Component) Component
		AddComponent(c Component)

		IsRoot() bool
		IsLeaf() bool
		HasChild() bool
		HasParent() bool
		Child() []GameObject
		AddChild(child GameObject)
		RemoveChild(id string)
		SetParent(parent GameObject)
		HierarchyLevel() uint8

		BoundingBox(padding float64) Rect
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

	EditorComponentResolver     = func(EditorComponentIdentifiable) EditorComponentIdentifiable
	EditorComponentIdentifiable interface {
		Id() string
	}
	EditorComponentCycleCreated interface {
		OnCreated(EditorComponentResolver)
	}
)

// --------------------------------------------
// State
// --------------------------------------------

const (
	QueryFlagExcludeLocked QueryFlag = 1 << iota
	QueryFlagExcludeRoots
	QueryFlagExcludeLeaf
	QueryFlagOnlyOnScreen
	QueryFlagOnlySelected
)

type (
	State interface {
		Camera() Camera
		Moment() Moment
		Mouse() Mouse
		Keyboard() Keyboard
		Movement() Movement
		EngineState() EngineState
		SoundMixer() SoundMixer
		Scene() Scene
		ObjectQueries() ObjectQueries
	}

	QueryFlag = uint32

	ObjectQueries interface {
		All() []GameObject
		AllIn(QueryFlag) []GameObject
	}

	Camera interface {
		Screen2World(screen Vec2d) Vec2d
		World2Screen(world Vec2d) Vec2d
		Position() Vec2d
		Width() int
		Height() int
		Scale() float64
		MoveTo(p Vec2d)
		CenterOn(p Vec2d)
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
		FrameId() uint64
		FrameDuration() time.Duration
		LimitDuration() time.Duration
		DeltaTime() float64
		SinceStart() time.Duration
	}
)

// --------------------------------------------
// Control
// --------------------------------------------

const (
	MousePropagationPriorityGame = iota + 1
	MousePropagationPriorityGameUI
	MousePropagationPriorityEditor
	MousePropagationPriorityEditorSelect
	MousePropagationPriorityEditorGizmos
	MousePropagationPriorityEditorHigh
	MousePropagationPriorityHightest
)

type (
	MousePropagationPriority = int

	Mouse interface {
		MouseCoords() Vec2d
		ScrollPosition() float64
		ScrollLastOffset() float64

		IsButtonsAvailable(priority int) bool
		StopPropagation(priority int)
		LeftPressed() bool
		LeftReleased() bool
		LeftDown() bool
		RightPressed() bool
		RightReleased() bool
		RightDown() bool
	}

	Keyboard interface {
		IsPressed(key rune) bool
		IsReleased(key rune) bool
		IsDown(key rune) bool
	}

	Movement interface {
		Vector() Vec2d
		Shift() bool
		Space() bool
	}
)

// --------------------------------------------
// APIs
// --------------------------------------------

const (
	RenderTargetMain = 0
)

const (
	RenderModeWorld RenderMode = iota
	RenderModeUI
)

type (
	RenderTarget = uint8
	RenderMode   = uint8

	Renderer interface {
		// base

		SetDrawColor(Color)
		DrawSquare(color Color, rect Rect)
		DrawSquareFilled(color Color, rect Rect)
		DrawCircle(color Color, circle Circle)
		DrawSquareEx(color Color, angle Angle, rect Rect)
		DrawLine(color Color, line Line)
		DrawVector(color Color, dist float64, vec Vec2d, angle Angle)
		DrawCrossLines(color Color, size int, vec Vec2d)
		DrawPoint(color Color, vec Vec2d)

		// sprite

		TextureQuery(res consts.AssetsPath) TextureInfo
		DrawSprite(res consts.AssetsPath, vec Vec2d)
		DrawSpriteAngle(res consts.AssetsPath, vec Vec2d, angle Angle)
		DrawSpriteEx(res consts.AssetsPath, src, dest Rect, angle Angle)

		// text

		DrawText(fontId consts.AssetsPath, color Color, vec Vec2d, text string)

		// system

		Gizmos() Gizmos
	}

	SoundMixer interface {
		Play(res consts.AssetsPath)
	}
)
