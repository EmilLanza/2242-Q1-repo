package main

import ("fmt")

type Triangle struct{
	base, hight float64
}

func (t Triangle)Area() (float64) {
	return (t.base * t.hight) / 2
}

func (t Triangle)perimeter() (float64){
	return (t.base + t.hight + t.hight)
}

