package utils

import "fmt"

type (
	CloseFn func() error
	Closer  struct {
		queue []CloseFn
	}
)

func NewCloser() *Closer {
	return &Closer{
		queue: make([]CloseFn, 0),
	}
}

func (c *Closer) Enqueue(fn CloseFn) {
	c.queue = append(c.queue, fn)
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
