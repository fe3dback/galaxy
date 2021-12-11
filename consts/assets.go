package consts

// engine wide hardcoded assets required for start up
const (
	// DefaultFont used for rendering debug/engine stats
	DefaultFont = "assets/default-font.ttf"

	// ScenesDefinitionXML point to xml file with scenes refs and settings
	ScenesDefinitionXML = "assets/scenes/scenes.xml"
)

type (
	// AssetsPath is game relative path to asset on disk
	AssetsPath = string
)
