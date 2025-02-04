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
