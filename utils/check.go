package utils

import "fmt"

func Check(operation string, err error) {
	if err == nil {
		return
	}

	panic(fmt.Sprintf("%s: %v", operation, err))
}

func CheckPanic(context string) {
	if data := recover(); data != nil {
		err := fmt.Errorf("recovered in %s: %v", context, data)
		panic(err)
	}
}
