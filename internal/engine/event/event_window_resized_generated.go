// This file generated at 2021-12-22 17:38:19.28178361 +0300 MSK m=+0.000562602
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // WindowResized Event Handler
    HandlerWindowResized func(windowResizedEvent WindowResizedEvent) error
)

func (d *Dispatcher) PublishEventWindowResized(windowResizedEvent WindowResizedEvent) {
    d.publish(eventTypeWindowResized, windowResizedEvent)
}

func (d *Dispatcher) OnWindowResized(h HandlerWindowResized) {
    d.subscribe(eventTypeWindowResized, func(e interface{}) error {
        evWindowResized, ok := e.(WindowResizedEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnWindowResized` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evWindowResized)
    })
}