package animator

import (
	"github.com/fe3dback/galaxy/render"
)

func (anim *Animator) OnDraw(r *render.Renderer) error {
	if !anim.initialized {
		anim.initialize(r)
	}

	seq := anim.activeSequence
	res := seq.resource
	frame := seq.frames[seq.currentFrame]
	entityPos := anim.entity.Position()

	dest := &render.Rect{
		X: int32(entityPos.X) + int32(seq.offsetX),
		Y: int32(entityPos.Y) + int32(seq.offsetY),
		W: int32(frame.w),
		H: int32(frame.h),
	}

	r.DrawSpriteEx(res, frame.TextureRect(), dest)

	return nil
}

func (anim *Animator) OnUpdate(_ float64) error {
	if anim.paused {
		return nil
	}

	seq := anim.activeSequence

	seq.currentFrame++

	// todo: fps
	// todo: bounce
	// todo: repeat
	// todo: playOnce

	if seq.currentFrame >= seq.lastFrame {
		seq.currentFrame = 0
	}

	return nil
}
