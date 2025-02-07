package main

import (
	"testing"
)

var (
	FilledBoardData BoardData = BoardData{
		{CellData{5, false}, CellData{8, false}, CellData{6, false}, CellData{2, false}, CellData{9, false}, CellData{3, false}, CellData{7, false}, CellData{1, false}, CellData{4, false}},
		{CellData{7, false}, CellData{2, false}, CellData{4, false}, CellData{6, false}, CellData{1, false}, CellData{8, false}, CellData{5, false}, CellData{9, false}, CellData{3, false}},
		{CellData{9, false}, CellData{3, false}, CellData{1, false}, CellData{4, false}, CellData{7, false}, CellData{5, false}, CellData{6, false}, CellData{8, false}, CellData{2, false}},
		{CellData{2, false}, CellData{6, false}, CellData{7, false}, CellData{3, false}, CellData{8, false}, CellData{9, false}, CellData{1, false}, CellData{4, false}, CellData{5, false}},
		{CellData{1, false}, CellData{9, false}, CellData{5, false}, CellData{7, false}, CellData{4, false}, CellData{2, false}, CellData{3, false}, CellData{6, false}, CellData{8, false}},
		{CellData{8, false}, CellData{4, false}, CellData{3, false}, CellData{5, false}, CellData{6, false}, CellData{1, false}, CellData{2, false}, CellData{7, false}, CellData{9, false}},
		{CellData{4, false}, CellData{1, false}, CellData{2, false}, CellData{8, false}, CellData{3, false}, CellData{7, false}, CellData{9, false}, CellData{5, false}, CellData{6, false}},
		{CellData{3, false}, CellData{7, false}, CellData{8, false}, CellData{9, false}, CellData{5, false}, CellData{6, false}, CellData{4, false}, CellData{2, false}, CellData{1, false}},
		{CellData{6, false}, CellData{5, false}, CellData{9, false}, CellData{1, false}, CellData{2, false}, CellData{4, false}, CellData{8, false}, CellData{3, false}, CellData{7, false}},
	}
)

func TestGeneratePossibleNumbers(t *testing.T) {
	rowArray := GeneratePossibleNumbers()

	for idx, elem := range rowArray {
		if elem != idx+1 {
			t.Errorf("elem: %v\nidx + 1: %v", elem, idx+1)
		}
	}
}

