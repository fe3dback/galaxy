package animator

import (
	"math"

	"github.com/fe3dback/galaxy/engine"
)

func (anim *Animator) OnDraw(r engine.Renderer) error {
	if !anim.initialized {
		anim.initialize(r)
	}

	seq := anim.activeSequence
	res := seq.resource
	frame := seq.frames[seq.currentFrame]
	entityPos := anim.entity.Position()

	dest := engine.Rect{
		X: int(entityPos.X) + seq.offsetX - (frame.w / 2),
		Y: int(entityPos.Y) + seq.offsetY - (frame.h / 2),
		W: frame.w,
		H: frame.h,
	}

	r.DrawSpriteEx(res, frame.TextureRect(), dest, anim.entity.Rotation().ToFloat())
	r.DrawSquare(
		engine.ColorGreen,
		dest.X,
		dest.Y,
		dest.W,
		dest.H,
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
