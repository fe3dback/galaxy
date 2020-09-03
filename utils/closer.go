package utils

import "fmt"

type (
	CloseFn func() error
	FreeFn  func()
	Closer  struct {
		queue []CloseFn
	}
)

func NewCloser() *Closer {
	return &Closer{
		queue: make([]CloseFn, 0),
	}
}

func (c *Closer) EnqueueClose(fn CloseFn) {
	c.queue = append(c.queue, fn)
}

func (c *Closer) EnqueueFree(fn FreeFn) {
	c.EnqueueClose(func() error {
		fn()

		return nil
	})
}

func (c *Closer) Close() error {
	for _, closeFn := range c.queue {
		err := closeFn()
		if err != nil {
			fmt.Printf("close error: %v", err)
		}
	}

	return nil
}
