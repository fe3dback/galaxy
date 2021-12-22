package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkPipeline struct {
		ref vulkan.Pipeline
		cfg *vkPipeLineCfg

		_free   bool
		_freeLd *vkLogicalDevice
	}
)

func (pl *vkPipeline) free() {
	if pl._free {
		return
	}

	pl._free = true
	vulkan.DestroyPipeline(pl._freeLd.ref, pl.ref, nil)
	pl.cfg.free()
}

func createPipeline(
	cfg *vkPipeLineCfg,
	ld *vkLogicalDevice,
	swapChain *vkSwapChain,
	shaderStages []vulkan.PipelineShaderStageCreateInfo,
	closer *utils.Closer,
) *vkPipeline {
	pl := &vkPipeline{
		cfg:     cfg,
		_freeLd: ld,
	}
	closer.EnqueueFree(pl.free)

	// data (input)
	inputAssemble := cfg.primitiveTopologyTriangle
	vertexInputInfo := &vulkan.PipelineVertexInputStateCreateInfo{ // todo: shader arguments??
		SType:                           vulkan.StructureTypePipelineVertexInputStateCreateInfo,
		VertexBindingDescriptionCount:   0,
		PVertexBindingDescriptions:      nil,
		VertexAttributeDescriptionCount: 0,
		PVertexAttributeDescriptions:    nil,
	}

	// viewport
	viewport := swapChain.viewport()
	scissor := swapChain.scissor()
	viewPortStage := &vulkan.PipelineViewportStateCreateInfo{
		SType:         vulkan.StructureTypePipelineViewportStateCreateInfo,
		ViewportCount: 1,
		PViewports:    []vulkan.Viewport{viewport},
		ScissorCount:  1,
		PScissors:     []vulkan.Rect2D{scissor},
	}

	pipelineCreateInfo := vulkan.GraphicsPipelineCreateInfo{
		SType:               vulkan.StructureTypeGraphicsPipelineCreateInfo,
		StageCount:          uint32(len(shaderStages)),
		PStages:             shaderStages,
		PVertexInputState:   vertexInputInfo,
		PInputAssemblyState: &inputAssemble,
		PViewportState:      viewPortStage,
		PRasterizationState: &cfg.rasterizerFill,
		PMultisampleState:   &cfg.multisampling,
		PColorBlendState:    &cfg.colorBlendDefault,
		Layout:              cfg.layout,
		RenderPass:          cfg.renderPass,
		Subpass:             0,
	}

	pipelines := make([]vulkan.Pipeline, 1)
	// todo: pipeline cache (optimization)
	result := vulkan.CreateGraphicsPipelines(
		ld.ref,
		nil,
		1,
		[]vulkan.GraphicsPipelineCreateInfo{pipelineCreateInfo},
		nil,
		pipelines,
	)

	vkAssert(result, fmt.Errorf("failed create graphics pipeline"))
	pl.ref = pipelines[0]
	return pl
}
