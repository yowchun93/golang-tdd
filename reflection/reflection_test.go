package reflection

import (
	"reflect"
	"testing"
)

// Look up GoLang structs
// https://gobyexample.com/structs

// Finding a way to use polymorphic function
// func TestWalk(t *testing.T) {

// 	expected := "Chris"
// 	var got []string

// 	x := struct {
// 		Name string
// 	}{expected}

// 	walk(x, func(input string) {
// 		got = append(got, input)
// 	})

// 	// if len(got) != 1 {
// 	// 	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
// 	// }
// 	if got[0] != expected {
// 		t.Errorf("got %q, want %q", got[0], expected)
// 	}
// }

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	// Writing this test, is something i want to get good at.
	// Gonna need lots of repetition
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
