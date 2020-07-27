package render

import (
	"fmt"

	"github.com/fe3dback/galaxy/utils"
)

var FontVars = FontsParamsMap{
	FontDefaultMono: FontParams{
		size: 14,
	},
}

const (
	FontDefaultMono FontId = "jet_brains_mono_regular"
)

type (
	FontId     string
	FontParams struct {
		size int
	}
	FontsMap        map[FontId]*Font
	FontsParamsMap  map[FontId]FontParams
	FontsCollection struct {
		fonts              FontsMap
		fontsDirectoryPath string
		closer             *utils.Closer
	}
)

func NewFontsCollection(fontsDirectoryPath string, closer *utils.Closer) *FontsCollection {
	return &FontsCollection{
		fonts:              map[FontId]*Font{},
		fontsDirectoryPath: fontsDirectoryPath,
		closer:             closer,
	}
}

func (fc *FontsCollection) Load(id FontId) {
	if _, ok := fc.fonts[id]; ok {
		panic(fmt.Sprintf("font `%s` already loaded", id))
	}

	params, ok := FontVars[id]
	if !ok {
		panic(fmt.Sprintf("font `%s` params not defined", id))
	}

	path := fmt.Sprintf("%s/%s.ttf", fc.fontsDirectoryPath, string(id))
	f, err := NewFont(path, fc.closer, params)
	if err != nil {
		panic(fmt.Sprintf("font `%s` loading failed: %v", id, err))
	}

	fc.fonts[id] = f
}

func (fc *FontsCollection) Get(id FontId) *Font {
	if font, ok := fc.fonts[id]; ok {
		return font
	}

	panic(fmt.Sprintf("font `%s` not loaded", id))
}
