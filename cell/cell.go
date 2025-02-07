package cell

import (
	"strconv"

	"github.com/jwkim960315/sugoku/types"
	"github.com/rivo/tview"
)

func GenerateCellContent(number int) string {
	return " " + strconv.Itoa(number) + " "
}

func GenerateCell(cellData *types.CellData) *tview.TableCell {
  str := GenerateCellContent(cellData.Number)
  return tview.NewTableCell(str).SetAlign(tview.AlignCenter)
}