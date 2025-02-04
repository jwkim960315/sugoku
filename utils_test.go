package main

import (
	"testing"
)

func TestGeneratePossibleNumbers(t *testing.T) {
	rowArray := GeneratePossibleNumbers()

	for idx, elem := range rowArray {
		if elem != uint(idx+1) {
			t.Errorf("elem: %v\nidx + 1: %v", elem, idx+1)
		}
	}
}

func TestGenerateRandomRowNumbers(t *testing.T) {
	rowArray := GeneratePossibleNumbers()
	shuffledRowArray := GenerateRandomRowNumbers(rowArray[:])

	if len(rowArray) != len(shuffledRowArray) {
		t.Errorf("\nOriginal array has a different length than the shuffled array\nOriginal array length: %v\nShuffled array length: %v", len(rowArray), len(shuffledRowArray))
	}

	var rowArraySum, shuffledRowArraySum uint
	for _, elem := range rowArray {
		rowArraySum += elem
	}

	for _, elem := range shuffledRowArray {
		shuffledRowArraySum += elem
	}

	if rowArraySum != shuffledRowArraySum {
		t.Errorf("\nOriginal array: %v\nShuffled array: %v", rowArray, shuffledRowArray)
	}
}

func TestIsValidRowForNumber(t *testing.T) {
	shuffledRow := BoardData{
		{
			CellData{3},
			CellData{2},
			CellData{8},
			CellData{4},
			CellData{5},
			CellData{9},
			CellData{7},
			CellData{1},
		},
	}

	if !IsValidRowForNumber(shuffledRow, 0, 6) {
		t.Errorf("Row validation failed for valid row")
	}

	if IsValidRowForNumber(shuffledRow, 0, 3) {
		t.Errorf("Row validation passed for invalid row")
	}
}
