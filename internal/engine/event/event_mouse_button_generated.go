// This file generated at 2021-12-09 22:45:44.775072098 +0300 MSK m=+0.000345509
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // MouseButton Event Handler
    HandlerMouseButton func(mouseButtonEvent MouseButtonEvent) error
)

func (d *Dispatcher) PublishEventMouseButton(mouseButtonEvent MouseButtonEvent) {
    d.publish(eventTypeMouseButton, mouseButtonEvent)
}

func (d *Dispatcher) OnMouseButton(h HandlerMouseButton) {
    d.subscribe(eventTypeMouseButton, func(e interface{}) error {
        evMouseButton, ok := e.(MouseButtonEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnMouseButton` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evMouseButton)
    })
}