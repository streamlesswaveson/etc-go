package main

import "testing"

func TestNaive4(t *testing.T) {

	v := []int{1, 2, 3}
	M := int(4)

	result := naiveCoinChange(M, v, 0)
	if result != 4 {
		t.Errorf("expected 4 got %v", result)
	}
}

func TestDP4(t *testing.T) {
	v := []uint64{1, 2, 3}
	M := uint64(4)
	result := dpCoinChange(v, M)
	if result != 4 {
		t.Errorf("expected 4 got %v", result)
	}

}
