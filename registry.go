package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/game/ui"

	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/render"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	registerFactory struct{}
	registry        struct {
		closer *utils.Closer
		sdl    *sdlRegistry
		engine *engineRegistry
		game   *gameRegistry
	}

	sdlRegistry struct {
		window *sdl.Window
	}

	engineRegistry struct {
		fontCollection *render.FontManager
		renderer       *render.Renderer
	}

	gameRegistry struct {
		options *gameOptions
		frames  *frames
		world   *game.World
		ui      *ui.UI
	}
)

func makeRegistry() *registry {
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
	options := reg.registerGameOptions()
	frames := reg.registerFrames(options.frames.targetFps)
	world := reg.registerWorld()

	// ui
	layerFPS := reg.registerUILayerFPS()
	gameUI := reg.registerUI(layerFPS)

	// build
	return &registry{
		closer: closer,
		engine: &engineRegistry{
			fontCollection: fontManager,
			renderer:       renderer,
		},
		sdl: &sdlRegistry{
			window: sdlWindow,
		},
		game: &gameRegistry{
			options: options,
			frames:  frames,
			world:   world,
			ui:      gameUI,
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
	fonts.Load(render.FontDefaultMono)

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

func (r registerFactory) registerGameOptions() *gameOptions {
	return &gameOptions{
		debug: debugOpt{
			inProfiling: *isProfiling,
			system:      false,
			frames:      true,
			world:       false,
		},
		frames: framesOpt{
			targetFps: 60,
		},
	}
}

func (r registerFactory) registerFrames(targetFps int) *frames {
	return NewFrames(targetFps)
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
