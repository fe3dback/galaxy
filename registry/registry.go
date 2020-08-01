package registry

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine/lib/control"

	"github.com/fe3dback/galaxy/engine"

	"github.com/fe3dback/galaxy/engine/lib/event"

	"github.com/fe3dback/galaxy/engine/lib"

	"github.com/fe3dback/galaxy/engine/lib/render"
	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/game/ui"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/system"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	registerFactory struct{}
	Registry        struct {
		Closer *utils.Closer
		Sdl    *SdlRegistry
		Engine *EngineRegistry
		Game   *GameRegistry
	}

	SdlRegistry struct {
		Window *sdl.Window
	}

	EngineRegistry struct {
		FontCollection *render.FontManager
		Renderer       *render.Renderer
		Dispatcher     *event.Dispatcher
	}

	GameRegistry struct {
		State   engine.State
		Options *system.GameOptions
		Frames  *system.Frames
		World   *game.World
		Ui      *ui.UI
	}
)

func makeRegistry(flags Flags) *Registry {
	reg := &registerFactory{}

	// main
	closer := reg.registerCloser()
	sdlLib := reg.registerSDLLib(closer)

	// system
	options := reg.registerGameOptions(flags)
	frames := reg.registerFrames(options.Frames.TargetFps)

	// sdl
	sdlWindow := reg.registerSdlWindow(
		sdlLib,
	)
	sdlRenderer := reg.registerSdlRenderer(
		sdlLib,
	)

	// engine
	camera := reg.registerCamera(sdlWindow)
	mouse := reg.registerMouse()
	fontManager := reg.registerFontManager(
		closer,
	)
	textureManager := reg.registerTextureManager(
		sdlRenderer,
		closer,
	)
	renderer := reg.registerRenderer(
		sdlWindow,
		sdlRenderer,
		fontManager,
		textureManager,
		camera,
	)
	dispatcher := reg.registerDispatcher(
		reg.eventQuit(frames),
	)

	// game
	world := reg.registerWorld()

	// ui
	layerFPS := reg.registerUILayerFPS()
	gameUI := reg.registerUI(
		layerFPS,
	)

	// game state
	movement := reg.registerMovement(dispatcher)
	gameState := reg.registerGameState(
		frames,
		camera,
		mouse,
		movement,
	)

	// build
	return &Registry{
		Closer: closer,
		Engine: &EngineRegistry{
			FontCollection: fontManager,
			Renderer:       renderer,
			Dispatcher:     dispatcher,
		},
		Sdl: &SdlRegistry{
			Window: sdlWindow,
		},
		Game: &GameRegistry{
			State:   gameState,
			Options: options,
			Frames:  frames,
			World:   world,
			Ui:      gameUI,
		},
	}
}

// ----------------------------------------
// Main
// ----------------------------------------

func (r registerFactory) registerCloser() *utils.Closer {
	return utils.NewCloser()
}

// ----------------------------------------
// SDL
// ----------------------------------------

func (r registerFactory) registerSDLLib(closer *utils.Closer) *lib.SDLLib {
	sdlLib, err := lib.NewSDLLib(
		closer,
	)

	if err != nil {
		panic(fmt.Sprintf("can`t provide sdl: %v", err))
	}

	return sdlLib
}

func (r registerFactory) registerSdlWindow(sdlLib *lib.SDLLib) *sdl.Window {
	return sdlLib.Window()
}

func (r registerFactory) registerSdlRenderer(sdlLib *lib.SDLLib) *sdl.Renderer {
	return sdlLib.Renderer()
}

// ----------------------------------------
// Engine
// ----------------------------------------

func (r registerFactory) registerFontManager(closer *utils.Closer) *render.FontManager {
	fonts := render.NewFontManager(closer)
	fonts.Load(generated.ResourcesFontsJetBrainsMonoRegular)

	return fonts
}

func (r registerFactory) registerTextureManager(sdlRenderer *sdl.Renderer, closer *utils.Closer) *render.TextureManager {
	return render.NewTextureManager(sdlRenderer, closer)
}

func (r registerFactory) registerCamera(window *sdl.Window) *render.Camera {
	w, h := window.GetSize()

	return render.NewCamera(engine.Vector2D{}, int(w), int(h))
}

func (r registerFactory) registerMouse() *control.Mouse {
	return control.NewMouse()
}

func (r registerFactory) registerMovement(dispatcher *event.Dispatcher) *control.Movement {
	return control.NewMovement(dispatcher)
}

func (r registerFactory) registerRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *render.FontManager,
	textureManager *render.TextureManager,
	camera *render.Camera,
) *render.Renderer {
	return render.NewRenderer(sdlWindow, sdlRenderer, fontManager, textureManager, camera)
}

// ----------------------------------------
// Game
// ----------------------------------------

func (r registerFactory) registerGameOptions(flags Flags) *system.GameOptions {
	return &system.GameOptions{
		Debug: system.DebugOpt{
			InProfiling: flags.IsProfiling,
			System:      false,
			Frames:      false,
			World:       false,
		},
		Frames: system.FramesOpt{
			TargetFps: 30,
		},
	}
}

func (r registerFactory) registerFrames(targetFps int) *system.Frames {
	return system.NewFrames(targetFps)
}

func (r registerFactory) registerWorld() *game.World {
	return game.NewLevel01()
}

func (r registerFactory) registerUI(layers ...ui.Layer) *ui.UI {
	return ui.NewUI(layers...)
}

func (r registerFactory) registerUILayerFPS() *ui.LayerFPS {
	return ui.NewLayerFPS()
}

func (r registerFactory) registerGameState(
	moment engine.Moment,
	camera engine.Camera,
	mouse engine.Mouse,
	movement engine.Movement,
) *engine.GameState {
	return engine.NewGameState(moment, camera, mouse, movement)
}
