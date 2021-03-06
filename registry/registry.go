package registry

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fe3dback/galaxy/engine/physics"

	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/engine/lib/sound"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/editor"
	editorcomponents "github.com/fe3dback/galaxy/editor/components"
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/engine/control"
	engineeditor "github.com/fe3dback/galaxy/engine/editor"
	"github.com/fe3dback/galaxy/engine/event"
	"github.com/fe3dback/galaxy/engine/lib"
	"github.com/fe3dback/galaxy/engine/lib/render"
	"github.com/fe3dback/galaxy/engine/loader"
	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/shared/ui"
	"github.com/fe3dback/galaxy/system"
	"github.com/fe3dback/galaxy/utils"
)

type (
	registerFactory struct{}
	Registry        struct {
		Closer *utils.Closer
		Logger *zap.SugaredLogger
		Sdl    *SdlRegistry
		Engine *EngineRegistry
		Game   *GameRegistry
		Editor *EditorRegistry
		State  engine.State
	}

	SdlRegistry struct {
		Window *sdl.Window
	}

	EngineRegistry struct {
		FontCollection *render.FontManager
		Renderer       *render.Renderer
		Dispatcher     *event.Dispatcher
		AppState       *engine.AppState
	}

	GameRegistry struct {
		Options      *system.GameOptions
		Frames       *system.Frames
		WorldManager *game.WorldManager
		Ui           *ui.UI
	}

	EditorRegistry struct {
		Manager *editor.Manager
		Ui      *ui.UI
	}
)

func makeRegistry(flags Flags) *Registry {
	reg := &registerFactory{}

	// core rand seed
	rand.Seed(flags.Seed)

	// main
	closer := reg.registerCloser()
	logger := reg.registerLogger(closer)
	sdlLib := reg.registerSDLLib(closer, flags.FullScreen)
	logger.Infow("Initialized random seed",
		zap.Int64("seed", flags.Seed),
	)

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
	soundManager := reg.registerSoundManager(closer)
	appState := reg.registerAppState()
	dispatcher := reg.registerDispatcher(
		reg.eventQuit(frames),
		reg.eventSwitchEditorState(appState),
	)
	camera := reg.registerCamera(dispatcher)
	mouse := reg.registerMouse(dispatcher)
	renderGizmos := reg.registerRenderGizmos(dispatcher, options.Debug.System)
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
		dispatcher,
		renderGizmos,
		appState,
	)

	// shared ui
	uiLayerSharedFPS := reg.registerUILayerSharedFPS()

	// editor
	editorManager := reg.registerEditorManager()

	// editor ui
	editorUI := reg.registerUI(
		uiLayerSharedFPS,
	)

	// physics
	physicsWorld := reg.registerPhysics(closer)

	// game
	assetsLoader := reg.registerAssetsLoader(soundManager)
	worldCreator := reg.registerGameWorldCreator(assetsLoader, soundManager, physicsWorld)
	worldManager := reg.registerWorldManager(worldCreator, dispatcher)

	// game ui
	gameUI := reg.registerUI(
		uiLayerSharedFPS,
	)

	// game state
	movement := reg.registerMovement(dispatcher)
	gameState := reg.registerGameState(
		frames,
		camera,
		mouse,
		movement,
		appState,
		soundManager,
	)

	// build
	return &Registry{
		Closer: closer,
		Logger: logger,
		Engine: &EngineRegistry{
			FontCollection: fontManager,
			Renderer:       renderer,
			Dispatcher:     dispatcher,
			AppState:       appState,
		},
		Sdl: &SdlRegistry{
			Window: sdlWindow,
		},
		State: gameState,
		Game: &GameRegistry{
			Options:      options,
			Frames:       frames,
			WorldManager: worldManager,
			Ui:           gameUI,
		},
		Editor: &EditorRegistry{
			Manager: editorManager,
			Ui:      editorUI,
		},
	}
}

// ----------------------------------------
// Main
// ----------------------------------------

func (r registerFactory) registerCloser() *utils.Closer {
	return utils.NewCloser()
}

