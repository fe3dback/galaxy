package animator

import (
	"math"

	"github.com/fe3dback/galaxy/engine"
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
		X: int32(entityPos.X) + int32(seq.offsetX) - (int32(frame.w) / 2),
		Y: int32(entityPos.Y) + int32(seq.offsetY) - (int32(frame.h) / 2),
		W: int32(frame.w),
		H: int32(frame.h),
	}

	r.DrawSpriteEx(res, frame.TextureRect(), dest, anim.entity.Rotation().ToFloat())
	r.DrawSquare(
		engine.ColorGreen,
		int(dest.X),
		int(dest.Y),
		int(dest.W),
		int(dest.H),
	)

	return nil
}

func (anim *Animator) OnUpdate(moment engine.Moment) error {
	if !anim.initialized || anim.paused {
		return nil
	}

	seq := anim.activeSequence

	// if one time sequence
	if seq.finished {
		return nil
	}

	// move frames
	deltaProgress := moment.DeltaTime() * float64(moment.TargetFPS())
	seq.progress += deltaProgress * seq.progressMod
	if seq.progress > float64(moment.TargetFPS()) {
		seq.progress = 0

		// end of sequence
		if !seq.repeat {
			seq.finished = true
			seq.currentFrame = seq.lastFrame
			return nil
		}

		// bounce
		if seq.bounce {
			anim.switchSequenceDirection(seq)
		}
	}

	// calculate current frame depend on fps and game time
	frameIndex := int(math.Floor(float64(seq.lastFrame) * seq.progress / float64(moment.TargetFPS())))

	if seq.direction == SequenceDirectionForward {
		seq.currentFrame = frameIndex
	} else {
		seq.currentFrame = seq.lastFrame - frameIndex
	}

	return nil
}

func (anim *Animator) switchSequenceDirection(seq *Sequence) {
	if seq.direction == SequenceDirectionForward {
		seq.direction = SequenceDirectionBackward
		return
	}

	seq.direction = SequenceDirectionForward
}
