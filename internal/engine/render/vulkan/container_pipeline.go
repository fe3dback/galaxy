package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/engine/render/vulkan/shader/shaderm"
)

var buildInShaders = map[shaderProgram]shaderPipelineFactory{
	&shaderm.Triangle{}: func(c *container, sp shaderProgram) vulkan.Pipeline {
		return c.createShaderPipelineTriangle(sp)
	},
}

func (c *container) createShaderPipelineTriangle(sp shaderProgram) vulkan.Pipeline {
	shaderModuleFrag := c.provideShaderManager().shaderModule(sp.ID() + shaderTypeFrag)
	shaderModuleVert := c.provideShaderManager().shaderModule(sp.ID() + shaderTypeVert)
	shaderStages := []vulkan.PipelineShaderStageCreateInfo{
		shaderModuleFrag.stageInfo,
		shaderModuleVert.stageInfo,
	}

	inputAssemble := c.createPipeLineAssembleStateTopologyTriangle()
	vertexInputInfo := c.createVertexInputInfo(sp)
	viewPortStage := c.createPipelineViewPortState()
	rasterizer := c.createPipeLineRasterizerLine()
	multisampling := c.createPipelineMultisampling()
	colorBlend := c.createPipelineColorBlendDefault()

	pipelineCreateInfo := vulkan.GraphicsPipelineCreateInfo{
		SType:               vulkan.StructureTypeGraphicsPipelineCreateInfo,
		StageCount:          uint32(len(shaderStages)),
		PStages:             shaderStages,
		PVertexInputState:   &vertexInputInfo,
		PInputAssemblyState: &inputAssemble,
		PViewportState:      &viewPortStage,
		PRasterizationState: &rasterizer,
		PMultisampleState:   &multisampling,
		PColorBlendState:    &colorBlend,
		Layout:              c.providePipelineLayout(),
		RenderPass:          c.defaultRenderPass(),
		Subpass:             0,
	}

	pipelines := make([]vulkan.Pipeline, 1)

	// todo: pipeline cache (optimization)
	result := vulkan.CreateGraphicsPipelines(
		c.provideVkLogicalDevice().ref,
		nil,
		1,
		[]vulkan.GraphicsPipelineCreateInfo{pipelineCreateInfo},
		nil,
		pipelines,
	)

	vkAssert(result, fmt.Errorf("failed create graphics pipeline"))
	return pipelines[0]
}

func (c *container) createPipeLineAssembleStateTopologyTriangle() vulkan.PipelineInputAssemblyStateCreateInfo {
	return c.createPipeLineAssembleState(vulkan.PrimitiveTopologyTriangleList)
}

func (c *container) createPipeLineAssembleState(topology vulkan.PrimitiveTopology) vulkan.PipelineInputAssemblyStateCreateInfo {
	return vulkan.PipelineInputAssemblyStateCreateInfo{
		SType:                  vulkan.StructureTypePipelineInputAssemblyStateCreateInfo,
		Topology:               topology,
		PrimitiveRestartEnable: vulkan.False,
	}
}

func (c *container) createVertexInputInfo(sp shaderProgram) vulkan.PipelineVertexInputStateCreateInfo {
	vertexBindings := sp.Bindings()
	vertexAttributes := sp.Attributes()

	return vulkan.PipelineVertexInputStateCreateInfo{
		SType:                           vulkan.StructureTypePipelineVertexInputStateCreateInfo,
		VertexBindingDescriptionCount:   uint32(len(vertexBindings)),
		PVertexBindingDescriptions:      vertexBindings,
		VertexAttributeDescriptionCount: uint32(len(vertexAttributes)),
		PVertexAttributeDescriptions:    vertexAttributes,
	}
}

func (c *container) createPipelineViewPortState() vulkan.PipelineViewportStateCreateInfo {
	swapChain := c.provideSwapChain()
	return vulkan.PipelineViewportStateCreateInfo{
		SType:         vulkan.StructureTypePipelineViewportStateCreateInfo,
		ViewportCount: 1,
		PViewports:    []vulkan.Viewport{swapChain.viewport()},
		ScissorCount:  1,
		PScissors:     []vulkan.Rect2D{swapChain.scissor()},
	}
}

