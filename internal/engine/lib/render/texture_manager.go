package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	utils2 "github.com/fe3dback/galaxy/internal/utils"
)

type TextureManager struct {
	sdlRenderer    *sdl.Renderer
	closer         *utils2.Closer
	loadedTextures map[consts.AssetsPath]*Texture
}

func NewTextureManager(sdlRenderer *sdl.Renderer, closer *utils2.Closer) *TextureManager {
	return &TextureManager{
		sdlRenderer:    sdlRenderer,
		closer:         closer,
		loadedTextures: map[consts.AssetsPath]*Texture{},
	}
}

func (manager *TextureManager) Get(path consts.AssetsPath) *Texture {
	if texture, ok := manager.loadedTextures[path]; ok {
		return texture
	}

	return manager.Load(path)
}

func (manager *TextureManager) Load(path consts.AssetsPath) *Texture {
	defer utils2.CheckPanic(fmt.Sprintf("texture `%s`", path))

	if _, ok := manager.loadedTextures[path]; ok {
		panic("texture already loadedTextures")
	}

	tex, err := img.LoadTexture(manager.sdlRenderer, string(path))
	utils2.Check("load", err)
	manager.closer.EnqueueClose(tex.Destroy)

	_, _, w, h, err := tex.Query()
	utils2.Check("query", err)

	texture := &Texture{
		Path:   path,
		Tex:    tex,
		Width:  w,
		Height: h,
	}

	manager.loadedTextures[path] = texture
	return texture
}
