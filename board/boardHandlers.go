package board

import (
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