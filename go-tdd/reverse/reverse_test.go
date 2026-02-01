package main

import "testing"

func TestReverse(t *testing.T) {
	result := Reverse("Hello")

	if result != "olleH" {
		t.Errorf("Expected 'olleH', Got '%s'", result)
	}
}
func TestReverseEmptystring(t *testing.T) {
	result := Reverse("")

	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}
func TestReverseUnicode(t *testing.T) {
	result := Reverse("世")

	if result != "世" {
		t.Errorf("expected '世', got '%s'", result)
	}
}

//This was where we left off yesterday
//lets see some more tests for the same function
//just know we need to convert to runes first before reversing then we can reverse sucessfuly   cause the byte version of 世 is not the same as the rune 世 and strings are represented in bytes
//forgot to save first sorry for that lets see now
