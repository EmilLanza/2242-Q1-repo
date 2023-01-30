package main

import ("fmt")

type Triangle struct{
	base, hight float64
}

func (t Triangle) Area() float64{
	return (t.base * t.hight) / 2
}

func (t Triangle) perimeter() float64{
	return (t.base + t.hight + t.hight)
}

func main(){
	area := Triangle{
		base: 24,
		hight: 13,
	}

	perimeter := Triangle{
		base: 2,
		hight: 3,
	}

	fmt.Println("Area", area.Area(), "perimeter", perimeter.perimeter())
}