// This file generated at 2020-08-15 15:11:56.505562455 +0300 MSK m=+0.000689745
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // FrameStart Event Handler
    HandlerFrameStart func(frameStartEvent FrameStartEvent) error
)

func (d *Dispatcher) PublishEventFrameStart(frameStartEvent FrameStartEvent) {
    d.publish(eventTypeFrameStart, frameStartEvent)
}

func (d *Dispatcher) OnFrameStart(h HandlerFrameStart) {
    d.subscribe(eventTypeFrameStart, func(e interface{}) error {
        evFrameStart, ok := e.(FrameStartEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnFrameStart` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evFrameStart)
    })
}