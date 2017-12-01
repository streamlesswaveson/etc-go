package main

import "testing"

func Test2(t *testing.T) {
	result := fibDP(2)
	if result != 1 {
		t.Errorf("expected 1, got %v", result)
	}
}

func Test10(t *testing.T) {
	result := fibDP(10)
	if result != 55 {
		t.Errorf("expected 55, got %v", result)
	}
}

func Test93(t *testing.T){
	result := fibDP(93)
	if result != 12200160415121876738 {
		t.Errorf("expected 12200160415121876738 got %v", result)
	}
}