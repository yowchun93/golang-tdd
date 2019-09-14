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

// Unlike traditional OOP namely Java, C# , you do not have to explicitly say type Foo implements
type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(r Rectangle) float64 {
	return (2 * (r.Width + r.Height))
}

// func Area(r Rectangle) float64 {
// 	return r.Width * r.Height
// }
