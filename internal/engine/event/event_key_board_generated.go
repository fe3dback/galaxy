// This file generated at 2020-08-15 14:53:05.971298031 +0300 MSK m=+0.000857991
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // KeyBoard Event Handler
    HandlerKeyBoard func(keyBoardEvent KeyBoardEvent) error
)

func (d *Dispatcher) PublishEventKeyBoard(keyBoardEvent KeyBoardEvent) {
    d.publish(eventTypeKeyBoard, keyBoardEvent)
}

func (d *Dispatcher) OnKeyBoard(h HandlerKeyBoard) {
    d.subscribe(eventTypeKeyBoard, func(e interface{}) error {
        evKeyBoard, ok := e.(KeyBoardEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnKeyBoard` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evKeyBoard)
    })
}