func (r registerFactory) registerLogger(closer *utils.Closer) *zap.SugaredLogger {
	const (
		logPathDebug  = "./galaxy-debug.log"
		logPathErrors = "./galaxy-error.log"
	)

	_ = os.Remove(logPathDebug)
	_ = os.Remove(logPathErrors)

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout", logPathDebug},
		ErrorOutputPaths:  []string{"stderr", logPathErrors},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("failed to create logger")
	}

	closedStd := zap.RedirectStdLog(logger)
	restoreLogger := zap.ReplaceGlobals(logger)

	closer.EnqueueClose(logger.Sync)
	closer.EnqueueFree(closedStd)
	closer.EnqueueFree(restoreLogger)

	return logger.Sugar()
}

// ----------------------------------------
// SDL
// ----------------------------------------

func (r registerFactory) registerSDLLib(closer *utils.Closer, fullscreen bool) *lib.SDLLib {
	sdlLib, err := lib.NewSDLLib(
		closer,
		960,
		540,
		fullscreen,
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

func (r registerFactory) registerSoundManager(closer *utils.Closer) *sound.Manager {
	return sound.NewManager(closer)
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

func (r registerFactory) registerCamera(dispatcher *event.Dispatcher) *render.Camera {
	return render.NewCamera(dispatcher)
}

func (r registerFactory) registerMouse(dispatcher *event.Dispatcher) *control.Mouse {
	return control.NewMouse(dispatcher)
}

func (r registerFactory) registerMovement(dispatcher *event.Dispatcher) *control.Movement {
	return control.NewMovement(dispatcher)
}

func (r registerFactory) registerRenderGizmos(dispatcher *event.Dispatcher, debugMode bool) engine.Gizmos {
	return engineeditor.NewDrawGizmos(dispatcher, debugMode)
}

func (r registerFactory) registerRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *render.FontManager,
	textureManager *render.TextureManager,
	camera *render.Camera,
	dispatcher *event.Dispatcher,
	gizmos engine.Gizmos,
	appState *engine.AppState,
) *render.Renderer {
	return render.NewRenderer(
		sdlWindow,
		sdlRenderer,
		fontManager,
		textureManager,
		camera,
		dispatcher,
		gizmos,
		appState,
	)
}

func (r registerFactory) registerAppState() *engine.AppState {
	return engine.NewAppState()
}

// ----------------------------------------
// Editor
// ----------------------------------------

func (r registerFactory) registerEditorManager() *editor.Manager {
	// todo auto register all editor components
	components := make([]editor.Component, 0)

	// register
	components = append(components, editorcomponents.NewCamera())

	return editor.NewManager(components)
}

// ----------------------------------------
// Game
// ----------------------------------------

func (r registerFactory) registerGameOptions(flags Flags) *system.GameOptions {
	return &system.GameOptions{
		Debug: system.DebugOpt{
			InProfiling: flags.IsProfiling,
			System:      true,
			Memory:      false,
			Frames:      false,
			World:       false,
		},
		Frames: system.FramesOpt{
			TargetFps: 60,
		},
	}
}

func (r registerFactory) registerFrames(targetFps int) *system.Frames {
	return system.NewFrames(targetFps)
}

func (r registerFactory) registerAssetsLoader(soundManager *sound.Manager) *loader.AssetsLoader {
	return loader.NewAssetsLoader(
		soundManager,
	)
}

func (r registerFactory) registerPhysics(closer *utils.Closer) *physics.World {
	return physics.NewWorld(closer)
}

func (r registerFactory) registerGameWorldCreator(
	assetsLoader engine.Loader,
	soundManager engine.SoundMixer,
	physics engine.Physics,
) *engine.GameWorldCreator {
	return engine.NewGameWorldCreator(
		assetsLoader,
		soundManager,
		physics,
	)
}

func (r registerFactory) registerWorldManager(worldCreator engine.WorldCreator, dispatcher *event.Dispatcher) *game.WorldManager {
	return game.NewWorldManager(game.NewLevel01(), worldCreator, dispatcher)
}

func (r registerFactory) registerUI(layers ...ui.Layer) *ui.UI {
	return ui.NewUI(layers...)
}

func (r registerFactory) registerUILayerSharedFPS() *ui.LayerFPS {
	return ui.NewLayerSharedFPS()
}

func (r registerFactory) registerGameState(
	moment engine.Moment,
	camera engine.Camera,
	mouse engine.Mouse,
	movement engine.Movement,
	appState *engine.AppState,
	soundMixer *sound.Manager,
) *engine.GameState {
	return engine.NewGameState(moment, camera, mouse, movement, appState, soundMixer)
}
