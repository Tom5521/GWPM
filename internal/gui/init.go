package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type ui struct {
	app        fyne.App
	mainWindow fyne.Window
}

func InitGUI() {
	app := app.New()
	ui := &ui{
		app:        app,
		mainWindow: app.NewWindow("Graphic Windows Package Manager"),
	}

	ui.mainWindow.ShowAndRun()
}
