// This file generated at 2020-08-15 15:11:56.506103411 +0300 MSK m=+0.001230698
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // FrameEnd Event Handler
    HandlerFrameEnd func(frameEndEvent FrameEndEvent) error
)

func (d *Dispatcher) PublishEventFrameEnd(frameEndEvent FrameEndEvent) {
    d.publish(eventTypeFrameEnd, frameEndEvent)
}

func (d *Dispatcher) OnFrameEnd(h HandlerFrameEnd) {
    d.subscribe(eventTypeFrameEnd, func(e interface{}) error {
        evFrameEnd, ok := e.(FrameEndEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnFrameEnd` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evFrameEnd)
    })
}