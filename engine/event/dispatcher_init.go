// This file generated at 2020-08-15 15:57:25.75767655 +0300 MSK m=+0.001468021
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