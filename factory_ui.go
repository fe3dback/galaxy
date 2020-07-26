package main

import (
	"github.com/fe3dback/galaxy/game/ui"
)

func (f *factory) provideUi() *ui.UI {
	return ui.NewUI(
		f.provideUiLayerFPS(),
	)
}

func (f *factory) provideUiLayerFPS() *ui.LayerFPS {
	return ui.NewLayerFPS(
		f.provideRenderer(),
		f.provideFrames(),
	)
}
