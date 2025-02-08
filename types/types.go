package types

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CellData struct {
	Number   int
	Editable bool
}

type BoardData = [][]CellData

type CellPos struct {
	RowIdx int
	ColIdx int
}

type Difficulty int

type InputCaptureHandler = func(event *tcell.EventKey) (*tcell.EventKey, bool)

type InputCapturePrimitive interface {
	SetInputCapture(func(*tcell.EventKey) *tcell.EventKey) *tview.Box
}
