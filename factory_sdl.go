package main

import (
	"fmt"

	"github.com/fe3dback/galaxy/render"
	"github.com/veandco/go-sdl2/sdl"
)

func (f *factory) provideSdl() *render.SDLLib {
	sdlLib, err := render.NewSDLLib(
		f.provideCloser(),
	)

	if err != nil {
		panic(fmt.Sprintf("can`t provide sdl: %v", err))
	}

	return sdlLib
}

func (f *factory) provideWindow() *sdl.Window {
	return f.provideSdl().Window()
}

func (f *factory) provideRenderer() *render.Renderer {
	return render.NewRenderer(
		f.provideWindow(),
		f.provideFontsCollection(),
		f.provideCloser(),
	)
}

func (f *factory) provideFontsCollection() *render.FontsCollection {
	fonts := render.NewFontsCollection(
		f.provideFontsDirectory(),
		f.provideCloser(),
	)

	fonts.Load(render.FontDefaultMono)

	return fonts
}
