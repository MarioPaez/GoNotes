package structs

import "math"

const PI = math.Pi

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func (t Triangle) Area() float64 {
	return (t.Height * t.Base) / 2
}
