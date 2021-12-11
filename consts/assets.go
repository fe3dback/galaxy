package consts

// engine wide hardcoded assets required for start up
const (
	// AssetDefaultFont used for rendering debug/engine stats
	AssetDefaultFont AssetsPath = "assets/default-font.ttf"

	// AssetScenesRoot point to scenes root directory
	AssetScenesRoot AssetsPath = "assets/scenes"

	// AssetScenesDefinitionXML point to xml file with scenes refs and settings
	AssetScenesDefinitionXML = AssetScenesRoot + "/scenes.xml"

	// AssetScenesObjectsFileName is file name where scene objects is defined
	AssetScenesObjectsFileName = "objects.xml"
)

type (
	// AssetsPath is game relative path to asset on disk
	AssetsPath = string
)
