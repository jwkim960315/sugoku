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

func GenerateBoard(boardData types.BoardData, app *tview.Application) *tview.Frame {
	table := tview.NewTable()

	tableFrame := tview.NewFrame(utils.GetCenteredComponent(table, 37, 19)).
		AddText("Press b to go back", false, tview.AlignCenter, tcell.ColorReset)
	tablePage := tview.NewFrame(utils.GetCenteredComponent(tableFrame, 37, 25))

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
	numberInputHandler := numberInputHandlerCurry(table, boardData, tableFrame)
	deleteCellNumberHandler := deleteCellNumberHandlerCurry(table, boardData)

	utils.RegisterInputCaptureHandlers(
		table,
		[]types.InputCaptureHandler{
			appQuitHandler,
			numberInputHandler,
			deleteCellNumberHandler,
		},
	)

	firstCellTextColor := cell.GetCellTextColor(&boardData[0][0], true)
	table.GetCell(0, 0).SetTextColor(firstCellTextColor)

	return tablePage
}
