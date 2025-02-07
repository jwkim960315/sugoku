package board

import "github.com/gdamore/tcell/v2"

type CellSelectionChangedHandler = func(row, col int)

type InputCaptureHandler = func(event *tcell.EventKey) (*tcell.EventKey, bool)
