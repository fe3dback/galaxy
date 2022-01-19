package vulkan

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/vulkan-go/vulkan"
)

type (
	dataInstance interface {
		Data() []byte
	}

	bindStats struct {
		instanceCount uint32
		buffers       []vulkan.Buffer
		offsets       []vulkan.DeviceSize
	}
)

func newDataBuffersManager(ld *vkLogicalDevice, pd *vkPhysicalDevice) *vkDataBuffersManager {
	return &vkDataBuffersManager{
		residentVertex: vkBufferTable{
			totalCapacity: 0,
			buffers:       []vkBuffer{},
		},
		ld: ld,
		pd: pd,
	}
}

func (vb *vkDataBuffersManager) free() {
	buffers := make([]vkBuffer, 0)
	buffers = append(buffers, vb.residentVertex.buffers...)

	for _, vkBuffer := range buffers {
		vulkan.DestroyBuffer(vb.ld.ref, vkBuffer.handle, nil)
		vulkan.FreeMemory(vb.ld.ref, vkBuffer.memory, nil)
	}

	log.Printf("vk: freed: vertex buffers\n")
}

func (vb *vkDataBuffersManager) resetVertexBuffers() {
	vb.residentVertex.framePageID = -1
	vb.residentVertex.framePageCapacity = 0
	vb.residentVertex.frameInstanceCounts = []uint32{}
	vb.residentVertex.frameStagedData = [][]byte{}

	for range vb.residentVertex.buffers {
		vb.residentVertex.frameStagedData = append(vb.residentVertex.frameStagedData, []byte{})
		vb.residentVertex.frameInstanceCounts = append(vb.residentVertex.frameInstanceCounts, 0)
	}
}

func (vb *vkDataBuffersManager) flushVertexBuffers() []bindStats {
	stats := make([]bindStats, 0)

	for pageID, stagedData := range vb.residentVertex.frameStagedData {
		buff := vb.residentVertex.buffers[pageID]
		vulkan.Memcopy(buff.dataPtr, stagedData)

		stats = append(stats, bindStats{
			instanceCount: vb.residentVertex.frameInstanceCounts[pageID],
			buffers:       []vulkan.Buffer{buff.handle},
			offsets:       []vulkan.DeviceSize{0},
		})
	}

	return stats
}

func (vb *vkDataBuffersManager) writeToVertexBuffers(instance dataInstance) {
	data := instance.Data()
	size := uint64(len(data))

	if vb.residentVertex.framePageCapacity < size {
		vb.residentVertex.framePageID++

		if int16(len(vb.residentVertex.buffers)-1) < vb.residentVertex.framePageID {
			vb.allocateNewVertexBuffer()
		}

		buff := vb.residentVertex.buffers[vb.residentVertex.framePageID]
		vb.residentVertex.framePageCapacity = uint64(buff.capacity)
	}

	// write to buffer
	vb.residentVertex.frameInstanceCounts[vb.residentVertex.framePageID]++
	stageData := &(vb.residentVertex.frameStagedData[vb.residentVertex.framePageID])
	*stageData = append(*stageData, data...)
	vb.residentVertex.framePageCapacity -= size
}

func (vb *vkDataBuffersManager) allocateNewVertexBuffer() {
	buff := allocatePersistBuffer(vb.ld, vb.pd)
	vb.residentVertex.buffers = append(vb.residentVertex.buffers, buff)
	vb.residentVertex.frameStagedData = append(vb.residentVertex.frameStagedData, []byte{})
	vb.residentVertex.frameInstanceCounts = append(vb.residentVertex.frameInstanceCounts, 0)
	vb.residentVertex.totalCapacity += uint64(buff.capacity)
}

func allocatePersistBuffer(ld *vkLogicalDevice, pd *vkPhysicalDevice) vkBuffer {
	// create new buffer page
	info := &vulkan.BufferCreateInfo{
		SType:       vulkan.StructureTypeBufferCreateInfo,
		Size:        vulkan.DeviceSize(vertexBufferSize),
		Usage:       vulkan.BufferUsageFlags(vulkan.BufferUsageVertexBufferBit),
		SharingMode: vulkan.SharingModeExclusive,
	}

	var buffer vulkan.Buffer
	vkAssert(
		vulkan.CreateBuffer(ld.ref, info, nil, &buffer),
		fmt.Errorf("failed create vertex buffer"),
	)

	// get device memory requirements for it
	var memoryReq vulkan.MemoryRequirements
	vulkan.GetBufferMemoryRequirements(ld.ref, buffer, &memoryReq)
	memoryReq.Deref()

	memoryTypeIndex := findVertexBufferMemoryType(
		pd,
		memoryReq,
		vulkan.MemoryPropertyFlags(
			vulkan.MemoryPropertyHostVisibleBit|
				vulkan.MemoryPropertyHostCoherentBit,
		),
	)

	memAllocInfo := &vulkan.MemoryAllocateInfo{
		SType:           vulkan.StructureTypeMemoryAllocateInfo,
		AllocationSize:  memoryReq.Size,
		MemoryTypeIndex: memoryTypeIndex,
	}

	var bufferMemory vulkan.DeviceMemory
	vkAssert(
		vulkan.AllocateMemory(ld.ref, memAllocInfo, nil, &bufferMemory),
		fmt.Errorf("failed allocate GPU memory for vertex buffer"),
	)

	vulkan.BindBufferMemory(ld.ref, buffer, bufferMemory, 0)

	var data unsafe.Pointer
	vulkan.MapMemory(ld.ref, bufferMemory, 0, info.Size, 0, &data)

	log.Printf("Buffer %dMB capacity - allocated", info.Size/1024)

	return vkBuffer{
		capacity: info.Size,
		dataPtr:  data,
		handle:   buffer,
		memory:   bufferMemory,
	}
}

func findVertexBufferMemoryType(pd *vkPhysicalDevice, memoryReq vulkan.MemoryRequirements, memFlags vulkan.MemoryPropertyFlags) uint32 {
	typeFilter := memoryReq.MemoryTypeBits

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
