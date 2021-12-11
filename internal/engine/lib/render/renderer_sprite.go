package render

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/galx"
	"github.com/fe3dback/galaxy/internal/utils"
)

func (r *Renderer) internalDrawTexture(tex *sdl.Texture, src, dest sdl.Rect, angle galx.Angle) {
	err := r.ref.CopyEx(tex, &src, &dest, angle.Flip().Degrees(),
		&sdl.Point{ // point relative to dest [X,Y]
			X: dest.W / 2,
			Y: dest.H / 2,
		},
		sdl.FLIP_NONE,
	)
	utils.Check("texture copy", err)
}

func (r *Renderer) getTexture(res consts.AssetsPath) *Texture {
	return r.textureManager.Get(res)
}
