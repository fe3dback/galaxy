package galaxy

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"

	"github.com/fe3dback/galaxy/cfg"
	"github.com/fe3dback/galaxy/internal/di"
)

type (
	Game struct {
		cfg       *cfg.InitFlags
		container *di.Container
	}
)

func NewGame(cfg *cfg.InitFlags) *Game {
	return &Game{
		cfg: cfg,
	}
}

func (g *Game) Run() int {
	runtime.LockOSThread()
	g.container = g.setup()
	return g.run()
}

func (g *Game) setup() *di.Container {
	return di.NewContainer(g.cfg)
}

func (g *Game) run() int {
	if g.container.Flags().IsProfiling() {
		profile(g.container.Flags().ProfilingPort())
	}

	// core rand seed
	rand.Seed(g.container.Flags().Seed())

	// run game loop
	// g.container.Logger().In
	err := gameLoop(g)
	if err != nil {
		log.Println(fmt.Errorf("game loop exited with error: %w", err))
		return 1
	}

	log.Printf("game loop sucessfully ended")
	return 0
}

func profile(port int) {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
		if err != nil {
			panic(fmt.Sprintf("can`t start http profiling tools at %d: %v", port, err))
		}
	}()
}
