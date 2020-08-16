package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/engine"
	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	ref    *ttf.Font
	closer *utils.Closer
}

func NewFont(resource FontId, params FontParams, closer *utils.Closer) *Font {
	defer utils.CheckPanic(fmt.Sprintf("create font `%s`", resource))

	f, err := ttf.OpenFont(string(resource), params.size)
	utils.Check("open", err)
	closer.EnqueueClose(func() error {
		f.Close()
		return nil
	})

	return &Font{
		ref:    f,
		closer: closer,
	}
}

func (f *Font) RenderText(text string, color engine.Color) *sdl.Surface {
	surface, err := f.ref.RenderUTF8Blended(text, transformColor(color))
	utils.Check("render text", err)

	return surface
}
