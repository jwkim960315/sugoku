package main

import (
	"fmt"
	"math/rand"
)

/***************************/
/***** General Purpose *****/
/***************************/

func DeepCopyBoardData(boardData BoardData) BoardData {
  copiedBoardData := make(BoardData, len(boardData))
  for i := range boardData {
    row := make([]CellData, len(boardData[i]))
    copy(row, boardData[i])
    copiedBoardData[i] = row
  }
  return copiedBoardData
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

func ShuffleSlice[T any](slice []T) []T {
  rand.Shuffle(len(slice), func(idx1, idx2 int) {
		(slice)[idx1], slice[idx2] = slice[idx2], slice[idx1]
	})

	return slice
}

/*********************************/
/***** Sudoku Initialization *****/
/*********************************/

const (
  Easy Difficulty = iota
  Medium
  Hard
)

const MaxNum = 9

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

func GeneratePossibleNumbers() [MaxNum]int {
	return [MaxNum]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func IsValidRowForNumber(boardData BoardData, rowIdx int, colIdx int, number int) bool {
	row := boardData[rowIdx]

	for currColIdx, elem := range row {
		if colIdx != currColIdx && elem.Number == number {
			return false
		}
	}

	return true
}

func IsValidColForNumber(boardData BoardData, rowIdx int, colIdx int, number int) bool {
	for currRowIdx, elem := range boardData {
		if currRowIdx != rowIdx && elem[colIdx].Number == number {
			return false
		}
	}

	return true
}

func IsValidInnerGridForNumber(boardData BoardData, rowIdx int, colIdx int, number int) bool {
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

func IsNumberValid(boardData BoardData, emptyCellRowIdx, emptyCellColIdx, number int) bool {
	return (IsValidRowForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidColForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidInnerGridForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number))
}

func GenerateEmptyBoardData() BoardData {
	boardData := make(BoardData, MaxNum)
	for rowIdx := range boardData {
		boardData[rowIdx] = make([]CellData, MaxNum)
		for colIdx := range boardData[rowIdx] {
			boardData[rowIdx][colIdx] = CellData{0, false}
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
				return &CellPos{rowIdx, colIdx}
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
	shuffledRowData := ShuffleSlice(rowData[:])

	for _, number := range shuffledRowData {
		emptyCellRowIdx := emptyCellPos.RowIdx
		emptyCellColIdx := emptyCellPos.ColIdx
		emptyCellData := &boardData[emptyCellRowIdx][emptyCellColIdx]
		if IsNumberValid(boardData, emptyCellRowIdx, emptyCellColIdx, number) {
			emptyCellData.Number = number
			if FillBoardData(boardData) {
				return true
			}
			boardData[emptyCellRowIdx][emptyCellColIdx].Number = 0
		}
	}

	return false
}

func GenerateFilledBoardData() BoardData {
	boardData := GenerateEmptyBoardData()
	FillBoardData(boardData)
	return boardData
}

func GenerateCellPositions() []CellPos {
  positions := make([]CellPos, MaxNum*MaxNum)
  for rowIdx := 0; rowIdx < MaxNum; rowIdx++ {
    for colIdx := 0; colIdx < MaxNum; colIdx++ {
      posIdx := rowIdx * MaxNum + colIdx
      positions[posIdx] = CellPos{rowIdx, colIdx}
    }
  }

  return positions
}

func countSolutionsHelper(boardData BoardData, emptyPosSlice []CellPos, idx int, count *int) {
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
    cellData.Number = 0
  }
}

func CountSolutions(boardData BoardData, emptyPosSlice []CellPos) int {
  count := 0
  countSolutionsHelper(boardData, emptyPosSlice, 0, &count)
  return count
}

func RemoveNumbers(boardData BoardData, numEmptyCells int) {
  cellPositions := GenerateCellPositions()
  shuffledPositions := ShuffleSlice(cellPositions)

  zeroPositions := make([]CellPos, 0);

  for idx := 0; len(zeroPositions) < int(numEmptyCells) && idx < len(shuffledPositions); idx++ {
    cellPos := &shuffledPositions[idx]
    cellNumber := boardData[cellPos.RowIdx][cellPos.ColIdx].Number
    boardData[cellPos.RowIdx][cellPos.ColIdx].Number = 0
    boardData[cellPos.RowIdx][cellPos.ColIdx].Editable = true
    zeroPositions = append(zeroPositions, *cellPos)
    numSolutions := CountSolutions(boardData, zeroPositions)
    if numSolutions > 1 {
      boardData[cellPos.RowIdx][cellPos.ColIdx].Number = cellNumber
      boardData[cellPos.RowIdx][cellPos.ColIdx].Editable = false
      zeroPositions = zeroPositions[:len(zeroPositions)-1]
    }
  }
}

func GenerateInitialBoardData(difficulty Difficulty) BoardData {
  numEmptyCells := GetNumEmptyCells(difficulty)
  boardData := GenerateFilledBoardData()
  RemoveNumbers(boardData, numEmptyCells)
  return boardData
}