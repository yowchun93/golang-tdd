package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

// func walk(x interface{}, fn func(input string)) {
// 	// call the anonymous function, passing in the argument
// 	fn("Random words")
// }

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)

// 	if val.Kind() == reflect.Slice {
// 		for i := 0; i < val.Len(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 		return
// 	}

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)

// 		switch field.Kind() {
// 		case reflect.String:
// 			fn(field.String())
// 		case reflect.Struct:
// 			walk(field.Interface(), fn)
// 		}
// 	}
// }

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	return val
}
