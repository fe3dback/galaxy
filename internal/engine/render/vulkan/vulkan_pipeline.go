package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"
)

type (
	vkPipeline struct {
		ref vulkan.Pipeline
	}
)

func (vk *Vk) vkCreatePipeline(opts vkCreateOptions) *vkPipeline {
	triangleVert := vk.logicalDevice.vkCreateShaderModule(opts, "triangle.vert", vulkan.ShaderStageVertexBit, commonShaderTriangleVert)
	triangleFrag := vk.logicalDevice.vkCreateShaderModule(opts, "triangle.frag", vulkan.ShaderStageFragmentBit, commonShaderTriangleFrag)

	shaderStages := []vulkan.PipelineShaderStageCreateInfo{triangleVert.stageInfo, triangleFrag.stageInfo}

	vertexInputInfo := &vulkan.PipelineVertexInputStateCreateInfo{
		SType:                           vulkan.StructureTypePipelineVertexInputStateCreateInfo,
		VertexBindingDescriptionCount:   0,
		PVertexBindingDescriptions:      nil,
		VertexAttributeDescriptionCount: 0,
		PVertexAttributeDescriptions:    nil,
	}

	inputAssemble := &vulkan.PipelineInputAssemblyStateCreateInfo{
		SType:                  vulkan.StructureTypePipelineInputAssemblyStateCreateInfo,
		Topology:               vulkan.PrimitiveTopologyTriangleList,
		PrimitiveRestartEnable: vulkan.False,
	}

	viewport := vulkan.Viewport{
		X:        0,
		Y:        0,
		Width:    float32(vk.swapChain.info.bufferSize.Width),
		Height:   float32(vk.swapChain.info.bufferSize.Height),
		MinDepth: 0.0,
		MaxDepth: 1.0,
	}

	scissor := vulkan.Rect2D{
		Offset: vulkan.Offset2D{
			X: 0,
			Y: 0,
		},
		Extent: vk.swapChain.info.bufferSize,
	}

	viewPortStage := &vulkan.PipelineViewportStateCreateInfo{
		SType:         vulkan.StructureTypePipelineViewportStateCreateInfo,
		ViewportCount: 1,
		PViewports:    []vulkan.Viewport{viewport},
		ScissorCount:  1,
		PScissors:     []vulkan.Rect2D{scissor},
	}

	rasterizer := &vulkan.PipelineRasterizationStateCreateInfo{
		SType:                   vulkan.StructureTypePipelineRasterizationStateCreateInfo,
		DepthClampEnable:        vulkan.False,
		RasterizerDiscardEnable: vulkan.False,
		PolygonMode:             vulkan.PolygonModeFill,
		CullMode:                vulkan.CullModeFlags(vulkan.CullModeBackBit),
		FrontFace:               vulkan.FrontFaceClockwise,
		DepthBiasEnable:         vulkan.False,
		DepthBiasConstantFactor: 0.0,
		DepthBiasClamp:          0.0,
		DepthBiasSlopeFactor:    0.0,
		LineWidth:               1.0,
	}

	multisampling := &vulkan.PipelineMultisampleStateCreateInfo{
		SType:                 vulkan.StructureTypePipelineMultisampleStateCreateInfo,
		RasterizationSamples:  vulkan.SampleCount1Bit,
		SampleShadingEnable:   vulkan.False,
		MinSampleShading:      1.0,
		PSampleMask:           nil,
		AlphaToCoverageEnable: vulkan.False,
		AlphaToOneEnable:      vulkan.False,
	}

	// finalColor.rgb = newAlpha * newColor + (1 - newAlpha) * oldColor;
	// finalColor.a = newAlpha.a;

	colorBlendAttachment := vulkan.PipelineColorBlendAttachmentState{
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
	}

	colorBlending := &vulkan.PipelineColorBlendStateCreateInfo{
		SType:           vulkan.StructureTypePipelineColorBlendStateCreateInfo,
		LogicOpEnable:   vulkan.False,
		LogicOp:         vulkan.LogicOpCopy,
		AttachmentCount: 1,
		PAttachments:    []vulkan.PipelineColorBlendAttachmentState{colorBlendAttachment},
		BlendConstants:  [4]float32{0, 0, 0, 0},
	}

	pipelineLayoutInfo := &vulkan.PipelineLayoutCreateInfo{
		SType:                  vulkan.StructureTypePipelineLayoutCreateInfo,
		SetLayoutCount:         0,
		PSetLayouts:            nil,
		PushConstantRangeCount: 0,
		PPushConstantRanges:    nil,
	}

	var pipelineLayout vulkan.PipelineLayout
	vkAssert(
		vulkan.CreatePipelineLayout(vk.logicalDevice.ref, pipelineLayoutInfo, nil, &pipelineLayout),
		fmt.Errorf("failed create pipeline layout"),
	)
	opts.closer.EnqueueFree(func() {
		vulkan.DestroyPipelineLayout(vk.logicalDevice.ref, pipelineLayout, nil)
	})

	pipelineCreateInfo := vulkan.GraphicsPipelineCreateInfo{
		SType:               vulkan.StructureTypeGraphicsPipelineCreateInfo,
		StageCount:          uint32(len(shaderStages)),
		PStages:             shaderStages,
		PVertexInputState:   vertexInputInfo,
		PInputAssemblyState: inputAssemble,
		PViewportState:      viewPortStage,
		PRasterizationState: rasterizer,
		PMultisampleState:   multisampling,
		PColorBlendState:    colorBlending,
		Layout:              pipelineLayout,
		RenderPass:          vk.renderPass,
		Subpass:             0,
	}

	pipelines := make([]vulkan.Pipeline, 1)
	result := vulkan.CreateGraphicsPipelines(
		vk.logicalDevice.ref,
		nil,
		1,
		[]vulkan.GraphicsPipelineCreateInfo{pipelineCreateInfo},
		nil,
		pipelines,
	)
	pipeline := pipelines[0]

	vkAssert(result, fmt.Errorf("failed create graphics pipeline"))
	opts.closer.EnqueueFree(func() {
		vulkan.DestroyPipeline(vk.logicalDevice.ref, pipeline, nil)
	})

	return &vkPipeline{
		ref: pipeline,
	}
}
