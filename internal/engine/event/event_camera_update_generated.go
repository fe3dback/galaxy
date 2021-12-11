// This file generated at 2020-08-15 15:57:25.757119062 +0300 MSK m=+0.000910572
// DO NOT MODIFY
package event

import (
    "fmt"
    "reflect"
)

type (
    // CameraUpdate Event Handler
    HandlerCameraUpdate func(cameraUpdateEvent CameraUpdateEvent) error
)

func (d *Dispatcher) PublishEventCameraUpdate(cameraUpdateEvent CameraUpdateEvent) {
    d.publish(eventTypeCameraUpdate, cameraUpdateEvent)
}

func (d *Dispatcher) OnCameraUpdate(h HandlerCameraUpdate) {
    d.subscribe(eventTypeCameraUpdate, func(e interface{}) error {
        evCameraUpdate, ok := e.(CameraUpdateEvent)

        if !ok {
            panic(fmt.Sprintf("can`t handle `OnCameraUpdate` unexpected event type `%s`", reflect.TypeOf(e)))
        }

        return h(evCameraUpdate)
    })
}