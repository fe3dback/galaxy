package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
)

type (
	Texture struct {
		Path   consts.AssetsPath
		Tex    *sdl.Texture
		Width  int32
		Height int32
	}
)
