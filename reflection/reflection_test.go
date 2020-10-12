package reflection

import (
	"reflect"
	"testing"
)

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
				City string
				Age  int
			}{"Chris", "London", 3},
			[]string{"Chris", "London"},
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
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

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
