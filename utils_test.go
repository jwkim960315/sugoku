package main

import "testing"

func TestGeneratePossibleNumbers(t *testing.T) {
	rowArray := GeneratePossibleNumbers();

	for idx, elem := range rowArray {
		if elem != uint(idx + 1) {
			t.Errorf("elem: %v\nidx + 1: %v", elem, idx + 1);
		}
	}
}