package board

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/cell"
	"github.com/jwkim960315/sugoku/types"
	"github.com/jwkim960315/sugoku/utils"
	"github.com/rivo/tview"
)

func customizeBoard(table *tview.Table) *tview.Table {
	return table.SetBorders(true).
		SetBordersColor(tcell.ColorReset).
		SetSelectable(true, true).
		Select(0, 0).
		SetSelectedStyle(
			tcell.StyleDefault.
				Foreground(tcell.ColorNone).
				Background(tcell.ColorNone),
		)
}

func insertCells(table *tview.Table, boardData types.BoardData) {
	for row := 0; row < utils.MaxNum; row++ {
		for col := 0; col < utils.MaxNum; col++ {
			textColor := cell.GetCellTextColor(&boardData[row][col], false)
			tableCell := cell.GenerateCell(&boardData[row][col]).SetTextColor(textColor)
			table.SetCell(row, col, tableCell)
		}
	}
}

func registerCellSelectionChangedHandlers(
	table *tview.Table,
	handlers []CellSelectionChangedHandler,
) *tview.Table {
	table.SetSelectionChangedFunc(func(row, col int) {
		for _, handler := range handlers {
			handler(row, col)
		}
	})

	return table
}

func registerInputCaptureHandlers(table *tview.Table, handlers []types.InputCaptureHandler) {
	table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, handler := range handlers {
			event, stop := handler(event)
			if stop {
				return event
			}
		}
		return event
	})
}

func GenerateBoard(boardData types.BoardData, app *tview.Application) *tview.Table {
	table := tview.NewTable()

	customizeBoard(table)

	insertCells(table, boardData)

	cellSelectionChangedHandlers := []CellSelectionChangedHandler{
		updateSelectedCellCurry(table, boardData),
	}

	registerCellSelectionChangedHandlers(
		table,
		cellSelectionChangedHandlers,
	)

	appQuitHandler := appQuitHandlerCurry(app)
	numberInputHandler := numberInputHandlerCurry(table, boardData)
	deleteCellNumberHandler := deleteCellNumberHandlerCurry(table, boardData)

	registerInputCaptureHandlers(
		table,
		[]types.InputCaptureHandler{
			appQuitHandler,
			numberInputHandler,
			deleteCellNumberHandler,
		},
	)

	firstCellTextColor := cell.GetCellTextColor(&boardData[0][0], true)
	table.GetCell(0, 0).SetTextColor(firstCellTextColor)

	return table
}
