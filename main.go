package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/fe3dback/galaxy/registry"
)

// -- flags
var isProfiling = flag.Bool("profile", false, "run in profile mode")
var profilingPort = flag.Int("profileport", 15600, "http port for profiling")
var fullScreen = flag.Bool("fullscreen", false, "run in fullscreen mode")

func main() {
	runtime.LockOSThread()

	flag.Parse()
	flags := registry.Flags{
		IsProfiling:   *isProfiling,
		ProfilingPort: *profilingPort,
		FullScreen:    *fullScreen,
		Seed:          time.Now().UnixNano(),
	}

	provider := registry.NewProvider(flags)
	run(provider)
}

func run(provider *registry.Provider) {
	if provider.Registry.Game.Options.Debug.InProfiling {
		profile()
	}

	err := gameLoop(provider)
	if err != nil {
		panic(err)
	}

	log.Printf("params loop sucessfully ended")
}

func profile() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *profilingPort), nil)
		if err != nil {
			panic(fmt.Sprintf("can`t start http profiling tools at %d: %v", *profilingPort, err))
		}
	}()
}
