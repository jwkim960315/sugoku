package board

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/types"
	"github.com/jwkim960315/sugoku/utils"
	"github.com/rivo/tview"
)

/*************************/
/*** Selection Changed ***/
/*************************/

func updateSelectedCellCurry(table *tview.Table, boardData types.BoardData) func(row, col int) {
	// Keep track of the previously-selected position
	prevRow, prevCol := 0, 0

	return func(row, col int) {
		// Reset the previously selected cell's border color.
		if cellComp := table.GetCell(prevRow, prevCol); cellComp != nil {
			textColor := getCellTextColor(&boardData[prevRow][prevCol], false)
			cellComp.SetTextColor(textColor).SetAttributes(tcell.AttrNone)
		}

		// Set the current selected cell's border color to red.
		if cellComp := table.GetCell(row, col); cellComp != nil {
			textColor := getCellTextColor(&boardData[row][col], true)
			cellComp.SetTextColor(textColor).SetAttributes(tcell.AttrBold)
		}

		prevRow, prevCol = row, col
	}
}

/*************************/
/***** Input Capture *****/
/*************************/

func numberInputHandlerCurry(table *tview.Table, boardData types.BoardData, tablePage *tview.Frame, done chan bool) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		switch event.Rune() {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			row, col := table.GetSelection()
			if !boardData[row][col].Editable {
				return event, true
			}

			selectedCell := table.GetCell(row, col)
			cellContent := createCellContent(int(event.Rune() - '0'))
			textColor := getCellTextColor(&boardData[row][col], true)
			selectedCell.SetText(cellContent).SetTextColor(textColor)

			boardData[row][col].Number = int(event.Rune() - '0')
			
			if utils.IsBoardComplete(boardData) {
				done <- true
				disableAllCells(boardData)
				tablePage.AddText("", false, tview.AlignCenter, tcell.ColorReset).
					AddText("Puzzle complete! 🎉", false, tview.AlignCenter, tcell.ColorGreen)
			}

			return event, true
		}

		return nil, false
	}
}

func deleteCellNumberHandlerCurry(table *tview.Table, boardData types.BoardData) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		if event.Key() == tcell.KeyBackspace || event.Key() == tcell.KeyBackspace2 {
			row, col := table.GetSelection()
			if !boardData[row][col].Editable {
				return event, true
			}
			table.GetCell(row, col).SetText(" 0 ")

			boardData[row][col].Number = 0

			return event, true
		}

		return nil, false
	}
}

/*************************/
/*****     Misc.     *****/
/*************************/

func disableAllCells(boardData types.BoardData) {
	for _, row := range boardData {
		for colIdx := range row {
			row[colIdx].Editable = false
		}
	}
}