package render

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
	const queueFamilyGraphics = 0
	const queueFamilyPresent = 1

	queueCreateInfo := []vulkan.DeviceQueueCreateInfo{
		{
			SType:            vulkan.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: queueFamilyGraphics,
			QueueCount:       1,
			PQueuePriorities: []float32{1.0},
		},
		{
			SType:            vulkan.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: queueFamilyPresent,
			QueueCount:       1,
			PQueuePriorities: []float32{1.0},
		},
	}

	createInfo := &vulkan.DeviceCreateInfo{
		SType:                vulkan.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount: uint32(len(queueCreateInfo)),
		PQueueCreateInfos:    queueCreateInfo,
		PEnabledFeatures:     []vulkan.PhysicalDeviceFeatures{pd.features},
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
	vulkan.GetDeviceQueue(logicalDevice, queueFamilyGraphics, 0, &graphicsQueue)
	vulkan.GetDeviceQueue(logicalDevice, queueFamilyPresent, 0, &presentQueue)

	return &vkLogicalDevice{
		ref:           logicalDevice,
		graphicsQueue: graphicsQueue,
		presetQueue:   presentQueue,
	}
}
