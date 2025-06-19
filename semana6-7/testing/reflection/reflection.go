package reflection

import (
	"fmt"
	"reflect"
)

func walk(x any, fn func(input string)) {
	val := newFunction(x)

	fmt.Println("valor de numField: ", val.NumField())
	for i := range val.NumField() {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func newFunction(x any) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
