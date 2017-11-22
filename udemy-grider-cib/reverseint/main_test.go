package main

import "testing"

func TestReverse0(t *testing.T) {
	assertReverse(0,0,t)
}

func TestReversePositive(t *testing.T) {
	assertReverse(123,321, t)
	assertReverse(100,1,t)
}

func TestReverseNegative(t *testing.T) {
	assertReverse(-123, -321, t)
	assertReverse(-90, -9, t)
}

func assertReverse(input int, expected int, t *testing.T) {
	actual := reverse(input)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
