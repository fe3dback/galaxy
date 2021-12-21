package vulkan

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

const shaderEntryPoint = "main"

type (
	shaderModuleID  = string
	vkShaderManager struct {
		ld      *vkLogicalDevice
		closer  *utils.Closer
		modules map[shaderModuleID]*vkShaderModule
	}

	vkShaderModule struct {
		module    vulkan.ShaderModule
		stageInfo vulkan.PipelineShaderStageCreateInfo
	}
)

func newShaderManager(ld *vkLogicalDevice, closer *utils.Closer) *vkShaderManager {
	return &vkShaderManager{
		ld:      ld,
		closer:  closer,
		modules: make(map[shaderModuleID]*vkShaderModule),
	}
}

func (sm *vkShaderManager) shaderModule(id shaderModuleID) *vkShaderModule {
	if module, exist := sm.modules[id]; exist {
		return module
	}

	sm.modules[id] = sm.allocModule(id)
	return sm.modules[id]
}

func (sm *vkShaderManager) allocModule(id shaderModuleID) *vkShaderModule {
	byteCode, ok := shadersByteCode[id]
	if !ok {
		panic(fmt.Errorf("failed find compiled bytecode for shader '%s'", id))
	}

	stageType, ok := shadersType[id]
	if !ok {
		panic(fmt.Errorf("failed find shader type for shader '%s'", id))
	}

	createInfo := &vulkan.ShaderModuleCreateInfo{
		SType:    vulkan.StructureTypeShaderModuleCreateInfo,
		CodeSize: uint(len(byteCode)),
		PCode:    vkTransformBytes(byteCode),
	}

	var shaderModule vulkan.ShaderModule
	vkAssert(
		vulkan.CreateShaderModule(sm.ld.ref, createInfo, nil, &shaderModule),
		fmt.Errorf("failed create shader module from '%s' shader", id),
	)

	sm.closer.EnqueueFree(func() {
		vulkan.DestroyShaderModule(sm.ld.ref, shaderModule, nil)
	})

	log.Printf("vk: shader allocated '%s', len=%d\n", id, len(byteCode))
	return &vkShaderModule{
		module: shaderModule,
		stageInfo: vulkan.PipelineShaderStageCreateInfo{
			SType:  vulkan.StructureTypePipelineShaderStageCreateInfo,
			Stage:  stageType,
			Module: shaderModule,
			PName:  fmt.Sprintf("%s\x00", shaderEntryPoint),
		},
	}
}
