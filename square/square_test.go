package main

import "testing"

func TestArea(t *testing.T){
	got := Area(side)
	if got != 12.5 {
	  t.Error("Expexted 28.26")
	}
}

func Testperimeter(t *testing.T){
	got := perimeter(side)
	if got != 15{
		t.Error("Expexted 43.96")
	}
}