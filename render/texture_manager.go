package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	sdlRenderer    *sdl.Renderer
	closer         *utils.Closer
	loadedTextures map[generated.ResourcePath]*Texture
}

func NewTextureManager(sdlRenderer *sdl.Renderer, closer *utils.Closer) *TextureManager {
	return &TextureManager{
		sdlRenderer:    sdlRenderer,
		closer:         closer,
		loadedTextures: map[generated.ResourcePath]*Texture{},
	}
}

func (manager *TextureManager) Get(path generated.ResourcePath) *Texture {
	if texture, ok := manager.loadedTextures[path]; ok {
		return texture
	}

	return manager.Load(path)
}

func (manager *TextureManager) Load(path generated.ResourcePath) *Texture {
	utils.PanicContext(fmt.Sprintf("texture `%s`", path))

	if _, ok := manager.loadedTextures[path]; ok {
		panic("texture already loadedTextures")
	}

	tex, err := img.LoadTexture(manager.sdlRenderer, string(path))
	utils.Check("load", err)
	manager.closer.Enqueue(tex.Destroy)

	_, _, w, h, err := tex.Query()
	utils.Check("query", err)

	texture := &Texture{
		Path:   path,
		Tex:    tex,
		Width:  w,
		Height: h,
	}

	manager.loadedTextures[path] = texture
	return texture
}
