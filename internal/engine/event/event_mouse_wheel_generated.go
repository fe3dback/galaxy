// This file generated at 2020-08-15 14:53:05.97174567 +0300 MSK m=+0.001305619
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // MouseWheel Event Handler
    HandlerMouseWheel func(mouseWheelEvent MouseWheelEvent) error
)

func (d *Dispatcher) PublishEventMouseWheel(mouseWheelEvent MouseWheelEvent) {
    d.publish(eventTypeMouseWheel, mouseWheelEvent)
}

func (d *Dispatcher) OnMouseWheel(h HandlerMouseWheel) {
    d.subscribe(eventTypeMouseWheel, func(e interface{}) error {
        evMouseWheel, ok := e.(MouseWheelEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnMouseWheel` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evMouseWheel)
    })
}