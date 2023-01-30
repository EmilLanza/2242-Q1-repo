package main

import "testing"

func TestArea(t *testing.T){
	got := Area()
	if got != 156{
		t.Error("Expected 156")
	}
}

	func Testperimeter(t *testing.T){
		got := perimeter()
	if got != 8{
		t.Error("Expected 8")
	}
	}