package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

// Perimeter returns a perimeter for a shape
// func Perimeter(width float64, height float64) float64 {
// 	return (2 * width) + (2 * height)
// }

// // Area returns area for a shape
// func Area(width float64, height float64) float64 {
// 	return width * height
// }

// Version 2
// This would not work if there is another type Shape
// Perimeter returns a perimeter for a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// // Area returns a area for a rectangle
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

// Version 3
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
