package board

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func customizeBoard(table *tview.Table) *tview.Table {
	return table.SetBorders(true).
				 SetBordersColor(tcell.ColorReset).
				 SetSelectable(true, true).
				 Select(0,0).
				 SetSelectedStyle(
				 	tcell.StyleDefault.
								Foreground(tcell.ColorNone).
				 				Background(tcell.ColorNone),
				 )
}