func (c *container) createPipeLineRasterizerFill() vulkan.PipelineRasterizationStateCreateInfo {
	return c.createPipeLineRasterizer(vulkan.PolygonModeFill)
}

func (c *container) createPipeLineRasterizerLine() vulkan.PipelineRasterizationStateCreateInfo {
	return c.createPipeLineRasterizer(vulkan.PolygonModeLine)
}

func (c *container) createPipeLineRasterizer(mode vulkan.PolygonMode) vulkan.PipelineRasterizationStateCreateInfo {
	return vulkan.PipelineRasterizationStateCreateInfo{
		SType:                   vulkan.StructureTypePipelineRasterizationStateCreateInfo,
		DepthClampEnable:        vulkan.False,
		RasterizerDiscardEnable: vulkan.False,
		PolygonMode:             mode,
		CullMode:                vulkan.CullModeFlags(vulkan.CullModeBackBit),
		FrontFace:               vulkan.FrontFaceClockwise,
		DepthBiasEnable:         vulkan.False,
		DepthBiasConstantFactor: 0.0,
		DepthBiasClamp:          0.0,
		DepthBiasSlopeFactor:    0.0,
		LineWidth:               1.0, // todo: require ext
	}
}

func (c *container) createPipelineMultisampling() vulkan.PipelineMultisampleStateCreateInfo {
	return vulkan.PipelineMultisampleStateCreateInfo{
		SType:                 vulkan.StructureTypePipelineMultisampleStateCreateInfo,
		RasterizationSamples:  vulkan.SampleCount1Bit,
		SampleShadingEnable:   vulkan.False,
		MinSampleShading:      1.0,
		PSampleMask:           nil,
		AlphaToCoverageEnable: vulkan.False,
		AlphaToOneEnable:      vulkan.False,
	}
}

func (c *container) createPipelineColorBlendDefault() vulkan.PipelineColorBlendStateCreateInfo {
	return vulkan.PipelineColorBlendStateCreateInfo{
		SType:           vulkan.StructureTypePipelineColorBlendStateCreateInfo,
		LogicOpEnable:   vulkan.False,
		LogicOp:         vulkan.LogicOpCopy,
		AttachmentCount: 1,
		PAttachments: []vulkan.PipelineColorBlendAttachmentState{{
			BlendEnable:         vulkan.True,
			SrcColorBlendFactor: vulkan.BlendFactorSrcAlpha,
			DstColorBlendFactor: vulkan.BlendFactorOneMinusSrcAlpha,
			ColorBlendOp:        vulkan.BlendOpAdd,
			SrcAlphaBlendFactor: vulkan.BlendFactorOne,
			DstAlphaBlendFactor: vulkan.BlendFactorZero,
			AlphaBlendOp:        vulkan.BlendOpAdd,
			ColorWriteMask: vulkan.ColorComponentFlags(
				vulkan.ColorComponentRBit | vulkan.ColorComponentGBit | vulkan.ColorComponentBBit | vulkan.ColorComponentABit,
			),
		}},
		BlendConstants: [4]float32{0, 0, 0, 0},
	}
}

// PipelineLayout used for input not vertex data into shaders (like textures, uniform buffers, etc...)
func newPipeLineLayout(ld *vkLogicalDevice) vulkan.PipelineLayout {
	info := &vulkan.PipelineLayoutCreateInfo{
		SType:                  vulkan.StructureTypePipelineLayoutCreateInfo,
		SetLayoutCount:         0,
		PSetLayouts:            nil,
		PushConstantRangeCount: 0,
		PPushConstantRanges:    nil,
	}

	var pipelineLayout vulkan.PipelineLayout
	vkAssert(
		vulkan.CreatePipelineLayout(ld.ref, info, nil, &pipelineLayout),
		fmt.Errorf("failed create pipeline layout"),
	)

	return pipelineLayout
}
