package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/fe3dback/galaxy/registry"
)

// -- flags
var isProfiling = flag.Bool("profile", false, "run in profile mode")
var profilingPort = flag.Int("profileport", 15600, "http port for profiling")

func main() {
	runtime.LockOSThread()

	flag.Parse()
	flags := registry.Flags{
		IsProfiling:   *isProfiling,
		ProfilingPort: *profilingPort,
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
