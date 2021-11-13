package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/fe3dback/galaxy/di"
)

// -- flags
var isProfiling = flag.Bool("profile", false, "run in profile mode")
var profilingPort = flag.Int("profileport", 15600, "http port for profiling")
var fullScreen = flag.Bool("fullscreen", false, "run in fullscreen mode")

func main() {
	runtime.LockOSThread()
	container := setup()

	os.Exit(run(container))
}

func setup() *di.Container {
	flag.Parse()
	flags := di.NewInitFlags()
	flags.IsProfiling = *isProfiling
	flags.ProfilingPort = *profilingPort
	flags.FullScreen = *fullScreen

	return di.NewContainer(flags)
}

func run(container *di.Container) int {
	if container.Flags().IsProfiling {
		profile()
	}

	// core rand seed
	rand.Seed(container.Flags().Seed)

	// run game loop
	err := gameLoop(container)
	if err != nil {
		log.Println(fmt.Errorf("game loop exited with error: %w", err))
		return 1
	}

	log.Printf("game loop sucessfully ended")
	return 0
}

func profile() {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *profilingPort), nil)
		if err != nil {
			panic(fmt.Sprintf("can`t start http profiling tools at %d: %v", *profilingPort, err))
		}
	}()
}
