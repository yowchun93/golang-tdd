package reflection

import "reflect"

//  Implementation 1
// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)
// 	field := val.Field(0)
// 	fn(field.String())
// }

// Implementation 2
// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		if field.Kind() == reflect.String {
// 			fn(field.String())
// 		}

// 		if field.Kind() == reflect.Struct {
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

// Implementation 3
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// pointers needs to be handled a little differently
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
