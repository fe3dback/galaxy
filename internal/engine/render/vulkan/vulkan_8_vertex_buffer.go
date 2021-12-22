package vulkan

import (
	"fmt"
	"unsafe"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vertexData interface {
		Size() uint64
		Data() []byte
		Bindings() []vulkan.VertexInputBindingDescription
		Attributes() []vulkan.VertexInputAttributeDescription
	}
)

func createVertexBuffer(pd *vkPhysicalDevice, ld *vkLogicalDevice, vertexData vertexData, closer *utils.Closer) vulkan.Buffer {
	info := &vulkan.BufferCreateInfo{
		SType:       vulkan.StructureTypeBufferCreateInfo,
		Size:        vulkan.DeviceSize(vertexData.Size()),
		Usage:       vulkan.BufferUsageFlags(vulkan.BufferUsageVertexBufferBit),
		SharingMode: vulkan.SharingModeExclusive,
	}

	var buffer vulkan.Buffer
	vkAssert(
		vulkan.CreateBuffer(ld.ref, info, nil, &buffer),
		fmt.Errorf("failed create vertex buffer"),
	)

	var memoryReq vulkan.MemoryRequirements
	vulkan.GetBufferMemoryRequirements(ld.ref, buffer, &memoryReq)
	memoryReq.Deref()

	memAllocInfo := &vulkan.MemoryAllocateInfo{
		SType:          vulkan.StructureTypeMemoryAllocateInfo,
		AllocationSize: memoryReq.Size,
		MemoryTypeIndex: findVertexBufferMemoryType(
			pd,
			memoryReq.MemoryTypeBits,
			vulkan.MemoryPropertyFlags(vulkan.MemoryPropertyHostVisibleBit|vulkan.MemoryPropertyHostCoherentBit),
		),
	}

	var bufferMemory vulkan.DeviceMemory
	vkAssert(
		vulkan.AllocateMemory(ld.ref, memAllocInfo, nil, &bufferMemory),
		fmt.Errorf("failed allocate GPU memory for vertex buffer"),
	)

	vulkan.BindBufferMemory(ld.ref, buffer, bufferMemory, 0)

	closer.EnqueueFree(func() {
		vulkan.DestroyBuffer(ld.ref, buffer, nil)
		vulkan.FreeMemory(ld.ref, bufferMemory, nil)
	})

	var data unsafe.Pointer
	vulkan.MapMemory(ld.ref, bufferMemory, 0, info.Size, 0, &data)
	vulkan.Memcopy(data, vertexData.Data())
	vulkan.UnmapMemory(ld.ref, bufferMemory)

	return buffer
}

func findVertexBufferMemoryType(pd *vkPhysicalDevice, typeFilter uint32, memFlags vulkan.MemoryPropertyFlags) uint32 {
	var memProperties vulkan.PhysicalDeviceMemoryProperties
	vulkan.GetPhysicalDeviceMemoryProperties(pd.ref, &memProperties)
	memProperties.Deref()

	for i := uint32(0); i < memProperties.MemoryTypeCount; i++ {
		memType := memProperties.MemoryTypes[i]
		memType.Deref()

		if (typeFilter&(1<<i) != 0) && ((memType.PropertyFlags & memFlags) == memFlags) {
			return i
		}
	}

	panic(fmt.Errorf("failed find suitable GPU memory for vertex buffer"))
}
