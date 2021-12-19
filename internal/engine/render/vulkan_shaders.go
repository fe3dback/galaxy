package render

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkShaderModule struct {
		module    vulkan.ShaderModule
		stageInfo vulkan.PipelineShaderStageCreateInfo
	}
)

func (ld *vkLogicalDevice) vkCreateShaderModule(opts vkCreateOptions, name string, stage vulkan.ShaderStageFlagBits, byteCode []byte) *vkShaderModule {
	createInfo := &vulkan.ShaderModuleCreateInfo{
		SType:    vulkan.StructureTypeShaderModuleCreateInfo,
		CodeSize: uint(len(byteCode)),
		PCode:    vkTransformBytes(byteCode),
	}

	var shaderModule vulkan.ShaderModule
	vkAssert(
		vulkan.CreateShaderModule(ld.ref, createInfo, nil, &shaderModule),
		fmt.Errorf("failed create shader module from '%s' shader", name),
	)

	opts.closer.EnqueueFree(func() {
		vulkan.DestroyShaderModule(ld.ref, shaderModule, nil)
	})

	return &vkShaderModule{
		module:    shaderModule,
		stageInfo: ld.vkCreateShaderStageInfo(shaderModule, stage),
	}
}

func (ld *vkLogicalDevice) vkCreateShaderStageInfo(module vulkan.ShaderModule, stage vulkan.ShaderStageFlagBits) vulkan.PipelineShaderStageCreateInfo {
	return vulkan.PipelineShaderStageCreateInfo{
		SType:  vulkan.StructureTypePipelineShaderStageCreateInfo,
		Stage:  stage,
		Module: module,
		PName:  "main\x00",
	}
}
