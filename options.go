package main

type (
	options struct {
		debug  debugOpt
		frames framesOpt
	}

	debugOpt struct {
		frames bool
	}

	framesOpt struct {
		targetFps int64
	}
)
