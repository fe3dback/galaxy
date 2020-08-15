// This file generated at 2020-08-15 14:53:05.970902818 +0300 MSK m=+0.000462792
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // Quit Event Handler
    HandlerQuit func(quitEvent QuitEvent) error
)

func (d *Dispatcher) PublishEventQuit(quitEvent QuitEvent) {
    d.publish(eventTypeQuit, quitEvent)
}

func (d *Dispatcher) OnQuit(h HandlerQuit) {
    d.subscribe(eventTypeQuit, func(e interface{}) error {
        evQuit, ok := e.(QuitEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnQuit` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evQuit)
    })
}