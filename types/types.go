package types

import "github.com/gdamore/tcell/v2"

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
