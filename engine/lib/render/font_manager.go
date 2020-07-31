package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
)

var FontVars = FontsParamsMap{
	generated.ResourcesFontsJetBrainsMonoRegular: FontParams{
		size: 14,
	},
}

type (
	FontId     = generated.ResourcePath
	FontParams struct {
		size int
	}
	FontsMap       map[FontId]*Font
	FontsParamsMap map[FontId]FontParams

	FontManager struct {
		fonts  FontsMap
		closer *utils.Closer
	}
)

func NewFontManager(closer *utils.Closer) *FontManager {
	return &FontManager{
		fonts:  map[FontId]*Font{},
		closer: closer,
	}
}

func (fc *FontManager) Load(id FontId) {
	if _, ok := fc.fonts[id]; ok {
		panic(fmt.Sprintf("font `%s` already loadedTextures", id))
	}

	params, ok := FontVars[id]
	if !ok {
		panic(fmt.Sprintf("font `%s` params not defined", id))
	}

	fc.fonts[id] = NewFont(id, params, fc.closer)
}

func (fc *FontManager) Get(id FontId) *Font {
	if font, ok := fc.fonts[id]; ok {
		return font
	}

	panic(fmt.Sprintf("font `%s` not loadedTextures", id))
}
