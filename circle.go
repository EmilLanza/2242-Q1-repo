package main

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * (radius * radius)
}
