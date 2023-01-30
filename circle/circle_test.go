package main

import "testing"

func TestArea(t *testing.T){
	got := Area()
	if got != 28.26 {
	  t.Error("Expexted 28.26")
	}
}

func TestPerimeter(t *testing.T){
	got := Perimeter()
	if got != 43.96{
		t.Error("Expexted 43.96")
	}
}