// This file generated at 2021-12-18 14:07:34.186813162 +0300 MSK m=+0.000783794
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // MouseMove Event Handler
    HandlerMouseMove func(mouseMoveEvent MouseMoveEvent) error
)

func (d *Dispatcher) PublishEventMouseMove(mouseMoveEvent MouseMoveEvent) {
    d.publish(eventTypeMouseMove, mouseMoveEvent)
}

func (d *Dispatcher) OnMouseMove(h HandlerMouseMove) {
    d.subscribe(eventTypeMouseMove, func(e interface{}) error {
        evMouseMove, ok := e.(MouseMoveEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnMouseMove` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evMouseMove)
    })
}