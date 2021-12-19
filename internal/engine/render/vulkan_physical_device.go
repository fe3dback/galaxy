package render

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"
)

type (
	vkPhysicalDevice struct {
		ref      vulkan.PhysicalDevice
		props    vulkan.PhysicalDeviceProperties
		features vulkan.PhysicalDeviceFeatures
		family   vkPhysicalDeviceFamily
	}

	vkPhysicalDeviceFamily struct {
		hasGraphics      bool
		hasCompute       bool
		hasTransfer      bool
		canWindowPresent bool
	}
)

func (inst *vkInstance) vkFindDevices() []*vkPhysicalDevice {
	count := uint32(0)
	vulkan.EnumeratePhysicalDevices(inst.ref, &count, nil)
	if count <= 0 {
		panic(fmt.Errorf("not found any GPU with vulkan support"))
	}

	refDevices := make([]vulkan.PhysicalDevice, count)
	vulkan.EnumeratePhysicalDevices(inst.ref, &count, refDevices)

	devices := make([]*vkPhysicalDevice, 0, len(refDevices))
	for _, refDevice := range refDevices {
		var props vulkan.PhysicalDeviceProperties
		vulkan.GetPhysicalDeviceProperties(refDevice, &props)
		props.Deref()

		var features vulkan.PhysicalDeviceFeatures
		vulkan.GetPhysicalDeviceFeatures(refDevice, &features)
		features.Deref()

		family := inst.vkPhysicalDeviceQueueFamilies(refDevice)

		devices = append(devices, &vkPhysicalDevice{
			ref:      refDevice,
			props:    props,
			features: features,
			family:   family,
		})
	}

	return devices
}

func (inst *vkInstance) vkPickPhysicalDevice() *vkPhysicalDevice {
	bestScore := -1
	var bestDevice *vkPhysicalDevice

	for _, device := range inst.vkFindDevices() {
		score := inst.vkPhysicalDeviceScore(device)
		if score < 0 {
			// refDevice is not suitable at all
			log.Printf("GPU '%s' is not suitable for use, ignoring\n", device.props.DeviceName)
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

// supported values:
// any value < 0 	- device is not suitable for use
// any value >= 0 	- device score
func (inst *vkInstance) vkPhysicalDeviceScore(device *vkPhysicalDevice) int {
	// filter
	if !device.family.hasGraphics {
		return -1
	}
	if !device.family.hasTransfer {
		return -1
	}
	if !device.family.hasCompute {
		return -1
	}
	if !device.family.canWindowPresent {
		return -1
	}

	// score
	score := 0
	if device.props.DeviceType == vulkan.PhysicalDeviceTypeDiscreteGpu {
		score += 1000
	}

	return score
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
	vulkan.GetPhysicalDeviceSurfaceSupport(device, 0, inst.surface.ref, &presentSupport)
	if presentSupport != 0 {
		result.canWindowPresent = true
	}

	return result
}
