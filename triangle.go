package main

type triangle struct{
	base float64
	hight float64
}

func (t triangle)Area() (float64, float64) {
	return t.base * t.hight / 2
}