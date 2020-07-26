package main

import (
	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/game/ui"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	gameParams struct {
		closer  *utils.Closer
		window  *sdl.Window
		world   *game.World
		ui      *ui.UI
		frames  *frames
		options *gameOptions
	}

	gameOptions struct {
		debug  debugOpt
		frames framesOpt
	}

	debugOpt struct {
		inProfiling bool
		system      bool
		frames      bool
		world       bool
	}

	framesOpt struct {
		targetFps int64
	}
)

func newGameParams(
	closer *utils.Closer,
	window *sdl.Window,
	world *game.World,
	ui *ui.UI,
	frames *frames,
	options *gameOptions,
) *gameParams {
	return &gameParams{
		closer:  closer,
		window:  window,
		world:   world,
		ui:      ui,
		frames:  frames,
		options: options,
	}
}
