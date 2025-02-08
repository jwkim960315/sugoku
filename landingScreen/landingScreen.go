package landingScreen

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/types"
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