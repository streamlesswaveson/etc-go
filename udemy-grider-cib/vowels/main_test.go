package main

import "testing"

func TestAll(t *testing.T) {

	result := vowels("aeiou")
	if result != 5 {
		t.Errorf("Expected 5 got %v", result)
	}

}

func TestUpperAll(t *testing.T) {
	result := vowels("AEIOU")
	if result != 5 {
		t.Errorf("Expected 5 got %v", result)
	}
}
