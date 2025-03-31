package my_interface

import "fmt"

//Interface is the collection of singnature of method
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length  int
	Breadth int
}

func (r Rectangle) Area() float64 {
	return float64(r.Length) * float64(r.Breadth)
}

func (r Rectangle) Perimeter() float64 {
	return 2*float64(r.Length) + float64(r.Breadth)
}

func (r Rectangle) DisplayRectangle() {
	fmt.Println("length=", r.Length, "Breadth=", r.Breadth)

}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * float64(c.Radius)
}

func DisplayAreaAndPerimeter(s Shape) {
	area := s.Area()
	perimeter := s.Perimeter()
	fmt.Println("Area:", area, "Perimeter", perimeter)
}
