package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) TextureQuery(res generated.ResourcePath) engine.TextureInfo {
	tex := r.getTexture(res)

	return engine.TextureInfo{
		Width:  int(tex.Width),
		Height: int(tex.Height),
	}
}

func (r *Renderer) DrawSprite(res generated.ResourcePath, p engine.Point) {
	src := engine.Rect{X: 0, Y: 0, W: 0, H: 0}
	dest := engine.Rect{X: p.X, Y: p.Y, W: 0, H: 0}

	r.draw(res, src, dest, 0)
}

func (r *Renderer) DrawSpriteEx(res generated.ResourcePath, src, dest engine.Rect, angle float64) {
	r.draw(res, src, dest, angle)
}

func (r *Renderer) draw(res generated.ResourcePath, src, dest engine.Rect, angle float64) {
	defer utils.CheckPanic(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.getTexture(res)

	if dest.W == 0 {
		dest.W = int(texture.Width)
	}
	if dest.H == 0 {
		dest.H = int(texture.Height)
	}

	if !r.isRectInsideCamera(dest) {
		return
	}

	if src.W == 0 {
		src.W = int(texture.Width)
	}
	if src.H == 0 {
		src.H = int(texture.Height)
	}

	err := r.ref.CopyEx(
		texture.Tex,
		&sdl.Rect{
			X: int32(src.X),
			Y: int32(src.Y),
			W: int32(src.W),
			H: int32(src.H),
		},
		r.screenRectPtr(dest),
		angle,
		&sdl.Point{ // point relative to dest [X,Y]
			X: int32(src.W / 2),
			Y: int32(src.H / 2),
		},
		sdl.FLIP_NONE,
	)
	utils.Check("texture copy", err)
}

func (r *Renderer) getTexture(res generated.ResourcePath) *Texture {
	return r.textureManager.Get(res)
}
