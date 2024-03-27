package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

var settings fyne.Preferences

type ui struct {
	app fyne.App

	manager pkg.Managerer
	pkgList struct {
		list *widget.List
		pkgs []pkg.Packager
	}

	lateral Lateral

	window fyne.Window
}

func InitGUI() {
	app := app.NewWithID("com.github.tom5521.gwpm")
	settings = app.Preferences()
	ui := &ui{
		app:    app,
		window: app.NewWindow("Graphic Windows Package Manager"),
	}
	ui.window.Resize(fyne.NewSize(828, 390))
	// Init Utils.
	ui.InitManager()
	ui.InitList()
	ui.lateral.Init()
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
	ui.pkgList.pkgs = ui.MakePkgSlice(ui.manager)
	ui.pkgList.list = ui.MakeList(ui.pkgList.pkgs)
	ui.pkgList.list.OnSelected = func(id widget.ListItemID) {
		ui.lateral.Load(ui.pkgList.pkgs[id])
	}
}

func (ui *ui) InitContent() {
	lateralBox := boxes.NewBorder(nil, nil, widget.NewSeparator(), nil, ui.lateral)
	content := boxes.NewBorder(nil, nil, nil, lateralBox, ui.pkgList.list)
	ui.window.SetContent(content)
}
