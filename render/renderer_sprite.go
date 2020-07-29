package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) GetTexture(res generated.ResourcePath) *Texture {
	return r.textureManager.Get(res)
}

func (r *Renderer) DrawSprite(res generated.ResourcePath, x, y int) {
	src := &Rect{X: 0, Y: 0, W: 0, H: 0}
	dest := &Rect{X: int32(x), Y: int32(y), W: 0, H: 0}

	r.draw(res, src, dest, 0)
}

func (r *Renderer) DrawSpriteEx(res generated.ResourcePath, src, dest *Rect, angle float64) {
	r.draw(res, src, dest, angle)
}

func (r *Renderer) draw(res generated.ResourcePath, src, dest *Rect, angle float64) {
	defer utils.CheckPanic(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.GetTexture(res)

	if src.W == 0 {
		src.W = texture.Width
	}
	if src.H == 0 {
		src.H = texture.Height
	}
	if dest.W == 0 {
		dest.W = texture.Width
	}
	if dest.H == 0 {
		dest.H = texture.Height
	}

	err := r.ref.CopyEx(texture.Tex, src, dest, angle, &sdl.Point{X: src.W / 2, Y: src.H / 2}, sdl.FLIP_NONE)
	utils.Check("texture copy", err)
}
