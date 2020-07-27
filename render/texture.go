package render

import (
	"github.com/fe3dback/galaxy/generated"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Texture struct {
		Path   generated.ResourcePath
		Tex    *sdl.Texture
		Width  int32
		Height int32
	}
)
