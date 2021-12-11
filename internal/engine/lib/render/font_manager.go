package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/consts"
	"github.com/fe3dback/galaxy/internal/utils"
)

var FontVars = FontsParamsMap{
	consts.AssetDefaultFont: FontParams{
		size: 14,
	},
}

type (
	FontId     = consts.AssetsPath
	FontParams struct {
		size int
	}
	FontsMap       map[FontId]*Font
	FontsParamsMap map[FontId]FontParams

	FontsManager struct {
		fonts  FontsMap
		closer *utils.Closer
	}
)

func NewFontsManager(closer *utils.Closer) *FontsManager {
	return &FontsManager{
		fonts:  map[FontId]*Font{},
		closer: closer,
	}
}

func (fc *FontsManager) Load(id FontId) {
	if _, ok := fc.fonts[id]; ok {
		panic(fmt.Sprintf("font `%s` already loadedTextures", id))
	}

	params, ok := FontVars[id]
	if !ok {
		panic(fmt.Sprintf("font `%s` params not defined", id))
	}

	fc.fonts[id] = NewFont(id, params, fc.closer)
}

func (fc *FontsManager) Get(id FontId) *Font {
	if font, ok := fc.fonts[id]; ok {
		return font
	}

	panic(fmt.Sprintf("font `%s` not loadedTextures", id))
}
