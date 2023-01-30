package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * (c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	area := Circle{
		radius: 3,
	}

	perimeter := Circle{
		radius: 7,
	}

	fmt.Println("Area of circle", area.Area(), "peremeter", perimeter.Perimeter())
}
