// This file generated at 2021-12-18 14:07:34.18743168 +0300 MSK m=+0.001402307
// DO NOT MODIFY
package event

func (d *Dispatcher) init() {
    d.handlers = make(map[eventType][]handlerFn)
    
    d.handlers[eventTypeQuit] = make([]handlerFn, 0)
    
    d.handlers[eventTypeKeyBoard] = make([]handlerFn, 0)
    
    d.handlers[eventTypeWindow] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseButton] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseWheel] = make([]handlerFn, 0)
    
    d.handlers[eventTypeMouseMove] = make([]handlerFn, 0)
    
    d.handlers[eventTypeFrameStart] = make([]handlerFn, 0)
    
    d.handlers[eventTypeFrameEnd] = make([]handlerFn, 0)
    
    d.handlers[eventTypeCameraUpdate] = make([]handlerFn, 0)
    
}