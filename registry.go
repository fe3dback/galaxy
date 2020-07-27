package main

import (
	"fmt"
	"os"
	"path/filepath"

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
		engine *engineRegistry
		sdl    *sdlRegistry
		game   *gameRegistry
	}

	engineRegistry struct {
		fontCollection *render.FontsCollection
		renderer       *render.Renderer
	}

	sdlRegistry struct {
		window *sdl.Window
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

	// engine
	fontsCollection := reg.registerFontsCollection(
		reg.dirFonts(),
		closer,
	)

	// sdl
	window := reg.registerWindow(
		sdlLib,
	)
	renderer := reg.registerRenderer(
		window,
		fontsCollection,
		closer,
	)

	// game
	options := reg.registerGameOptions()
	frames := reg.registerFrames(options.frames.targetFps)
	world := reg.registerWorld()

	// ui
	layerFPS := reg.registerUILayerFPS(renderer, frames)
	gameUI := reg.registerUI(layerFPS)

	// build
	return &registry{
		closer: closer,
		engine: &engineRegistry{
			fontCollection: fontsCollection,
			renderer:       renderer,
		},
		sdl: &sdlRegistry{
			window: window,
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
// Engine
// ----------------------------------------

func (r registerFactory) registerFontsCollection(fontsDir string, closer *utils.Closer) *render.FontsCollection {
	fonts := render.NewFontsCollection(fontsDir, closer)
	fonts.Load(render.FontDefaultMono)

	return fonts
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

func (r registerFactory) registerWindow(sdlLib *render.SDLLib) *sdl.Window {
	return sdlLib.Window()
}

func (r registerFactory) registerRenderer(
	window *sdl.Window,
	collection *render.FontsCollection,
	closer *utils.Closer,
) *render.Renderer {
	return render.NewRenderer(window, collection, closer)
}

// ----------------------------------------
// Game
// ----------------------------------------

func (r registerFactory) registerGameOptions() *gameOptions {
	return &gameOptions{
		debug: debugOpt{
			inProfiling: *isProfiling,
			system:      true,
			frames:      false,
			world:       true,
		},
		frames: framesOpt{
			targetFps: 60,
		},
	}
}

func (r registerFactory) registerFrames(targetFps int64) *frames {
	return NewFrames(targetFps)
}

func (r registerFactory) registerWorld() *game.World {
	return game.NewLevel01()
}

func (r registerFactory) registerUI(layers ...ui.Layer) *ui.UI {
	return ui.NewUI(layers...)
}

func (r registerFactory) registerUILayerFPS(renderer *render.Renderer, frm *frames) *ui.LayerFPS {
	return ui.NewLayerFPS(
		renderer,
		frm,
	)
}

// ----------------------------------------
// Path
// ----------------------------------------

func (r registerFactory) dirRoot() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Sprintf("can`t provide root dir: %v", err))
	}

	return dir
}

func (r registerFactory) dirResources() string {
	return fmt.Sprintf("%s/resources", r.dirRoot())
}

func (r registerFactory) dirFonts() string {
	return fmt.Sprintf("%s/fonts", r.dirResources())
}
