package board

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/types"
	"github.com/jwkim960315/sugoku/utils"
	"github.com/rivo/tview"
)

/******************/
/* 			Cell 			*/
/******************/

func getCellTextColor(cellData *types.CellData, selected bool) tcell.Color {
	if selected {
		return tcell.ColorRed
	}

	if cellData.Editable {
		return tcell.ColorReset
	}

	return tcell.ColorGreen
}

func createCellContent(number int) string {
	return " " + strconv.Itoa(number) + " "
}

func createCell(cellData *types.CellData) *tview.TableCell {
	str := createCellContent(cellData.Number)
	return tview.NewTableCell(str).SetAlign(tview.AlignCenter)
}

/*******************/
/*		  Board 		 */
/*******************/

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
			textColor := getCellTextColor(&boardData[row][col], false)
			tableCell := createCell(&boardData[row][col]).SetTextColor(textColor)
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

func registerBoardSelectionChangedHandlers(table *tview.Table, boardData types.BoardData) {
	cellSelectionChangedHandlers := []CellSelectionChangedHandler{
		updateSelectedCellCurry(table, boardData),
	}

	registerCellSelectionChangedHandlers(
		table,
		cellSelectionChangedHandlers,
	)
}

func registerTableInputCaptureHandlers(table *tview.Table, boardData types.BoardData, tableFrame *tview.Frame) {
	numberInputHandler := numberInputHandlerCurry(table, boardData, tableFrame)
	deleteCellNumberHandler := deleteCellNumberHandlerCurry(table, boardData)

	utils.RegisterInputCaptureHandlers(
		table,
		[]types.InputCaptureHandler{
			numberInputHandler,
			deleteCellNumberHandler,
		},
	)
}

func focusFirstCell(table *tview.Table, boardData types.BoardData) {
	firstCellTextColor := getCellTextColor(&boardData[0][0], true)
	table.GetCell(0, 0).SetTextColor(firstCellTextColor)
}

func GenerateBoard(boardData types.BoardData, app *tview.Application) *tview.Frame {
	table := tview.NewTable()

	tableFrame := tview.NewFrame(utils.GetCenteredComponent(table, 37, 19)).
		AddText("Press b to go back", false, tview.AlignCenter, tcell.ColorReset)
	tablePage := tview.NewFrame(utils.GetCenteredComponent(tableFrame, 37, 25))

	customizeBoard(table)

	insertCells(table, boardData)

	registerBoardSelectionChangedHandlers(table, boardData)

	registerTableInputCaptureHandlers(table, boardData, tableFrame)

	focusFirstCell(table, boardData)

	return tablePage
}
