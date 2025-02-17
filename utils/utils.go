package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/types"
	"github.com/rivo/tview"
)

/***************************/
/***** General Purpose *****/
/***************************/

func DeepCopyBoardData(boardData types.BoardData) types.BoardData {
	copiedBoardData := make(types.BoardData, len(boardData))
	for i := range boardData {
		row := make([]types.CellData, len(boardData[i]))
		copy(row, boardData[i])
		copiedBoardData[i] = row
	}
	return copiedBoardData
}

func PrintBoardData(boardData types.BoardData) string {
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

func IsBoardComplete(boardData types.BoardData) bool {
  for rowIdx, row := range boardData {
    for colIdx := range row {
      cellData := &row[colIdx]
      if cellData.Number == 0 || !IsNumberValid(boardData, rowIdx, colIdx, cellData.Number) {
        return false
      }
    }
  }

  return true
}

func FormatTime(duration time.Duration) string {
	hours := duration / time.Hour
	minutes := duration % time.Hour / time.Minute
	seconds := duration % time.Minute / time.Second
	milliseconds := duration % time.Second / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d:%03d", hours, minutes, seconds, milliseconds)
}

func StartTimer(timeTextView *tview.TextView, app *tview.Application) chan bool {
  startTime := time.Now()
  done := make(chan bool)
  go func() {
    for {
      select {
      case <-done:
        timeTextView.SetTextColor(tcell.ColorGreen)
        return
      default:
        duration := time.Since(startTime)
        timeStr := FormatTime(duration)
        app.QueueUpdateDraw(func() {
          timeTextView.SetText(timeStr)
        })
        time.Sleep(time.Millisecond)
      }
    }
  }()
  return done
}


/*********************************/
/***** Sudoku Initialization *****/
/*********************************/

const (
	Easy types.Difficulty = iota
	Medium
	Hard
)

const MaxNum = 9

var (
	cellsToRemoveByDifficulty = map[types.Difficulty]int{
		Easy:   20,
		Medium: 30,
		Hard:   40,
	}
)

func GetNumEmptyCells(difficulty types.Difficulty) int {
	numEmptyCells := cellsToRemoveByDifficulty[difficulty]
	if numEmptyCells == 0 {
		panic(fmt.Sprintf("Invalid difficulty value: %v", difficulty))
	}

	return numEmptyCells
}

func GeneratePossibleNumbers() [MaxNum]int {
	return [MaxNum]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
}

func IsValidRowForNumber(boardData types.BoardData, rowIdx, colIdx, number int) bool {
	row := boardData[rowIdx]

	for currColIdx, elem := range row {
		if colIdx != currColIdx && elem.Number == number {
			return false
		}
	}

	return true
}

func IsValidColForNumber(boardData types.BoardData, rowIdx, colIdx, number int) bool {
	for currRowIdx, elem := range boardData {
		if currRowIdx != rowIdx && elem[colIdx].Number == number {
			return false
		}
	}

	return true
}

func IsValidInnerGridForNumber(boardData types.BoardData, rowIdx, colIdx, number int) bool {
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

func IsNumberValid(boardData types.BoardData, emptyCellRowIdx, emptyCellColIdx, number int) bool {
	return (IsValidRowForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidColForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number) &&
		IsValidInnerGridForNumber(boardData, emptyCellRowIdx, emptyCellColIdx, number))
}

func GenerateEmptyBoardData() types.BoardData {
	boardData := make(types.BoardData, MaxNum)
	for rowIdx := range boardData {
		boardData[rowIdx] = make([]types.CellData, MaxNum)
		for colIdx := range boardData[rowIdx] {
			boardData[rowIdx][colIdx] = types.CellData{Number: 0, Editable: false}
		}
	}
	return boardData
}

func FindNextEmptyCellPos(boardData types.BoardData) *types.CellPos {
	for rowIdx := range boardData {
		rowData := boardData[rowIdx]
		for colIdx := range rowData {
			cellData := boardData[rowIdx][colIdx]
			if cellData.Number == 0 {
				return &types.CellPos{RowIdx: rowIdx, ColIdx: colIdx}
			}
		}
	}

	return nil
}

func FillBoardData(boardData types.BoardData) bool {
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

func GenerateFilledBoardData() types.BoardData {
	boardData := GenerateEmptyBoardData()
	FillBoardData(boardData)
	return boardData
}

func GenerateCellPositions() []types.CellPos {
	positions := make([]types.CellPos, MaxNum*MaxNum)
	for rowIdx := 0; rowIdx < MaxNum; rowIdx++ {
		for colIdx := 0; colIdx < MaxNum; colIdx++ {
			posIdx := rowIdx*MaxNum + colIdx
			positions[posIdx] = types.CellPos{RowIdx: rowIdx, ColIdx: colIdx}
		}
	}

	return positions
}

func countSolutionsHelper(boardData types.BoardData, emptyPosSlice []types.CellPos, idx int, count *int) {
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

func CountSolutions(boardData types.BoardData, emptyPosSlice []types.CellPos) int {
	count := 0
	countSolutionsHelper(boardData, emptyPosSlice, 0, &count)
	return count
}

func RemoveNumbers(boardData types.BoardData, numEmptyCells int) {
	cellPositions := GenerateCellPositions()
	shuffledPositions := ShuffleSlice(cellPositions)

	zeroPositions := make([]types.CellPos, 0)

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

func GenerateInitialBoardData(difficulty types.Difficulty) types.BoardData {
	numEmptyCells := GetNumEmptyCells(difficulty)
	boardData := GenerateFilledBoardData()
	RemoveNumbers(boardData, numEmptyCells)
	return boardData
}

/***************************/
/***** 			 UI 			 *****/
/***************************/

func CreateCenteredPrimitive[T tview.Primitive](primitive T, width, height int) *tview.Flex {
	wrappedPrimitive := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(primitive, width, 1, true).
		AddItem(nil, 0, 1, false).
		SetDirection(tview.FlexColumn)

	wrappedPrimitive = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(wrappedPrimitive, height, 1, true).
		AddItem(nil, 0, 1, false)
	return wrappedPrimitive
}

func RegisterInputCaptureHandlers[P types.InputCapturePrimitive](primitive P, handlers []types.InputCaptureHandler) {
	primitive.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, handler := range handlers {
			event, stop := handler(event)
			if stop {
				return event
			}
		}
		return event
	})
}