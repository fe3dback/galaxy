// This file generated at 2020-08-15 14:53:05.971524957 +0300 MSK m=+0.001084917
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // Window Event Handler
    HandlerWindow func(windowEvent WindowEvent) error
)

func (d *Dispatcher) PublishEventWindow(windowEvent WindowEvent) {
    d.publish(eventTypeWindow, windowEvent)
}

func (d *Dispatcher) OnWindow(h HandlerWindow) {
    d.subscribe(eventTypeWindow, func(e interface{}) error {
        evWindow, ok := e.(WindowEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnWindow` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evWindow)
    })
}