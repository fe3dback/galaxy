package vulkan_depr

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkLogicalDevice struct {
		ref           vulkan.Device
		queueGraphics vulkan.Queue
		queuePresent  vulkan.Queue
	}
)

func (pd *vkPhysicalDevice) createLogicalDevice(closer *utils.Closer) *vkLogicalDevice {
	queueCreateInfos := make([]vulkan.DeviceQueueCreateInfo, 0)
	for _, familyId := range pd.families.uniqueIDs() {
		queueCreateInfos = append(queueCreateInfos, vulkan.DeviceQueueCreateInfo{
			SType:            vulkan.StructureTypeDeviceQueueCreateInfo,
			QueueFamilyIndex: familyId,
			QueueCount:       1,
			PQueuePriorities: []float32{1.0},
		})
	}

	deviceCreateInfo := &vulkan.DeviceCreateInfo{
		SType:                   vulkan.StructureTypeDeviceCreateInfo,
		QueueCreateInfoCount:    uint32(len(queueCreateInfos)),
		PQueueCreateInfos:       queueCreateInfos,
		PEnabledFeatures:        []vulkan.PhysicalDeviceFeatures{pd.features},
		EnabledExtensionCount:   uint32(len(requiredDeviceExtensions)),
		PpEnabledExtensionNames: vkStringsToStringLabels(requiredDeviceExtensions),
	}

	var logicalDevice vulkan.Device
	vkAssert(
		vulkan.CreateDevice(pd.ref, deviceCreateInfo, nil, &logicalDevice),
		fmt.Errorf("failed create logical device"),
	)
	closer.EnqueueFree(func() {
		vulkan.DestroyDevice(logicalDevice, nil)
	})

	var queueGraphics vulkan.Queue
	var queuePresent vulkan.Queue
	vulkan.GetDeviceQueue(logicalDevice, pd.families.graphicsFamilyId, 0, &queueGraphics)
	vulkan.GetDeviceQueue(logicalDevice, pd.families.presentFamilyId, 0, &queuePresent)

	log.Printf("vk: logical device created (graphicsQ: %d, presentQ: %d)\n",
		pd.families.graphicsFamilyId,
		pd.families.presentFamilyId,
	)

	return &vkLogicalDevice{
		ref:           logicalDevice,
		queueGraphics: queueGraphics,
		queuePresent:  queuePresent,
	}
}
