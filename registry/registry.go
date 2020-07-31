package registry

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine/render"
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
	}

	GameRegistry struct {
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

	// sdl
	sdlWindow := reg.registerSdlWindow(
		sdlLib,
	)
	sdlRenderer := reg.registerSdlRenderer(
		sdlLib,
	)

	// engine
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
	)

	// game
	options := reg.registerGameOptions(flags)
	frames := reg.registerFrames(options.Frames.TargetFps)
	world := reg.registerWorld()

	// ui
	layerFPS := reg.registerUILayerFPS()
	gameUI := reg.registerUI(layerFPS)

	// build
	return &Registry{
		Closer: closer,
		Engine: &EngineRegistry{
			FontCollection: fontManager,
			Renderer:       renderer,
		},
		Sdl: &SdlRegistry{
			Window: sdlWindow,
		},
		Game: &GameRegistry{
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

func (r registerFactory) registerSDLLib(closer *utils.Closer) *render.SDLLib {
	sdlLib, err := render.NewSDLLib(
		closer,
	)

	if err != nil {
		panic(fmt.Sprintf("can`t provide sdl: %v", err))
	}

	return sdlLib
}

func (r registerFactory) registerSdlWindow(sdlLib *render.SDLLib) *sdl.Window {
	return sdlLib.Window()
}

func (r registerFactory) registerSdlRenderer(sdlLib *render.SDLLib) *sdl.Renderer {
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

func (r registerFactory) registerRenderer(
	sdlWindow *sdl.Window,
	sdlRenderer *sdl.Renderer,
	fontManager *render.FontManager,
	textureManager *render.TextureManager,
) *render.Renderer {
	return render.NewRenderer(sdlWindow, sdlRenderer, fontManager, textureManager)
}

// ----------------------------------------
// Game
// ----------------------------------------

func (r registerFactory) registerGameOptions(flags Flags) *system.GameOptions {
	return &system.GameOptions{
		Debug: system.DebugOpt{
			InProfiling: flags.IsProfiling,
			System:      false,
			Frames:      true,
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

func (r registerFactory) registerWorld() *game.World {
	return game.NewLevel01()
}

func (r registerFactory) registerUI(layers ...ui.Layer) *ui.UI {
	return ui.NewUI(layers...)
}

func (r registerFactory) registerUILayerFPS() *ui.LayerFPS {
	return ui.NewLayerFPS()
}
