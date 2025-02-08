package board

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/cell"
	"github.com/jwkim960315/sugoku/types"
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
			textColor := cell.GetCellTextColor(&boardData[prevRow][prevCol], false)
			cellComp.SetTextColor(textColor)
		}

		// Set the current selected cell's border color to red.
		if cellComp := table.GetCell(row, col); cellComp != nil {
			textColor := cell.GetCellTextColor(&boardData[row][col], true)
			cellComp.SetTextColor(textColor)
		}

		prevRow, prevCol = row, col
	}
}

/*************************/
/***** Input Capture *****/
/*************************/

func appQuitHandlerCurry(app *tview.Application) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		if event.Rune() == 'q' {
			app.Stop()
			return event, true
		}

		return nil, false
	}
}

func numberInputHandlerCurry(table *tview.Table, boardData types.BoardData) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		switch event.Rune() {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			row, col := table.GetSelection()
			if !boardData[row][col].Editable {
				return event, true
			}

			selectedCell := table.GetCell(row, col)
			cellContent := cell.GenerateCellContent(int(event.Rune() - '0'))
			textColor := cell.GetCellTextColor(&boardData[row][col], true)
			selectedCell.SetText(cellContent).SetTextColor(textColor)
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
			return event, true
		}

		return nil, false
	}
}
