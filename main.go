package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// -- flags
var isProfiling = flag.Bool("profile", false, "run in profile mode")
var profilingPort = flag.Int("profileport", 15600, "http port for profiling")

func main() {
	runtime.LockOSThread()

	flag.Parse()
	factory := newFactory()

	run(factory.provideGameParams())
}

func run(params *gameParams) {
	if params.options.debug.inProfiling {
		profile()
	}

	err := gameLoop(params)
	if err != nil {
		panic(err)
	}

	fmt.Printf("params loop sucessfully ended\n")
}

func profile() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *profilingPort), nil)
		if err != nil {
			panic(fmt.Sprintf("can`t start http profiling tools at %d: %v", *profilingPort, err))
		}
	}()
}
