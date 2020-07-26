package main

import (
	"github.com/fe3dback/galaxy/game"
	"github.com/fe3dback/galaxy/utils"
)

type factory struct {
}

func newFactory() *factory {
	return &factory{}
}

func (f *factory) provideGameOptions() *gameOptions {
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

func (f *factory) provideCloser() *utils.Closer {
	return utils.NewCloser()
}

func (f *factory) provideFrames() *frames {
	opt := f.provideGameOptions()

	return NewFrames(opt.frames.targetFps)
}

func (f *factory) provideWorld() *game.World {
	return game.NewLevel01()
}

func (f *factory) provideGameParams() *gameParams {
	return newGameParams(
		f.provideCloser(),
		f.provideWindow(),
		f.provideWorld(),
		f.provideUi(),
		f.provideFrames(),
		f.provideGameOptions(),
	)
}
