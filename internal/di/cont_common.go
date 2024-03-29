package di

import (
	"os"

	"go.uber.org/zap"

	"github.com/fe3dback/galaxy/internal/utils"
)

func (c *Container) Closer() *utils.Closer {
	if c.memstate.closer != nil {
		return c.memstate.closer
	}

	c.memstate.closer = utils.NewCloser()
	return c.memstate.closer
}

func (c *Container) logger() *zap.SugaredLogger {
	if c.memstate.logger != nil {
		return c.memstate.logger
	}

	const (
		logPathDebug  = "./galaxy-debug.log"
		logPathErrors = "./galaxy-error.log"
	)

	_ = os.Remove(logPathDebug)
	_ = os.Remove(logPathErrors)

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "console",
		EncoderConfig:     zap.NewDevelopmentEncoderConfig(),
		OutputPaths:       []string{"stdout", logPathDebug},
		ErrorOutputPaths:  []string{"stderr", logPathErrors},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic("failed to create logger")
	}

	closedStd := zap.RedirectStdLog(logger)
	restoreLogger := zap.ReplaceGlobals(logger)

	closer := c.Closer()
	// closer.EnqueueClose(logger.Sync) // todo
	closer.EnqueueFree(closedStd)
	closer.EnqueueFree(restoreLogger)

	c.memstate.logger = logger.Sugar()
	return c.memstate.logger
}
