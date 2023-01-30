package main

import ("fmt")

func Area(side float64) float64 {
	return (side * side) / 2
}

func perimeter(side float64) float64 {
	return (side + side + side)
}

func main(){
	var sideOftriangle, firstanswer, secondanswer float64
	sideOftriangle = 5
	firstanswer = Area(sideOftriangle)
	secondanswer = perimeter(sideOftriangle)
	fmt.Println("Area", firstanswer, "Perimeter", secondanswer)
}