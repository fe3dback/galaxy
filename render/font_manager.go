package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/generated"
	"github.com/fe3dback/galaxy/utils"
)

const (
	FontDefaultMono FontId = "jet_brains_mono_regular"
)

var FontVars = FontsParamsMap{
	FontDefaultMono: FontParams{
		resourcePath: generated.ResourcesFontsJetBrainsMonoRegular,
		size:         14,
	},
}

type (
	FontId     string
	FontParams struct {
		resourcePath generated.ResourcePath
		size         int
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

	f, err := NewFont(params, fc.closer)
	if err != nil {
		panic(fmt.Sprintf("font `%s` loading failed: %v", id, err))
	}

	fc.fonts[id] = f
}

func (fc *FontManager) Get(id FontId) *Font {
	if font, ok := fc.fonts[id]; ok {
		return font
	}

	panic(fmt.Sprintf("font `%s` not loadedTextures", id))
}
