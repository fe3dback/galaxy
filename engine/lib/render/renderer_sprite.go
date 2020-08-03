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

func (r *Renderer) DrawSprite(res generated.ResourcePath, vec engine.Vec) {
	src := engine.Rect{}
	dest := engine.Rect{Min: vec}

	r.draw(res, src, dest, 0)
}

func (r *Renderer) DrawSpriteAngle(res generated.ResourcePath, vec engine.Vec, angle engine.Angle) {
	src := engine.Rect{}
	dest := engine.Rect{Min: vec}

	r.draw(res, src, dest, angle)
}

func (r *Renderer) DrawSpriteEx(res generated.ResourcePath, src, dest engine.Rect, angle engine.Angle) {
	r.draw(res, src, dest, angle)
}

func (r *Renderer) draw(res generated.ResourcePath, src, dest engine.Rect, angle engine.Angle) {
	defer utils.CheckPanic(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.getTexture(res)

	if dest.Max.X == 0 {
		dest.Max.X = float64(texture.Width)
	}
	if dest.Max.Y == 0 {
		dest.Max.Y = float64(texture.Height)
	}

	// apply offset tex to dest
	dest.Min.X -= dest.Max.X / 2
	dest.Min.Y -= dest.Max.Y / 2

	// check is visible
	if !r.isRectInsideCamera(dest) {
		return
	}

	if src.Max.X == 0 {
		src.Max.X = float64(texture.Width)
	}
	if src.Max.Y == 0 {
		src.Max.Y = float64(texture.Height)
	}

	err := r.ref.CopyEx(
		texture.Tex,
		&sdl.Rect{
			X: int32(src.Min.X),
			Y: int32(src.Min.Y),
			W: int32(src.Max.X),
			H: int32(src.Max.Y),
		},
		r.screenRectPtr(dest),
		angle.Degrees(),
		&sdl.Point{ // point relative to dest [X,Y]
			X: int32(src.Max.X / 2),
			Y: int32(src.Max.Y / 2),
		},
		sdl.FLIP_NONE,
	)
	utils.Check("texture copy", err)
}

func (r *Renderer) getTexture(res generated.ResourcePath) *Texture {
	return r.textureManager.Get(res)
}
