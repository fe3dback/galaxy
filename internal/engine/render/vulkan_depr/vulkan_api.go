package vulkan_depr

import "github.com/fe3dback/galaxy/internal/engine/render/vulkan_depr/shader/shaderm"

func (vk *Vk) FrameBegin() {

}

func (vk *Vk) FrameEnd() {

}

func (vk *Vk) DrawTriangle(triangle *shaderm.Triangle) {
	vk.frameQueue = append(vk.frameQueue, triangle)
}
