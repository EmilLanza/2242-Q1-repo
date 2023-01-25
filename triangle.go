package main

type triangle struct{
	base float64
	hight float64
}

func (t triangle)Area() (float64, float64) {
	var a int
	a = t.base * t.hight / 2
	return a
}

func (t triangle)perimeter() (float64, float64){
	var p int
	p = t.base + t.hight + t.hight
	return p
}