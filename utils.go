package main

import (
	"math/rand"
)

func GeneratePossibleNumbers() [9]uint {
	return [9]uint{1,2,3,4,5,6,7,8,9}
}

func GenerateRandomRowNumbers(rowSlice []uint) []uint {
	rand.Shuffle(len(rowSlice), func(idx1, idx2 int) {
		(rowSlice)[idx1], rowSlice[idx2] = rowSlice[idx2], rowSlice[idx1]
	});
  
  return rowSlice
}

func IsValidRowForNumber(boardData BoardData, rowIdx uint, number uint) bool {
  row := boardData[rowIdx]

  for _, elem := range row {
    if elem.Number == number {
      return false
    }
  }

  return true
}

func IsValidColForNumber(boardData BoardData, colIdx uint, number uint) bool {
  for _, elem := range boardData {
    if elem[colIdx].Number == number {
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
      if boardData[row][col].Number == number {
        return false
      }
    }
  }

  return true
}

func IsNumberValid(boardData BoardData, emptyCellRowIdx, emptyCellColIdx, number uint) bool {
  return (
    IsValidRowForNumber(boardData, emptyCellRowIdx, number) &&
    IsValidColForNumber(boardData, emptyCellColIdx, number) &&
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