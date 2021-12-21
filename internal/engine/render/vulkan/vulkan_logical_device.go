package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkLogicalDevice struct {
		ref           vulkan.Device
		graphicsQueue vulkan.Queue
		presetQueue   vulkan.Queue
	}
)

func (pd *vkPhysicalDevice) createLogicalDevice(opts vkCreateOptions) *vkLogicalDevice {
	queueCreateInfo := make(map[uint32]vulkan.DeviceQueueCreateInfo)
	uniqueFamilies := make(map[uint32]struct{})
	uniqueFamilies[pd.family.graphicsFamilyId] = struct{}{}
	uniqueFamilies[pd.family.presentFamilyId] = struct{}{}

	for uniqueFamilyId := range uniqueFamilies {
		queueCreateInfo[uniqueFamilyId] = vulkan.DeviceQueueCreateInfo{
			SType:            vulkan.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: pd.family.graphicsFamilyId,
			QueueCount:       1,
			PQueuePriorities: []float32{1.0},
		}
	}

	infos := make([]vulkan.DeviceQueueCreateInfo, 0, len(queueCreateInfo))
	for _, info := range queueCreateInfo {
		infos = append(infos, info)
	}

	createInfo := &vulkan.DeviceCreateInfo{
		SType:                   vulkan.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount:    uint32(len(infos)),
		PQueueCreateInfos:       infos,
		PEnabledFeatures:        []vulkan.PhysicalDeviceFeatures{pd.features},
		EnabledExtensionCount:   uint32(len(requiredDeviceExtensions)),
		PpEnabledExtensionNames: vkStringsToStringLabels(requiredDeviceExtensions),
	}

	var logicalDevice vulkan.Device
	vkAssert(
		vulkan.CreateDevice(pd.ref, createInfo, nil, &logicalDevice),
		fmt.Errorf("failed create logical device"),
	)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroyDevice(logicalDevice, nil)
	})

	var graphicsQueue vulkan.Queue
	var presentQueue vulkan.Queue
	vulkan.GetDeviceQueue(logicalDevice, pd.family.graphicsFamilyId, 0, &graphicsQueue)
	vulkan.GetDeviceQueue(logicalDevice, pd.family.presentFamilyId, 0, &presentQueue)

	return &vkLogicalDevice{
		ref:           logicalDevice,
		graphicsQueue: graphicsQueue,
		presetQueue:   presentQueue,
	}
}
