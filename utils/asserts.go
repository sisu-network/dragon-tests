package utils

import (
	"fmt"
	"reflect"
)

func Assert(expected interface{}, actual interface{}, message string) {
	if expected == actual {
		return
	}

	expectedType := fmt.Sprintf("%T", expected)
	actualType := fmt.Sprintf("%T", actual)

	if expectedType != actualType {
		fmt.Println("Expect type ", expectedType, "but get actual Type", actualType)
		panic(message)
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)

	if expectedValue != actualValue {
		fmt.Println("Expect value ", expectedValue, "but get actual value", actualValue)
		panic(message)
	}
}
