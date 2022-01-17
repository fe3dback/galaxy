package vulkan

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"
)

const (
	applicationName = "Game"
	engineName      = "Galaxy"
)

var requiredValidationLayers = []string{
	"VK_LAYER_KHRONOS_validation",
}

func newVkInstance(requiredExt []string, debugMode bool) *vkInstance {
	var inst vulkan.Instance

	vkAssert(
		vulkan.CreateInstance(instanceCreateInfo(requiredExt, debugMode), nil, &inst),
		fmt.Errorf("create vulkan instance failed"),
	)

	return &vkInstance{
		ref: inst,
	}
}

func (inst *vkInstance) free() {
	vulkan.DestroyInstance(inst.ref, nil)

	log.Printf("vk: freed: vulkan instance\n")
}

func instanceCreateInfo(requiredExt []string, debugMode bool) *vulkan.InstanceCreateInfo {
	log.Printf("vk: init '%s', required extensions: [%v]\n", engineName, requiredExt)

	instInfo := &vulkan.InstanceCreateInfo{
		SType: vulkan.StructureTypeInstanceCreateInfo,
		PApplicationInfo: &vulkan.ApplicationInfo{
			SType:              vulkan.StructureTypeApplicationInfo,
			PApplicationName:   applicationName,
			ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
			PEngineName:        engineName,
			EngineVersion:      vulkan.MakeVersion(1, 0, 0),
			ApiVersion:         vulkan.ApiVersion11,
		},
	}

	// setup extensions
	instInfo.PpEnabledExtensionNames = requiredExt
	instInfo.EnabledExtensionCount = uint32(len(instInfo.PpEnabledExtensionNames))

	// setup validation (debug)
	validationLayers := validationLayers(debugMode)
	instInfo.EnabledLayerCount = uint32(len(validationLayers))
	instInfo.PpEnabledLayerNames = validationLayers

	return instInfo
}

func validationLayers(isDebugMode bool) []string {
	if !isDebugMode {
		return []string{}
	}

	layersCount := uint32(0)
	vkAssert(
		vulkan.EnumerateInstanceLayerProperties(&layersCount, nil),
		fmt.Errorf("failed enumerate layer properties"),
	)

	availableLayers := make([]vulkan.LayerProperties, layersCount)
	vkAssert(
		vulkan.EnumerateInstanceLayerProperties(&layersCount, availableLayers),
		fmt.Errorf("failed enumerate layer properties"),
	)

	foundLayers := make(map[string]struct{})
	for _, layer := range availableLayers {
		layer.Deref()
		foundLayers[vkLabelToString(layer.LayerName)] = struct{}{}
	}

	notFound := make([]string, 0)
	found := make([]string, 0)

	for _, requiredLayer := range requiredValidationLayers {
		layerLabel := vkRepackLabel(requiredLayer)
		if _, exist := foundLayers[layerLabel]; !exist {
			notFound = append(notFound, layerLabel)
			continue
		}

		found = append(found, layerLabel)
	}

	log.Printf("vk: available layers: [%v]\n", found)

	if len(notFound) > 0 {
		log.Printf("vk: debug may not work (turn off it in game config), because some of extensions not found: %v\n",
			notFound,
		)
	}

	return found
}
