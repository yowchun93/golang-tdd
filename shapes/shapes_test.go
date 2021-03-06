package shapes

import "testing"

// Version 1
// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {
// 	got := Area(12.0, 6.0)
// 	want := 72.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// Perimeter and Area using structs
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Version 2
// func TestArea(t *testing.T) {
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Area()
// 		want := 100.0

// 		if got != want {
// 			t.Errorf("got %.2f want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// Version 3
// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
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

// Table driven test
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
