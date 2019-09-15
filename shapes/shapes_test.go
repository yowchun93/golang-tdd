package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	want := 40.0
	got := Perimeter(rectangle)
	if want != got {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Version 1
// func TestArea(t *testing.T) {

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.1592653589793
// 		got := circle.Area()
// 		if want != got {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area()
// 		want := 100.0
// 		if want != got {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})
// }

// Version 2
// func TestArea(t *testing.T) {
// 	// creating Helper test function
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		checkArea(t, rectangle, 100.0)
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

// Version 3
// refactored using Table driven tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{10, 10}, 50.0},
	}

	// this is beautiful, i know for a fact what i am iterating through
	for _, test := range areaTests {
		got := test.shape.Area()
		want := test.want
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
}
