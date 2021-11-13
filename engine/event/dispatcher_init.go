// This file generated at 2021-11-13 20:45:50.021970524 +0300 MSK m=+0.000773339
// DO NOT MODIFY
package event

func (d *Dispatcher) init() {
	d.handlers = make(map[eventType][]handlerFn)

	d.handlers[eventTypeQuit] = make([]handlerFn, 0)

	d.handlers[eventTypeKeyBoard] = make([]handlerFn, 0)

	d.handlers[eventTypeWindow] = make([]handlerFn, 0)

	d.handlers[eventTypeMouseWheel] = make([]handlerFn, 0)

	d.handlers[eventTypeFrameStart] = make([]handlerFn, 0)

	d.handlers[eventTypeFrameEnd] = make([]handlerFn, 0)

	d.handlers[eventTypeCameraUpdate] = make([]handlerFn, 0)

}
