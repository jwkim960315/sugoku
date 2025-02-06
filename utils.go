package main

import (
	"fmt"
	"math/rand"
)

const (
  Easy Difficulty = iota
  Medium
  Hard
)

var (
  cellsToRemoveByDifficulty = map[Difficulty]int{
    Easy: 20,
    Medium: 30,
    Hard: 40,
  }
)

func GetNumEmptyCells(difficulty Difficulty) int {
  numEmptyCells := cellsToRemoveByDifficulty[difficulty]
  if numEmptyCells == 0 {
    panic(fmt.Sprintf("Invalid difficulty value: %v", difficulty))
  }

  return numEmptyCells
}

func GeneratePossibleNumbers() [9]uint {
	return [9]uint{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func GenerateRandomRowNumbers(rowSlice []uint) []uint {
	rand.Shuffle(len(rowSlice), func(idx1, idx2 int) {
		(rowSlice)[idx1], rowSlice[idx2] = rowSlice[idx2], rowSlice[idx1]
	})

	return rowSlice
}

func IsValidRowForNumber(boardData BoardData, rowIdx uint, colIdx uint, number uint) bool {
	row := boardData[rowIdx]

	for currColIdx, elem := range row {
		if colIdx != uint(currColIdx) && elem.Number == number {
			return false
		}
	}

	return true
}

func IsValidColForNumber(boardData BoardData, rowIdx uint, colIdx uint, number uint) bool {
	for currRowIdx, elem := range boardData {
		if uint(currRowIdx) != rowIdx && elem[colIdx].Number == number {
			return false
		}
	}

	return true
}

func IsValidInnerGridForNumber(boardData BoardData, rowIdx uint, colIdx uint, number uint) bool {
	rowStartIdx := 3 * (rowIdx / 3)
	colStartIdx := 3 * (colIdx / 3)

	for row := rowStartIdx; row < rowStartIdx+3; row++ {
		for col := colStartIdx; col < colStartIdx+3; col++ {
			if (row != rowIdx || col != colIdx) && boardData[row][col].Number == number {
				return false
			}
		}
	}

	return true
}

func IsNumberValid(boardData BoardData, emptyCellRowIdx, emptyCellColIdx, number uint) bool {
	return (IsValidRowForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidColForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidInnerGridForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number))
}

func GenerateEmptyBoardData() BoardData {
	boardData := make(BoardData, 9)
	for rowIdx := range boardData {
		boardData[rowIdx] = make([]CellData, 9)
		for colIdx := range boardData[rowIdx] {
			boardData[rowIdx][colIdx] = CellData{0}
		}
	}
	return boardData
}

func FindNextEmptyCellPos(boardData BoardData) *CellPos {
	for rowIdx := range boardData {
		rowData := boardData[rowIdx]
		for colIdx := range rowData {
			cellData := boardData[rowIdx][colIdx]
			if cellData.Number == 0 {
				return &CellPos{uint(rowIdx), uint(colIdx)}
			}
		}
	}

	return nil
}

func FillBoardData(boardData BoardData) bool {
	emptyCellPos := FindNextEmptyCellPos(boardData)

	if emptyCellPos == nil {
		return true
	}

	rowData := GeneratePossibleNumbers()
	shuffledRowData := GenerateRandomRowNumbers(rowData[:])

	for _, number := range shuffledRowData {
		emptyCellRowIdx := emptyCellPos.RowIdx
		emptyCellColIdx := emptyCellPos.ColIdx
		emptyCellData := &boardData[emptyCellRowIdx][emptyCellColIdx]
		if IsNumberValid(boardData, emptyCellRowIdx, emptyCellColIdx, number) {
			emptyCellData.Number = number
			if FillBoardData(boardData) {
				return true
			}
			boardData[emptyCellRowIdx][emptyCellColIdx].Number = uint(0)
		}
	}

	return false
}

func GenerateFilledBoardData() BoardData {
	boardData := GenerateEmptyBoardData()
	FillBoardData(boardData)
	return boardData
}

func PrintBoardData(boardData BoardData) string {
	printStr := "-------------------------\n"

	for rowIdx, row := range boardData {
		printStr += "|"
		for colIdx := range row {
			cellData := &row[colIdx]
			printStr += " "
			printStr += fmt.Sprintf("%v", cellData.Number)
			if (colIdx+1)%3 == 0 {
				printStr += " |"
			}
		}
		printStr += "\n"

		if (rowIdx+1)%3 == 0 {
			printStr += "-------------------------\n"
		}
	}

	return printStr
}

func GenerateCellPositions() []CellPos {
  positions := make([]CellPos, 81)
  for rowIdx := 0; rowIdx < 9; rowIdx++ {
    for colIdx := 0; colIdx < 9; colIdx++ {
      posIdx := rowIdx * 9 + colIdx
      positions[posIdx] = CellPos{uint(rowIdx), uint(colIdx)}
    }
  }

  return positions
}

func GenerateRandomPositions(positions []CellPos) []CellPos {
  rand.Shuffle(len(positions), func(idx1, idx2 int) {
    (positions)[idx1], positions[idx2] = positions[idx2], positions[idx1]
  })

  return positions
}

func countSolutionsHelper(boardData BoardData, emptyPosSlice []CellPos, idx int, count *uint) {
  if idx == len(emptyPosSlice) {
    *count++
    return
  }

  cellPos := &emptyPosSlice[idx]
  possibleNumbers := GeneratePossibleNumbers()

  for _, num := range possibleNumbers {
    cellData := &boardData[cellPos.RowIdx][cellPos.ColIdx]
    cellData.Number = num
    if IsNumberValid(boardData, cellPos.RowIdx, cellPos.ColIdx, num) {
      countSolutionsHelper(boardData, emptyPosSlice, idx+1, count)
    }
    cellData.Number = uint(0)
  }
}

func CountSolutions(boardData BoardData, emptyPosSlice []CellPos) uint {
  count := uint(0)
  countSolutionsHelper(boardData, emptyPosSlice, 0, &count)
  return count
}
