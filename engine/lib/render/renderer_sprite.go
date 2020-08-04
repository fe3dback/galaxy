package render

import (
	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (r *Renderer) internalDrawTexture(tex *sdl.Texture, src, dest sdl.Rect, angle engine.Angle) {
	err := r.ref.CopyEx(tex, &src, &dest, angle.Flip().Degrees(),
		&sdl.Point{ // point relative to dest [X,Y]
			X: dest.W / 2,
			Y: dest.H / 2,
		},
		sdl.FLIP_NONE,
	)
	utils.Check("texture copy", err)
}

func (r *Renderer) getTexture(res generated.ResourcePath) *Texture {
	return r.textureManager.Get(res)
}
