package main

type (
	gameOptions struct {
		debug  debugOpt
		frames framesOpt
	}

	debugOpt struct {
		inProfiling bool
		system      bool
		frames      bool
		world       bool
	}

	framesOpt struct {
		targetFps int
	}
)
