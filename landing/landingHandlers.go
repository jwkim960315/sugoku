package landing

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/board"
	"github.com/jwkim960315/sugoku/types"
	"github.com/jwkim960315/sugoku/utils"
	"github.com/rivo/tview"
)

/*************************/
/***** Input Capture *****/
/*************************/

func navigateButtonsHandlerCurry(buttonsContainer *tview.Flex, selectedItemIdx *int, app *tview.Application) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		itemCount := buttonsContainer.GetItemCount()
		switch event.Key() {
		case tcell.KeyUp:
			if *selectedItemIdx == 0 {
				*selectedItemIdx = itemCount - 1
			} else {
				*selectedItemIdx -= 2
			}

			selectedItem := buttonsContainer.GetItem(*selectedItemIdx)
			app.SetFocus(selectedItem)
			return event, true
		case tcell.KeyDown:
			if *selectedItemIdx == itemCount - 1 {
				*selectedItemIdx = 0
			} else {
				*selectedItemIdx += 2
			}
			
			selectedItem := buttonsContainer.GetItem(*selectedItemIdx)
			app.SetFocus(selectedItem)
			return event, true
		default:
			return nil, false
		}
	}
}

func chooseDifficultyButtonHandlerCurry(app *tview.Application, selectedItemIdx *int, pageIdx *int) types.InputCaptureHandler {
	return func(event *tcell.EventKey) (*tcell.EventKey, bool) {
		if event.Key() == tcell.KeyEnter {
			difficulty := utils.Easy
			switch *selectedItemIdx {
			case 2:
				difficulty = utils.Medium
			case 4:
				difficulty = utils.Hard
			}
			boardData := utils.GenerateInitialBoardData(difficulty)
			tablePage := board.GenerateBoard(boardData, app)
			app.SetRoot(tablePage, true)
			*pageIdx++
			return event, true
		}
		return nil, false
	}
}