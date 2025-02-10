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

func createTimerTextView() *tview.TextView {
	timerTextView := tview.NewTextView().
		SetText("00:00:00:000").
		SetTextAlign(tview.AlignCenter)
	timerTextView.SetBorderPadding(0,1,0,0)
	return timerTextView
}

func createBoardWithTimer(board *tview.Table, timer *tview.TextView) *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
			timer,
			2,
			1,
			false,
		).
		AddItem(board, 19, 1, true)
}

func createBoardFrame(boardWithTimer *tview.Flex) *tview.Frame {
	centeredTable := utils.CreateCenteredPrimitive(boardWithTimer, 37, 21)
	return tview.NewFrame(centeredTable).
		AddText(
			"Press b to go back", 
			false, 
			tview.AlignCenter, 
			tcell.ColorReset,
		)
}

func createBoardPage(tableFrame *tview.Frame) *tview.Frame {
	centeredTableFrame := utils.CreateCenteredPrimitive(tableFrame, 37, 27)
	return tview.NewFrame(centeredTableFrame)
}

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

func registerTableInputCaptureHandlers(table *tview.Table, boardData types.BoardData, tableFrame *tview.Frame, done chan bool) {
	numberInputHandler := numberInputHandlerCurry(table, boardData, tableFrame, done)
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

/*
|----------BoardPage--------|
|	|-------BoardFrame------|	|
| |	|--BoardWithTimer---| | |
| | |       Timer       | | |
|	|	| |-----Board-----|	| | |
|	|	| |								|	| | |
|	|	| |								|	| | |
|	|	| |								|	| | |
|	|	| |								|	| | |
|	|	| |								|	| | |
|	|	| |---------------|	| | |
|	|	|-------------------| | |
|	|	         Msg          | |
|	|-----------------------|	|
|---------------------------|
*/
func GenerateBoard(boardData types.BoardData, app *tview.Application) *tview.Frame {
	timer := createTimerTextView()

	board := tview.NewTable()

	boardWithTimer := createBoardWithTimer(board, timer)

	boardFrame := createBoardFrame(boardWithTimer)

	boardPage := createBoardPage(boardFrame)

	done := utils.StartTimer(timer, app)

	customizeBoard(board)

	insertCells(board, boardData)

	focusFirstCell(board, boardData)

	registerBoardSelectionChangedHandlers(board, boardData)

	registerTableInputCaptureHandlers(board, boardData, boardFrame, done)

	return boardPage
}
