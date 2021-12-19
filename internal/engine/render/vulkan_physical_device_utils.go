package render

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

func (inst *vkInstance) vkPhysicalDeviceGetExtensions(refDevice vulkan.PhysicalDevice) []vulkan.ExtensionProperties {
	count := uint32(0)
	vkAssert(
		vulkan.EnumerateDeviceExtensionProperties(refDevice, "", &count, nil),
		fmt.Errorf("failed enumerate device extensions"),
	)
	if count == 0 {
		return nil
	}

	ext := make([]vulkan.ExtensionProperties, count)
	vkAssert(
		vulkan.EnumerateDeviceExtensionProperties(refDevice, "", &count, ext),
		fmt.Errorf("failed enumerate device extensions"),
	)

	result := make([]vulkan.ExtensionProperties, 0, count)
	for _, properties := range ext {
		properties.Deref()
		result = append(result, properties)
	}

	return result
}

func (inst *vkInstance) vkPhysicalDeviceGetSurfaceCapabilities(device vulkan.PhysicalDevice) vkPhysicalDeviceSurface {
	// Capabilities
	var capabilities vulkan.SurfaceCapabilities
	vkAssert(
		vulkan.GetPhysicalDeviceSurfaceCapabilities(device, inst.surface.ref, &capabilities),
		fmt.Errorf("failed get physical device capabilities"),
	)

	// Surface formats
	formatsCount := uint32(0)
	vkAssert(
		vulkan.GetPhysicalDeviceSurfaceFormats(device, inst.surface.ref, &formatsCount, nil),
		fmt.Errorf("failed get physical device surface formats"),
	)

	surfaceFormats := make([]vulkan.SurfaceFormat, formatsCount)
	vkAssert(
		vulkan.GetPhysicalDeviceSurfaceFormats(device, inst.surface.ref, &formatsCount, surfaceFormats),
		fmt.Errorf("failed get physical device surface formats"),
	)

	for _, format := range surfaceFormats {
		format.Deref()
	}

	// Present Modes
	modesCount := uint32(0)
	vkAssert(
		vulkan.GetPhysicalDeviceSurfacePresentModes(device, inst.surface.ref, &modesCount, nil),
		fmt.Errorf("failed get physical device present modes"),
	)

	presentModes := make([]vulkan.PresentMode, modesCount)
	vkAssert(
		vulkan.GetPhysicalDeviceSurfacePresentModes(device, inst.surface.ref, &modesCount, presentModes),
		fmt.Errorf("failed get physical device present modes"),
	)

	// Deref data
	capabilities.Deref()
	resultFormats := make([]vulkan.SurfaceFormat, 0, len(surfaceFormats))
	for _, format := range surfaceFormats {
		format.Deref()
		resultFormats = append(resultFormats, format)
	}

	// return
	return vkPhysicalDeviceSurface{
		capabilities: capabilities,
		formats:      resultFormats,
		presentModes: presentModes,
	}
}

func (inst *vkInstance) vkPhysicalDeviceQueueFamilies(device vulkan.PhysicalDevice) vkPhysicalDeviceFamily {
	count := uint32(0)
	vulkan.GetPhysicalDeviceQueueFamilyProperties(device, &count, nil)

	if count == 0 {
		return vkPhysicalDeviceFamily{}
	}

	family := make([]vulkan.QueueFamilyProperties, count)
	vulkan.GetPhysicalDeviceQueueFamilyProperties(device, &count, family)

	result := vkPhysicalDeviceFamily{}

	for _, properties := range family {
		properties.Deref()
		if properties.QueueFlags&vulkan.QueueFlags(vulkan.QueueGraphicsBit) != 0 {
			result.hasGraphics = true
		}
		if properties.QueueFlags&vulkan.QueueFlags(vulkan.QueueComputeBit) != 0 {
			result.hasCompute = true
		}
		if properties.QueueFlags&vulkan.QueueFlags(vulkan.QueueTransferBit) != 0 {
			result.hasTransfer = true
		}
	}

	var presentSupport vulkan.Bool32
	vkAssert(
		vulkan.GetPhysicalDeviceSurfaceSupport(device, 0, inst.surface.ref, &presentSupport),
		fmt.Errorf("failed check device surface support"),
	)
	if presentSupport != 0 {
		result.canWindowPresent = true
	}

	return result
}
