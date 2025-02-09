package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jwkim960315/sugoku/landing"
	"github.com/rivo/tview"
)

func SetupApp(app *tview.Application, pageIdx *int) *tview.Application {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
			return event
		}
		
		if *pageIdx == 1 && event.Rune() == 'b' {
			landingScreenPage := landing.CreateLandingPage(app, pageIdx)
			app.SetRoot(landingScreenPage, true)
			*pageIdx--
		}
		return event
	})

	return app
}