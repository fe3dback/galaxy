// This file generated at 2021-12-22 17:39:40.389672587 +0300 MSK m=+0.000393214
// DO NOT MODIFY
package event

func (d *Dispatcher) init() {
    d.handlers = make(map[eventType][]handlerFn)
    
    d.handlers[eventTypeQuit] = make([]handlerFn, 0)
    
    d.handlers[eventTypeKeyBoard] = make([]handlerFn, 0)
    
    d.handlers[eventTypeWindowResized] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseButton] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseWheel] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseMove] = make([]handlerFn, 0)
    
    d.handlers[eventTypeFrameStart] = make([]handlerFn, 0)
    
    d.handlers[eventTypeFrameEnd] = make([]handlerFn, 0)
    
    d.handlers[eventTypeCameraUpdate] = make([]handlerFn, 0)
    
}