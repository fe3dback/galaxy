package render

import (
	"fmt"
	"log"
	"strings"

	"github.com/vulkan-go/vulkan"
)

type (
	vkPhysicalDevice struct {
		ref        vulkan.PhysicalDevice
		props      vulkan.PhysicalDeviceProperties
		features   vulkan.PhysicalDeviceFeatures
		extensions []vulkan.ExtensionProperties
		family     vkPhysicalDeviceFamily
		surface    vkPhysicalDeviceSurface
	}

	vkPhysicalDeviceFamily struct {
		hasGraphics      bool
		hasCompute       bool
		hasTransfer      bool
		canWindowPresent bool
	}

	vkPhysicalDeviceSurface struct {
		capabilities vulkan.SurfaceCapabilities
		formats      []vulkan.SurfaceFormat
		presentModes []vulkan.PresentMode
	}
)

var requiredDeviceExtensions = []string{
	"VK_KHR_swapchain", // require for display buffer to screen
}

func (pd *vkPhysicalDevice) isSupportAllRequiredExtensions() bool {
	supportedExt := make(map[string]struct{})
	for _, extension := range pd.extensions {
		supportedExt[vkLabelToString(extension.ExtensionName)] = struct{}{}
	}

	notSupported := make([]string, 0)
	for _, extension := range requiredDeviceExtensions {
		if _, supported := supportedExt[vkLabelToString(vkStringToLabel(extension))]; supported {
			continue
		}

		notSupported = append(notSupported, extension)
	}

	if len(notSupported) > 0 {
		log.Printf("Device '%s' not support all required extensions: [%s]\n", pd.props.DeviceName, strings.Join(notSupported, ", "))
		return false
	}

	return true
}

func (inst *vkInstance) vkPickPhysicalDevice() *vkPhysicalDevice {
	bestScore := -1
	var bestDevice *vkPhysicalDevice

	for _, device := range inst.vkFindDevices() {
		score := inst.vkPhysicalDeviceScore(device)
		if score < 0 {
			// device is not suitable at all
			continue
		}

		log.Printf("GPU '%s' is suitable for use, score = %d\n", device.props.DeviceName, score)

		// score other found refDevices
		if score > bestScore {
			bestScore = score
			bestDevice = device
		}
	}

	if bestDevice == nil {
		panic(fmt.Errorf("not found suitable vulkan GPU for rendering"))
	}

	log.Printf("Using GPU: %s\n", bestDevice.props.DeviceName)
	return bestDevice
}

func (inst *vkInstance) vkFindDevices() []*vkPhysicalDevice {
	count := uint32(0)
	vkAssert(
		vulkan.EnumeratePhysicalDevices(inst.ref, &count, nil),
		fmt.Errorf("failed EnumeratePhysicalDevices"),
	)
	if count <= 0 {
		panic(fmt.Errorf("not found any GPU with vulkan support"))
	}

	refDevices := make([]vulkan.PhysicalDevice, count)
	vkAssert(
		vulkan.EnumeratePhysicalDevices(inst.ref, &count, refDevices),
		fmt.Errorf("failed EnumeratePhysicalDevices"),
	)

	devices := make([]*vkPhysicalDevice, 0, len(refDevices))
	for _, refDevice := range refDevices {
		devices = append(devices, inst.vkAssemblePhysicalDevice(refDevice))
	}

	return devices
}

func (inst *vkInstance) vkAssemblePhysicalDevice(refDevice vulkan.PhysicalDevice) *vkPhysicalDevice {
	var props vulkan.PhysicalDeviceProperties
	vulkan.GetPhysicalDeviceProperties(refDevice, &props)
	props.Deref()

	var features vulkan.PhysicalDeviceFeatures
	vulkan.GetPhysicalDeviceFeatures(refDevice, &features)
	features.Deref()

	return &vkPhysicalDevice{
		ref:        refDevice,
		props:      props,
		features:   features,
		family:     inst.vkPhysicalDeviceQueueFamilies(refDevice),
		extensions: inst.vkPhysicalDeviceGetExtensions(refDevice),
		surface:    inst.vkPhysicalDeviceGetSurfaceCapabilities(refDevice),
	}
}

// supported values:
// any value < 0 	- device is not suitable for use
// any value >= 0 	- device score
func (inst *vkInstance) vkPhysicalDeviceScore(device *vkPhysicalDevice) int {
	required := map[bool]string{
		// family
		device.family.hasGraphics:      "graphics operations not supported",
		device.family.hasTransfer:      "transfer operations not supported",
		device.family.hasCompute:       "compute operations not supported",
		device.family.canWindowPresent: "window present not supported",

		// extensions
		device.isSupportAllRequiredExtensions(): "not all required extensions supported",

		// swap chain
		len(device.surface.formats) > 0:              "not GPU",
		len(device.surface.presentModes) > 0:         "not GPU",
		device.surface.richColorSpaceFormat() != nil: "rich colorspace not supported",
	}

	// filter
	for passed, reason := range required {
		if !passed {
			log.Printf("GPU '%s' not pass check: %s\n", device.props.DeviceName, reason)
			return -1
		}
	}

	// score
	score := 0
	if device.props.DeviceType == vulkan.PhysicalDeviceTypeDiscreteGpu {
		score += 1000
	}

	return score
}
