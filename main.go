package main

import (
	"github.com/jwkim960315/sugoku/landingScreen"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	pageIdx := 0
	
	SetupApp(app, &pageIdx)

	landingScreenPage := landingScreen.GenerateLandingScreen(app, &pageIdx)
	
	if err := app.SetRoot(landingScreenPage, true).Run(); err != nil {
		panic(err)
	}
}
