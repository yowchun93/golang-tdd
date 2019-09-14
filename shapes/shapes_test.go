package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {
// 	// takes a Shape
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})
// t.Run("rectangle", func(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	got := rectangle.Area()
// 	want := 72.0
// 	if got != want {
// 		t.Errorf("got %g want %g", got, want)
// 	}
// })

// t.Run("circle", func(t *testing.T) {
// 	circle := Circle{10}
// 	got := circle.Area()
// 	want := 314.1592653589793
// 	if got != want {
// 		t.Errorf("got %g want %g", got, want)
// 	}
// })
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
