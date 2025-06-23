package reflection

import (
	"reflect"
)

func walk(x any, fn func(input string)) {
	val := calculateVal(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := range val.NumField() {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := range val.Len() {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func calculateVal(x any) reflect.Value {
	val := reflect.ValueOf(x) // Nos da el valor del tipo.
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
