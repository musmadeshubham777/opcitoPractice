package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radious float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Circle) Area() float64 {
	return math.Pi * r.radious * r.radious
}
func (r Circle) Perimeter() float64 {
	d := r.radious + r.radious
	return math.Pi * d
}
func (r Rectangle) Perimeter() float64 {
	w := r.width + r.width + r.height + r.height

	return w
}

func CalculateArea(s Shape) (float64, float64) {
	area := s.Area()
	perimeter := s.Perimeter()
	return area, perimeter
}

func main() {
	rect := Rectangle{width: 10, height: 20}
	cicl := Circle{radious: 20}

	// fmt.Println("Area of Rectangle", CalculateArea(rect))
	// fmt.Println("Area of Circle", CalculateArea(cicl))

	area, perimeter := CalculateArea(rect)
	fmt.Println("Rectangle Area:", area, "Perimeter:", perimeter)

	carea, cperimeter := CalculateArea(cicl)
	fmt.Println("Circle Area:", carea, "Perimeter:", cperimeter)

}
