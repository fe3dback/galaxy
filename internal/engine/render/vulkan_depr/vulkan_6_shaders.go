package vulkan_depr

import (
	"fmt"
	"log"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

const shaderEntryPoint = "main"
const shaderTypeVert = ".vert"
const shaderTypeFrag = ".frag"

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

func (sm *vkShaderManager) preloadShaders() {
	for _, shader := range shaderModules {
		sm.modules[shader.ID()+shaderTypeFrag] = sm.allocModule(shader.ID()+shaderTypeFrag, shader.ProgramFrag(), vulkan.ShaderStageFragmentBit)
		sm.modules[shader.ID()+shaderTypeVert] = sm.allocModule(shader.ID()+shaderTypeVert, shader.ProgramVert(), vulkan.ShaderStageVertexBit)
	}
}

func (sm *vkShaderManager) shaderPipeline(id shaderModuleID) vulkan.GraphicsPipelineCreateInfo {
	shaderStages := []vulkan.PipelineShaderStageCreateInfo{
		sm.shaderModule(id + shaderTypeVert).stageInfo,
		sm.shaderModule(id + shaderTypeFrag).stageInfo,
	}

	return vulkan.GraphicsPipelineCreateInfo{
		SType:      vulkan.StructureTypeGraphicsPipelineCreateInfo,
		StageCount: uint32(len(shaderStages)),
		PStages:    shaderStages,
	}
}

func (sm *vkShaderManager) shaderModule(id shaderModuleID) *vkShaderModule {
	if module, exist := sm.modules[id]; exist {
		return module
	}

	return sm.modules[id]
}

func (sm *vkShaderManager) allocModule(id shaderModuleID, byteCode []byte, stageType vulkan.ShaderStageFlagBits) *vkShaderModule {
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
