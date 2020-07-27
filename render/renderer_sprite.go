package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
)

func (r *Renderer) DrawSprite(res generated.ResourcePath, x, y int32) {
	utils.PanicContext(fmt.Sprintf("draw sprite `%s`", res))

	texture := r.textureManager.Get(res)

	src := &Rect{
		X: 0,
		Y: 0,
		W: texture.Width,
		H: texture.Height,
	}

	dest := &Rect{
		X: x,
		Y: y,
		W: texture.Width,
		H: texture.Height,
	}

	err := r.ref.Copy(texture.Tex, src, dest)
	utils.Check("texture copy", err)
}
