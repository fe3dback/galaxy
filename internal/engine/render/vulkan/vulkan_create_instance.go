package vulkan

import (
	"fmt"
	"log"
	"strings"

	"github.com/vulkan-go/vulkan"
)

type (
	vkInstance struct {
		ref     vulkan.Instance
		surface *vkSurface
	}
)

var requiredValidationLayers = []string{
	"VK_LAYER_KHRONOS_validation",
}

func vkCreateInstance(opts vkCreateOptions) *vkInstance {
	var inst vulkan.Instance

	// create vk instance
	vkAssert(vulkan.CreateInstance(
		vkInstanceCreateInfo(opts),
		nil,
		&inst,
	), fmt.Errorf("create vulkan instance failed"))

	opts.closer.EnqueueFree(func() {
		vulkan.DestroyInstance(inst, nil)
	})

	// return
	return &vkInstance{
		ref: inst,
	}
}

func vkSetupValidationLayers(opts vkCreateOptions) []string {
	if !opts.debugVulkan {
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
		layerLabel := vkLabelToString(vkStringToLabel(requiredLayer)) // string(any) -> [256]byte -> string(256)
		if _, exist := foundLayers[layerLabel]; !exist {
			notFound = append(notFound, layerLabel)
			continue
		}

		found = append(found, layerLabel)
	}

	if len(notFound) > 0 {
		log.Printf("vulkan debug may not work (turn off it in game config), because some of extensions not found [%s]. Available in system: [%s]",
			strings.Join(notFound, ", "),
			strings.Join(vkMapToList(foundLayers), ", "),
		)
	}

	return found
}

func vkAppInfo() *vulkan.ApplicationInfo {
	return &vulkan.ApplicationInfo{
		SType:              vulkan.StructureTypeApplicationInfo,
		PApplicationName:   "Game",
		ApplicationVersion: vulkan.MakeVersion(1, 0, 0),
		PEngineName:        "Galaxy",
		EngineVersion:      vulkan.MakeVersion(0, 0, 1),
		ApiVersion:         vulkan.ApiVersion11,
	}
}

func vkInstanceCreateInfo(opts vkCreateOptions) *vulkan.InstanceCreateInfo {
	instInfo := &vulkan.InstanceCreateInfo{
		SType:             vulkan.StructureTypeInstanceCreateInfo,
		PApplicationInfo:  vkAppInfo(),
		EnabledLayerCount: 0,
	}

	// setup extensions
	instInfo.PpEnabledExtensionNames = opts.window.GetRequiredInstanceExtensions()
	instInfo.EnabledExtensionCount = uint32(len(instInfo.PpEnabledExtensionNames))

	// setup validation (debug)
	validationLayers := vkSetupValidationLayers(opts)
	instInfo.EnabledLayerCount = uint32(len(validationLayers))
	instInfo.PpEnabledLayerNames = validationLayers

	return instInfo
}
