package main

import ("fmt")

func Area(side float64) float64 {
	return (side * side) / 2
}

func perimeter(side float64) float64 {
	return (side + side + side)
}

func main(){
	var sideOftriangle, answer float64
	sideOftriangle = 5
	answer = Area(sideOftriangle)
	answer = perimeter(sideOftriangle)
	fmt.Println(answer)
}