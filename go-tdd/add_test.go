// Today  I will try and write my first test for a simple add programm
package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)

	if sum != 5 {
		t.Errorf("Expected 5, Got %d", sum)
	}
}
func TestAddWithNegativeNumbers(t *testing.T) {
	sum := Add(-2, 3)
	if sum != 1 {
		t.Errorf("Expected -1, Got %d", sum)
	}
}

//lets add another test case
//this test should fail till we write a function that fulfills the expectations
