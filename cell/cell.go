package cell

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/types"
	"github.com/rivo/tview"
)

func GetCellTextColor(cellData *types.CellData, selected bool) tcell.Color {
	if selected {
		return tcell.ColorRed
	}

	if cellData.Number == 0 {
		return tcell.ColorReset
	}

	return tcell.ColorGreen
}

func GenerateCellContent(number int) string {
	return " " + strconv.Itoa(number) + " "
}

func GenerateCell(cellData *types.CellData) *tview.TableCell {
  str := GenerateCellContent(cellData.Number)
  return tview.NewTableCell(str).SetAlign(tview.AlignCenter)
}