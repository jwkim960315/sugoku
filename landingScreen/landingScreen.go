package landingScreen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/button"
	"github.com/jwkim960315/sugoku/types"
	"github.com/jwkim960315/sugoku/utils"
	"github.com/rivo/tview"
)

func registerInputCaptureHandlers(flex *tview.Flex, handlers []types.InputCaptureHandler) {
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for _, handler := range handlers {
			event, stop := handler(event)
			if stop {
				return event
			}
		}
		return event
	})
}

func getDifficultyButtonContainer(buttons []*tview.Button) *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	focus := true
	for i := 0; i < len(buttons) * 2 - 1; i++ {
		if i % 2 == 0 {
			flex.AddItem(buttons[i / 2], 3, 1, focus)
			focus = false
			continue
		}

		flex.AddItem(nil, 1, 0, false)
	}

	return flex
}

func centerButtonsHorizontally(buttonContainer *tview.Flex) *tview.Flex {
  return tview.NewFlex().
    SetDirection(tview.FlexColumn).
    AddItem(nil, 0, 1, false).
    AddItem(buttonContainer, 30, 1, true).
    AddItem(nil, 0, 1, false)
}

func GenerateLandingScreen(app *tview.Application, pageIdx *int) *tview.Frame {
  easyButton := button.GenerateDifficultyButton("Easy")
  mediumButton := button.GenerateDifficultyButton("Medium")
  hardButton := button.GenerateDifficultyButton("Hard")
  
  buttonContainer := getDifficultyButtonContainer([]*tview.Button{easyButton, mediumButton, hardButton})

  selectedItemIdx := 0
  navigateButtonsHandler := navigateButtonsHandlerCurry(buttonContainer, &selectedItemIdx, app)
  chooseDifficultyButtonHandler := chooseDifficultyButtonHandlerCurry(app, &selectedItemIdx, pageIdx)
  inputCaptureHandlers := []types.InputCaptureHandler{navigateButtonsHandler, chooseDifficultyButtonHandler}

  registerInputCaptureHandlers(buttonContainer, inputCaptureHandlers)

  wrappedFlexBox := centerButtonsHorizontally(buttonContainer)

  // Need this for the background color
  frame := tview.NewFrame(wrappedFlexBox)

	landingScreenPage := tview.NewFrame(frame).
		AddText("Sudoku Puzzle", true, tview.AlignCenter, tcell.ColorReset).
		SetBorders(2, 2, 0, 0, 2, 2).
		AddText("Select a difficulty level to begin your Sudoku challenge", true, tview.AlignCenter, tcell.ColorReset)

	wrappedFrame := utils.GetCenteredComponent(landingScreenPage, 60, 19)

	landingScreenPage = tview.NewFrame(wrappedFrame)

  return landingScreenPage
}
