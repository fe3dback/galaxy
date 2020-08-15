package event

import (
	"container/list"
	"fmt"
)

type (
	handlerFn func(interface{}) error
	event     interface{}
	meta      struct {
		eventType eventType
		event     event
	}

	Dispatcher struct {
		queue    *list.List
		handlers map[eventType][]handlerFn
	}
)

func NewDispatcher() *Dispatcher {
	dispatcher := &Dispatcher{
		queue: list.New(),
	}

	dispatcher.init()
	return dispatcher
}

func (d *Dispatcher) Dispatch() {
	for d.queue.Len() > 0 {
		elm := d.queue.Front()
		meta := elm.Value.(meta)

		for _, handler := range d.handlers[meta.eventType] {
			err := handler(meta.event)
			if err != nil {
				panic(fmt.Sprintf("failed to handle event %d '%T': %v",
					meta.eventType,
					meta.event,
					err,
				))
			}
		}

		d.queue.Remove(elm)
	}
}

func (d *Dispatcher) publish(eventType eventType, event event) {
	d.queue.PushBack(meta{
		eventType: eventType,
		event:     event,
	})
}

func (d *Dispatcher) subscribe(eventType eventType, handler handlerFn) {
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}
