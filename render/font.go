package render

import (
	"fmt"
	"image/color"

	"github.com/fe3dback/galaxy/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	ref    *ttf.Font
	closer *utils.Closer
}

func NewFont(path string, closer *utils.Closer, params FontParams) (font *Font, err error) {
	utils.Recover(fmt.Sprintf("create font `%s`", path), &err)

	f, err := ttf.OpenFont(path, params.size)
	utils.Check("open", err)
	closer.Enqueue(func() error {
		f.Close()
		return nil
	})

	return &Font{
		ref:    f,
		closer: closer,
	}, nil
}

func (f *Font) RenderText(text string, color color.RGBA) *sdl.Surface {
	surface, err := f.ref.RenderUTF8Blended(text, sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	})
	utils.Check("render text", err)

	return surface
}
