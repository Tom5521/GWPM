package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

var settings fyne.Preferences

type ui struct {
	app fyne.App

	manager pkg.Managerer
	pkgList *widget.List

	window fyne.Window

	lateralMenu        *fyne.Container
	lateralMenuContent struct {
	}
}

func InitGUI() {
	app := app.NewWithID("com.github.tom5521.gwpm")
	settings = app.Preferences()
	ui := &ui{
		app:    app,
		window: app.NewWindow("Graphic Windows Package Manager"),
	}
	ui.InitManager()
	ui.InitList()
	ui.InitContent()

	ui.window.ShowAndRun()
}

func (ui *ui) InitManager() {
	m, err := ui.MakeManager()
	if err != nil {
		popups.FatalError(err)
	}
	ui.manager = m
}

func (ui *ui) InitList() {
	l, err := ui.MakeList(ui.manager)
	if err != nil {
		popups.FatalError(err)
	}
	ui.pkgList = l
}

func (ui *ui) InitContent() {
	ui.window.SetContent(ui.pkgList)
}
