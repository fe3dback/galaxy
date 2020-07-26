package main

import (
	"fmt"
)

type (
	closeFn func() error
	closer  struct {
		queue []closeFn
	}
)

func newCloser() *closer {
	return &closer{
		queue: make([]closeFn, 0),
	}
}

func (c *closer) Enqueue(fn closeFn) {
	c.queue = append(c.queue, fn)
}

func (c *closer) Close() error {
	for _, closeFn := range c.queue {
		err := closeFn()
		if err != nil {
			fmt.Printf("close error: %v", err)
		}
	}

	return nil
}

func formatBytes(b uint64) string {
	const unit = uint64(1000)
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1d %cB", b/div, "kMGTPE"[exp])
}
