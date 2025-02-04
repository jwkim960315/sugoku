package main

import (
	"testing"
)

var (
	FilledBoardData BoardData = BoardData{
		{CellData{5}, CellData{8}, CellData{6}, CellData{2}, CellData{9}, CellData{3}, CellData{7}, CellData{1}, CellData{4}},
		{CellData{7}, CellData{2}, CellData{4}, CellData{6}, CellData{1}, CellData{8}, CellData{5}, CellData{9}, CellData{3}},
		{CellData{9}, CellData{3}, CellData{1}, CellData{4}, CellData{7}, CellData{5}, CellData{6}, CellData{8}, CellData{2}},
		{CellData{2}, CellData{6}, CellData{7}, CellData{3}, CellData{8}, CellData{9}, CellData{1}, CellData{4}, CellData{5}},
		{CellData{1}, CellData{9}, CellData{5}, CellData{7}, CellData{4}, CellData{2}, CellData{3}, CellData{6}, CellData{8}},
		{CellData{8}, CellData{4}, CellData{3}, CellData{5}, CellData{6}, CellData{1}, CellData{2}, CellData{7}, CellData{9}},
		{CellData{4}, CellData{1}, CellData{2}, CellData{8}, CellData{3}, CellData{7}, CellData{9}, CellData{5}, CellData{6}},
		{CellData{3}, CellData{7}, CellData{8}, CellData{9}, CellData{5}, CellData{6}, CellData{4}, CellData{2}, CellData{1}},
		{CellData{6}, CellData{5}, CellData{9}, CellData{1}, CellData{2}, CellData{4}, CellData{8}, CellData{3}, CellData{7}},
	}
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

	if !IsValidRowForNumber(shuffledRow, 0, 8, 6) {
		t.Errorf("Row validation failed for valid row")
	}

	if IsValidRowForNumber(shuffledRow, 0, 8, 3) {
		t.Errorf("Row validation passed for invalid row")
	}
}

func TestIsValidColForNumber(t *testing.T) {
	shuffledCol := BoardData{
		{CellData{3}},
		{CellData{2}},
		{CellData{8}},
		{CellData{4}},
		{CellData{5}},
		{CellData{9}},
		{CellData{7}},
		{CellData{1}},
	}

	if !IsValidColForNumber(shuffledCol, 8, 0, 6) {
		t.Errorf("Column validation failed for valid column")
	}

	if IsValidColForNumber(shuffledCol, 8, 0, 3) {
		t.Errorf("Column validation passed for invalid column")
	}
}

func TestIsValidInnerGridForNumber(t *testing.T) {
	boardData := BoardData{
		{CellData{5}, CellData{8}, CellData{6}, CellData{2}, CellData{9}, CellData{3}, CellData{7}, CellData{1}, CellData{4}},
		{CellData{7}, CellData{2}, CellData{4}, CellData{6}, CellData{1}, CellData{8}, CellData{5}, CellData{9}, CellData{3}},
		{CellData{9}, CellData{3}, CellData{1}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
	}

	if !IsValidInnerGridForNumber(boardData, 2, 3, 4) {
		t.Errorf("Inner grid validation failed for valid number")
	}

	if IsValidInnerGridForNumber(boardData, 2, 3, 8) {
		t.Errorf("Inner grid validation passed for invalid number")
	}
}

func TestIsNumberValid(t *testing.T) {
	boardData := BoardData{
		{CellData{5}, CellData{8}, CellData{6}, CellData{2}, CellData{9}, CellData{3}, CellData{7}, CellData{1}, CellData{4}},
		{CellData{7}, CellData{2}, CellData{4}, CellData{6}, CellData{1}, CellData{8}, CellData{5}, CellData{9}, CellData{3}},
		{CellData{9}, CellData{3}, CellData{1}, CellData{4}, CellData{7}, CellData{5}, CellData{6}, CellData{8}, CellData{2}},
		{CellData{2}, CellData{6}, CellData{7}, CellData{3}, CellData{8}, CellData{9}, CellData{1}, CellData{4}, CellData{5}},
		{CellData{1}, CellData{9}, CellData{5}, CellData{7}, CellData{4}, CellData{2}, CellData{3}, CellData{6}, CellData{8}},
		{CellData{8}, CellData{4}, CellData{3}, CellData{5}, CellData{6}, CellData{1}, CellData{2}, CellData{7}, CellData{9}},
		{CellData{4}, CellData{1}, CellData{2}, CellData{8}, CellData{3}, CellData{7}, CellData{9}, CellData{5}, CellData{6}},
		{CellData{3}, CellData{7}, CellData{8}, CellData{9}, CellData{5}, CellData{6}, CellData{4}, CellData{2}, CellData{1}},
		{CellData{6}, CellData{5}, CellData{9}, CellData{1}, CellData{2}, CellData{4}, CellData{8}, CellData{3}, CellData{0}},
	}

	if !IsNumberValid(boardData, 8, 8, 7) {
		t.Error("Number validation fails with valid number")
	}

	if IsNumberValid(boardData, 8, 8, 4) {
		t.Error("Number validation passes with invalid number")
	}
}

func TestGenerateEmptyBoardData(t *testing.T) {
	boardData := GenerateEmptyBoardData()

	for rowIdx := range boardData {
		row := boardData
		for colIdx := range row {
			cellData := &boardData[rowIdx][colIdx]
			if cellData.Number != 0 {
				t.Errorf("Cell data value isn't zero:%v", cellData.Number)
			}
		}
	}
}

func TestFindNextEmptyCellPos(t *testing.T) {
	boardData1 := GenerateEmptyBoardData()
	boardData2 := BoardData{
		{CellData{5}, CellData{8}, CellData{6}, CellData{2}, CellData{9}, CellData{3}, CellData{7}, CellData{1}, CellData{4}},
		{CellData{7}, CellData{2}, CellData{4}, CellData{6}, CellData{1}, CellData{8}, CellData{5}, CellData{9}, CellData{3}},
		{CellData{9}, CellData{3}, CellData{1}, CellData{4}, CellData{7}, CellData{5}, CellData{6}, CellData{8}, CellData{2}},
		{CellData{2}, CellData{6}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
		{CellData{1}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
		{CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
		{CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
		{CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
		{CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}, CellData{0}},
	}
	boardData3 := FilledBoardData

	emptyCellPos1 := FindNextEmptyCellPos(boardData1)
	emptyCellPos2 := FindNextEmptyCellPos(boardData2)
	emptyCellPos3 := FindNextEmptyCellPos(boardData3)

	if emptyCellPos1.RowIdx != 0 || emptyCellPos1.ColIdx != 0 {
		t.Errorf("Incorrect empty cell position: (%v,%v)", emptyCellPos1.RowIdx, emptyCellPos1.ColIdx)
	}

	if emptyCellPos2.RowIdx != 3 || emptyCellPos2.ColIdx != 2 {
		t.Errorf("Incorrect empty cell position: (%v,%v)", emptyCellPos2.RowIdx, emptyCellPos2.ColIdx)
	}

	if emptyCellPos3 != nil {
		t.Errorf("Found empty cell position with fully filled board data")
	}
}

func TestGenerateFilledBoardData(t *testing.T) {
	boardData := GenerateFilledBoardData()

	for rowIdx, row := range boardData {
		for colIdx := range row {
			cellData := &row[colIdx]

			if !IsNumberValid(boardData, uint(rowIdx), uint(colIdx), cellData.Number) {
				t.Errorf("\nInvalid number %v at (%v,%v)\n%v", cellData.Number, rowIdx, colIdx, PrintBoardData(boardData))
			}
		}
	}
}