func TestShuffleSlice(t *testing.T) {
	rowArray := GeneratePossibleNumbers()
	shuffledRowArray := ShuffleSlice(rowArray[:])

	if len(rowArray) != len(shuffledRowArray) {
		t.Errorf("\nOriginal array has a different length than the shuffled array\nOriginal array length: %v\nShuffled array length: %v", len(rowArray), len(shuffledRowArray))
	}

	var rowArraySum, shuffledRowArraySum int
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
			CellData{3, false},
			CellData{2, false},
			CellData{8, false},
			CellData{4, false},
			CellData{5, false},
			CellData{9, false},
			CellData{7, false},
			CellData{1, false},
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
		{CellData{3, false}},
		{CellData{2, false}},
		{CellData{8, false}},
		{CellData{4, false}},
		{CellData{5, false}},
		{CellData{9, false}},
		{CellData{7, false}},
		{CellData{1, false}},
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
		{CellData{5, false}, CellData{8, false}, CellData{6, false}, CellData{2, false}, CellData{9, false}, CellData{3, false}, CellData{7, false}, CellData{1, false}, CellData{4, false}},
		{CellData{7, false}, CellData{2, false}, CellData{4, false}, CellData{6, false}, CellData{1, false}, CellData{8, false}, CellData{5, false}, CellData{9, false}, CellData{3, false}},
		{CellData{9, false}, CellData{3, false}, CellData{1, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
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
		{CellData{5, false}, CellData{8, false}, CellData{6, false}, CellData{2, false}, CellData{9, false}, CellData{3, false}, CellData{7, false}, CellData{1, false}, CellData{4, false}},
		{CellData{7, false}, CellData{2, false}, CellData{4, false}, CellData{6, false}, CellData{1, false}, CellData{8, false}, CellData{5, false}, CellData{9, false}, CellData{3, false}},
		{CellData{9, false}, CellData{3, false}, CellData{1, false}, CellData{4, false}, CellData{7, false}, CellData{5, false}, CellData{6, false}, CellData{8, false}, CellData{2, false}},
		{CellData{2, false}, CellData{6, false}, CellData{7, false}, CellData{3, false}, CellData{8, false}, CellData{9, false}, CellData{1, false}, CellData{4, false}, CellData{5, false}},
		{CellData{1, false}, CellData{9, false}, CellData{5, false}, CellData{7, false}, CellData{4, false}, CellData{2, false}, CellData{3, false}, CellData{6, false}, CellData{8, false}},
		{CellData{8, false}, CellData{4, false}, CellData{3, false}, CellData{5, false}, CellData{6, false}, CellData{1, false}, CellData{2, false}, CellData{7, false}, CellData{9, false}},
		{CellData{4, false}, CellData{1, false}, CellData{2, false}, CellData{8, false}, CellData{3, false}, CellData{7, false}, CellData{9, false}, CellData{5, false}, CellData{6, false}},
		{CellData{3, false}, CellData{7, false}, CellData{8, false}, CellData{9, false}, CellData{5, false}, CellData{6, false}, CellData{4, false}, CellData{2, false}, CellData{1, false}},
		{CellData{6, false}, CellData{5, false}, CellData{9, false}, CellData{1, false}, CellData{2, false}, CellData{4, false}, CellData{8, false}, CellData{3, false}, CellData{0, false}},
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
		for colIdx := range boardData[rowIdx] {
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
		{CellData{5, false}, CellData{8, false}, CellData{6, false}, CellData{2, false}, CellData{9, false}, CellData{3, false}, CellData{7, false}, CellData{1, false}, CellData{4, false}},
		{CellData{7, false}, CellData{2, false}, CellData{4, false}, CellData{6, false}, CellData{1, false}, CellData{8, false}, CellData{5, false}, CellData{9, false}, CellData{3, false}},
		{CellData{9, false}, CellData{3, false}, CellData{1, false}, CellData{4, false}, CellData{7, false}, CellData{5, false}, CellData{6, false}, CellData{8, false}, CellData{2, false}},
		{CellData{2, false}, CellData{6, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{1, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
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
			if !IsNumberValid(boardData, rowIdx, colIdx, cellData.Number) {
				t.Errorf("\nInvalid number %v at (%v,%v)\n%v", cellData.Number, rowIdx, colIdx, PrintBoardData(boardData))
			}
		}
	}
}

func TestGenerateCellPositions(t *testing.T) {
	cellPositions := GenerateCellPositions()
	for i := 0; i < MaxNum; i++ {
		for j := 0; j < MaxNum; j++ {
			cellPos := cellPositions[i*MaxNum+j]
			if cellPos.RowIdx != i || cellPos.ColIdx != j {
				t.Errorf("Incorrect position (%v, %v); want (%v, %v)", cellPos.RowIdx, cellPos.ColIdx, i, j)
			}
		}
	}
}

func TestGenerateRandomPositions(t *testing.T) {
	cellPositions := GenerateCellPositions()
	shuffledPositions := ShuffleSlice(cellPositions)
	if len(cellPositions) != len(shuffledPositions) {
		t.Errorf("\nOriginal slice has a different length than the shuffled slice\nOriginal slice length: %v\nShuffled slice length: %v", len(cellPositions), len(shuffledPositions))
	}

	numMatchingPos := 0
	for _, cellPos := range cellPositions {
		hasFound := false
		for _, randCellPos := range shuffledPositions {
			if randCellPos.ColIdx == cellPos.ColIdx && randCellPos.RowIdx == cellPos.RowIdx {
				hasFound = true
				break
			}
		}
		if hasFound {
			numMatchingPos++
		}
	}

	if numMatchingPos != len(cellPositions) {
		t.Errorf("\nSome positions are either missing or mismatching\nOriginal slice: %v\nShuffled slice: %v", cellPositions, shuffledPositions)
	}
}

func TestCountSolutions(t *testing.T) {
	multipleSolutionBoardData := GenerateEmptyBoardData()
	multipleEmptyPos := []CellPos{{0, 0}, {0, 1}}
	numSolutions := CountSolutions(multipleSolutionBoardData, multipleEmptyPos)
	expectedNumSolutions := 72
	if numSolutions != expectedNumSolutions {
		t.Errorf("Expected %v solutions for board with only first row filled, got %v", expectedNumSolutions, numSolutions)
	}

	uniqueSolutionBoardData := BoardData{
		{CellData{5, false}, CellData{3, false}, CellData{0, false}, CellData{0, false}, CellData{7, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{6, false}, CellData{0, false}, CellData{0, false}, CellData{1, false}, CellData{9, false}, CellData{5, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}},
		{CellData{0, false}, CellData{9, false}, CellData{8, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{6, false}, CellData{0, false}},
		{CellData{8, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{6, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{3, false}},
		{CellData{4, false}, CellData{0, false}, CellData{0, false}, CellData{8, false}, CellData{0, false}, CellData{3, false}, CellData{0, false}, CellData{0, false}, CellData{1, false}},
		{CellData{7, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{2, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{6, false}},
		{CellData{0, false}, CellData{6, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{2, false}, CellData{8, false}, CellData{0, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{4, false}, CellData{1, false}, CellData{9, false}, CellData{0, false}, CellData{0, false}, CellData{5, false}},
		{CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{0, false}, CellData{8, false}, CellData{0, false}, CellData{0, false}, CellData{7, false}, CellData{9, false}},
	}

	var uniqueEmptyPos []CellPos
	for i := 0; i < MaxNum; i++ {
		for j := 0; j < MaxNum; j++ {
			if uniqueSolutionBoardData[i][j].Number == 0 {
				uniqueEmptyPos = append(uniqueEmptyPos, CellPos{i, j})
			}
		}
	}

	uniqueSolutions := CountSolutions(uniqueSolutionBoardData, uniqueEmptyPos)
	if uniqueSolutions != 1 {
		t.Errorf("Expected exactly one solution for valid Sudoku puzzle, got %v", uniqueSolutions)
	}
}

func TestGetNumEmptyCells(t *testing.T) {
	expectedNumEmptyCellsForEasy := 20
	expectedNumEmptyCellsForMedium := 30
	expectedNumEmptyCellsForHard := 40

	numEmptyCellsForEasy := GetNumEmptyCells(Easy)
	numEmptyCellsForMedium := GetNumEmptyCells(Medium)
	numEmptyCellsForHard := GetNumEmptyCells(Hard)

	if numEmptyCellsForEasy != expectedNumEmptyCellsForEasy {
		t.Errorf("\nNumber of empty cells for easy level is incorrect\nExpected output: %v\n Actual output: %v", expectedNumEmptyCellsForEasy, numEmptyCellsForEasy)
	}

	if numEmptyCellsForMedium != expectedNumEmptyCellsForMedium {
		t.Errorf("\nNumber of empty cells for medium level is incorrect\nExpected output: %v\n Actual output: %v", expectedNumEmptyCellsForMedium, numEmptyCellsForMedium)
	}

	if numEmptyCellsForHard != expectedNumEmptyCellsForHard {
		t.Errorf("\nNumber of empty cells for hard level is incorrect\nExpected output: %v\n Actual output: %v", expectedNumEmptyCellsForHard, numEmptyCellsForHard)
	}
}

func TestRemoveNumbers(t *testing.T) {
	boardData := DeepCopyBoardData(FilledBoardData)

	numEmptyCells := GetNumEmptyCells(Medium)

	RemoveNumbers(boardData, numEmptyCells)

	numZero := 0
	for i := 0; i < len(boardData); i++ {
		for j := 0; j < len(boardData[i]); j++ {
			if boardData[i][j].Number == 0 {
				numZero++
			}
		}
	}

	if numZero != int(numEmptyCells) {
		t.Errorf("Expected %d cells to be removed, got %d", numEmptyCells, numZero)
	}
}

func TestGenerateInitialBoardData(t *testing.T) {
	boardData := GenerateInitialBoardData(Easy)

	// Check board dimensions are correct
	if len(boardData) != MaxNum {
		t.Errorf("Board has incorrect number of rows. Expected %d, got %d", MaxNum, len(boardData))
	}
	for i := 0; i < len(boardData); i++ {
		if len(boardData[i]) != MaxNum {
			t.Errorf("Row %d has incorrect length. Expected %d, got %d", i, MaxNum, len(boardData[i]))
		}
	}

	// Count empty cells
	numEmptyCells := GetNumEmptyCells(Easy)
	numZero := 0
	for i := 0; i < len(boardData); i++ {
		for j := 0; j < len(boardData[i]); j++ {
			if boardData[i][j].Number == 0 {
				numZero++
			}
		}
	}

	if numZero != numEmptyCells {
		t.Errorf("Expected %d empty cells, got %d", numEmptyCells, numZero)
	}

	// Verify cell editability
	for i := 0; i < len(boardData); i++ {
		for j := 0; j < len(boardData[i]); j++ {
			// Zero cell
			if boardData[i][j].Number == 0 && !boardData[i][j].Editable {
				t.Errorf("Expected (%v,%v) to be editable", i, j)
			}

			// Non-zero cell
			if boardData[i][j].Number != 0 && boardData[i][j].Editable {
				t.Errorf("Expected (%v,%v) to be non-editable", i, j)
			}
		}
	}

	// Verify board has exactly one solution
	emptyPositions := make([]CellPos, 0)
	for i := 0; i < len(boardData); i++ {
		for j := 0; j < len(boardData[i]); j++ {
			if boardData[i][j].Number == 0 {
				emptyPositions = append(emptyPositions, CellPos{i, j})
			}
		}
	}

	numSolutions := CountSolutions(boardData, emptyPositions)
	if numSolutions != 1 {
		t.Errorf("Board should have exactly one solution, got %d solutions", numSolutions)
	}
}