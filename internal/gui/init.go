package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

var settings fyne.Preferences

type ui struct {
	app fyne.App

	manager pkg.Managerer
	pkgList struct {
		*widget.List
		pkgs []pkg.Packager
	}

	window fyne.Window
}

func InitGUI() {
	app := app.NewWithID("com.github.tom5521.gwpm")
	settings = app.Preferences()
	ui := &ui{
		app:    app,
		window: app.NewWindow("Graphic Windows Package Manager"),
	}
	// Init Utils.
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
	ui.pkgList.pkgs = ui.MakePkgList(ui.manager)
	ui.pkgList.List = ui.MakeList(ui.pkgList.pkgs)
	ui.pkgList.OnSelected = func(id widget.ListItemID) {
		ui.lateral = InitLateral(ui.pkgList.pkgs[id])
		ui.lateral.Show()
	}
}

func (ui *ui) InitContent() {
	ui.lateral.Form = widget.NewForm()
	content := container.NewBorder(nil, nil, nil, ui.lateral, ui.pkgList)
	ui.window.SetContent(content)
}
