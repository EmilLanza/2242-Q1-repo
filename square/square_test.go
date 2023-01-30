package main

import "testing"

func TestArea(t *testing.T){
	got := Area(5)
	if got != 12.5{
	t.Error("Expected 12.5")
	}
}

func TestPerimeter(t *testing.T){
	got := perimeter(5)
	if got != 15{
	t.Error("Expected 15")
	}
}