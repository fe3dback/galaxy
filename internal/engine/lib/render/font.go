package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/galx"
	utils2 "github.com/fe3dback/galaxy/internal/utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	ref    *ttf.Font
	closer *utils2.Closer
}

func NewFont(resource FontId, params FontParams, closer *utils2.Closer) *Font {
	defer utils2.CheckPanic(fmt.Sprintf("create font `%s`", resource))

	f, err := ttf.OpenFont(string(resource), params.size)
	utils2.Check("open", err)
	closer.EnqueueClose(func() error {
		f.Close()
		return nil
	})

	return &Font{
		ref:    f,
		closer: closer,
	}
}

func (f *Font) RenderText(text string, color galx.Color) *sdl.Surface {
	surface, err := f.ref.RenderUTF8Blended(text, transformColor(color))
	utils2.Check("render text", err)

	return surface
}
