package sandbox

import "reflect"

type ComponentA struct {
}

type ComponentB struct {
}

func main() {
	a := reflect.TypeOf(ComponentA{})
	b := reflect.TypeOf(ComponentB{})

	_, _ = a, b
}
