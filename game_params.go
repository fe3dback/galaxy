package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/render"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	initializer func(*gameParams)

	gameParams struct {
		options gameOptions
		closer  *closer
		window  *sdl.Window
		world   *game.World
		frames  *frames
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

func newGame() *gameParams {
	return newGameParams(
		withSettings,
		withDebug,
		withWindow,
		withFrames,
		withWorld,
	)
}

func newGameParams(args ...initializer) *gameParams {
	params := &gameParams{
		closer: newCloser(),
	}

	for _, initializer := range args {
		initializer(params)
	}

	return params
}

func withSettings(game *gameParams) {
	game.options.frames = framesOpt{
		targetFps: 60,
	}
}

func withDebug(game *gameParams) {
	game.options.debug = debugOpt{
		inProfiling: *isProfiling,
		system:      true,
		frames:      false,
		world:       true,
	}
}

func withWindow(params *gameParams) {
	lib, err := render.CreateSDL()
	if err != nil {
		panic(fmt.Sprintf("can`t create sdl: %v", err))
	}

	params.closer.Enqueue(lib.Close)
	params.window = lib.Window()
}

func withWorld(params *gameParams) {
	params.world = game.NewLevel01()
}

func withFrames(params *gameParams) {
	params.frames = NewFrames(params.options.frames.targetFps)
}
