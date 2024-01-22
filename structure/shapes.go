package structure

import "math"

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base float64
	height float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func(r Rectangle) Area() float64 {
	return r.height * r.width
}

func(r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func(c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func(c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func(t Triangle) Area() float64{
	return 0.5 * t.base * t.height
}

func(t Triangle) Perimeter() float64{
	return 0
}