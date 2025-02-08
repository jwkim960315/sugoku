package button

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GenerateDifficultyButton(content string) *tview.Button {
	button := tview.NewButton(content).
		SetActivatedStyle(tcell.StyleDefault.
			Background(tcell.ColorRed).
			Foreground(tcell.ColorWhite).
			Bold(true),
		).
		SetStyle(tcell.StyleDefault.
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorBlack),
		)

	return button
}