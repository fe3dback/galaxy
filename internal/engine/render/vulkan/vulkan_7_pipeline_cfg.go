package vulkan

import (
	"fmt"

	"github.com/vulkan-go/vulkan"

	"github.com/fe3dback/galaxy/internal/utils"
)

type (
	vkPipeLineCfg struct {
		// primitives: topology
		// ---------------------------------
		primitiveTopologyTriangle vulkan.PipelineInputAssemblyStateCreateInfo
		primitiveTopologyLine     vulkan.PipelineInputAssemblyStateCreateInfo

		// rasterizer
		// ---------------------------------
		rasterizerFill vulkan.PipelineRasterizationStateCreateInfo
		rasterizerLine vulkan.PipelineRasterizationStateCreateInfo

		// color blends
		// ---------------------------------
		colorBlendDefault vulkan.PipelineColorBlendStateCreateInfo

		// other
		// ---------------------------------
		multisampling vulkan.PipelineMultisampleStateCreateInfo

		// allocated
		// ---------------------------------
		renderPass vulkan.RenderPass
		layout     vulkan.PipelineLayout
	}
)

func newPipeLineCfg(ld *vkLogicalDevice, swapChain *vkSwapChain, closer *utils.Closer) *vkPipeLineCfg {
	return &vkPipeLineCfg{
		// primitives: topology
		// ---------------------------------
		primitiveTopologyTriangle: createPipeLineAssembleState(vulkan.PrimitiveTopologyTriangleList),
		primitiveTopologyLine:     createPipeLineAssembleState(vulkan.PrimitiveTopologyLineStrip),

		// rasterizer
		// ---------------------------------
		rasterizerFill: createPipeLineRasterizer(vulkan.PolygonModeFill),
		rasterizerLine: createPipeLineRasterizer(vulkan.PolygonModeLine),

		// color blends
		// ---------------------------------
		colorBlendDefault: vulkan.PipelineColorBlendStateCreateInfo{
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
		},

		// other
		// ---------------------------------
		multisampling: vulkan.PipelineMultisampleStateCreateInfo{
			SType:                 vulkan.StructureTypePipelineMultisampleStateCreateInfo,
			RasterizationSamples:  vulkan.SampleCount1Bit,
			SampleShadingEnable:   vulkan.False,
			MinSampleShading:      1.0,
			PSampleMask:           nil,
			AlphaToCoverageEnable: vulkan.False,
			AlphaToOneEnable:      vulkan.False,
		},

		// allocated
		// ---------------------------------
		layout:     createPipeLineLayout(ld, closer),
		renderPass: createPipeLineRenderPass(ld, swapChain, closer),
	}
}

func createPipeLineAssembleState(topology vulkan.PrimitiveTopology) vulkan.PipelineInputAssemblyStateCreateInfo {
	return vulkan.PipelineInputAssemblyStateCreateInfo{
		SType:                  vulkan.StructureTypePipelineInputAssemblyStateCreateInfo,
		Topology:               topology,
		PrimitiveRestartEnable: vulkan.False,
	}
}

func createPipeLineRasterizer(mode vulkan.PolygonMode) vulkan.PipelineRasterizationStateCreateInfo {
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

// PipelineLayout used for input not vertex data into shaders (like textures, uniform buffers, etc...)
func createPipeLineLayout(ld *vkLogicalDevice, closer *utils.Closer) vulkan.PipelineLayout {
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
	closer.EnqueueFree(func() {
		vulkan.DestroyPipelineLayout(ld.ref, pipelineLayout, nil)
	})

	return pipelineLayout
}

// todo: shader input params?
func createPipeLineRenderPass(ld *vkLogicalDevice, swapChain *vkSwapChain, closer *utils.Closer) vulkan.RenderPass {
	colorAttachmentRef := vulkan.AttachmentReference{
		Attachment: 0,
		Layout:     vulkan.ImageLayoutColorAttachmentOptimal,
	}

	subPass := vulkan.SubpassDescription{
		PipelineBindPoint:       vulkan.PipelineBindPointGraphics,
		InputAttachmentCount:    0,
		PInputAttachments:       nil, // todo: shader params?
		ColorAttachmentCount:    1,
		PColorAttachments:       []vulkan.AttachmentReference{colorAttachmentRef},
		PResolveAttachments:     nil,
		PDepthStencilAttachment: nil,
		PreserveAttachmentCount: 0,
		PPreserveAttachments:    nil,
	}

	dependency := vulkan.SubpassDependency{
		SrcSubpass:    vulkan.SubpassExternal,
		DstSubpass:    0,
		SrcStageMask:  vulkan.PipelineStageFlags(vulkan.PipelineStageColorAttachmentOutputBit),
		DstStageMask:  vulkan.PipelineStageFlags(vulkan.PipelineStageColorAttachmentOutputBit),
		SrcAccessMask: 0,
		DstAccessMask: vulkan.AccessFlags(vulkan.AccessColorAttachmentWriteBit),
	}

	colorAttachment := vulkan.AttachmentDescription{
		Format:         swapChain.info.imageFormat,
		Samples:        vulkan.SampleCount1Bit,
		LoadOp:         vulkan.AttachmentLoadOpClear,
		StoreOp:        vulkan.AttachmentStoreOpStore,
		StencilLoadOp:  vulkan.AttachmentLoadOpDontCare,
		StencilStoreOp: vulkan.AttachmentStoreOpDontCare,
		InitialLayout:  vulkan.ImageLayoutUndefined,
		FinalLayout:    vulkan.ImageLayoutPresentSrc,
	}

	renderPassInfo := &vulkan.RenderPassCreateInfo{
		SType:           vulkan.StructureTypeRenderPassCreateInfo,
		AttachmentCount: 1,
		PAttachments:    []vulkan.AttachmentDescription{colorAttachment},
		SubpassCount:    1,
		PSubpasses:      []vulkan.SubpassDescription{subPass},
		DependencyCount: 1,
		PDependencies:   []vulkan.SubpassDependency{dependency},
	}

	var renderPass vulkan.RenderPass
	vkAssert(
		vulkan.CreateRenderPass(ld.ref, renderPassInfo, nil, &renderPass),
		fmt.Errorf("failed create render pass"),
	)
	closer.EnqueueFree(func() {
		vulkan.DestroyRenderPass(ld.ref, renderPass, nil)
	})

	return renderPass
}
