package system

type (
	GameOptions struct {
		Debug  DebugOpt
		Frames FramesOpt
	}

	DebugOpt struct {
		InProfiling bool
		System      bool
		Frames      bool
		World       bool
	}

	FramesOpt struct {
		TargetFps int
	}
)
