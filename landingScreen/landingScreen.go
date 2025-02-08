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