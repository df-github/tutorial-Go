package main

import "testing"

func TestIsEven(t *testing.T){
	result := isEven(3)
	if result {
		t.Errorf("Expected 3 not to be even but got true")
	}

	result = isEven(4)
	if !result {
		t.Errorf("Expected 3 not to be even but got true")
	}
}
