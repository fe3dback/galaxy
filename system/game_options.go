package system

type (
	GameOptions struct {
		Debug  DebugOpt
		Frames FramesOpt
	}

	DebugOpt struct {
		InProfiling bool
		System      bool
		Memory      bool
		Frames      bool
		World       bool
	}

	FramesOpt struct {
		TargetFps int
	}
)
