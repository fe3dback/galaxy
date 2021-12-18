package event

type (
	FrameEndEvent struct {
		FrameID   uint64
		DeltaTime float64
	}
)
