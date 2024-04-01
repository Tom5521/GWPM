package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/choco"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
	"github.com/Tom5521/GWPM/pkg/scoop"

	boxes "fyne.io/fyne/v2/container"
)

type packager struct {
	pkg.Packager
	Checked bool
}

type ui struct {
	app      fyne.App
	settings fyne.Preferences

	mainMenu   *MainMenu
	mainWindow fyne.Window
	mainBox    *fyne.Container

	manager pkg.Managerer

	search  *Search
	sideBar *SideBar

	packages []packager
	list     *widget.List
}

var cui *ui

func InitGUI() {
	app := app.NewWithID("com.github.tom5521.gwpm")
	cui = &ui{
		app:      app,
		settings: app.Preferences(),

		// Initialize structures.
		sideBar:  new(SideBar),
		search:   new(Search),
		mainMenu: new(MainMenu),
	}
	cui.mainWindow = app.NewWindow("Graphic Windows Package Manager")
	cui.mainWindow.SetMaster()
	cui.mainWindow.Resize(fyne.NewSize(830, 390))

	// Initialize methods.
	cui.InitManager()
	cui.sideBar.Init()
	cui.search.Init()
	cui.search.InitSelect()
	cui.InitList()
	cui.InitBoxes()
	cui.mainMenu.Init()
	cui.mainWindow.SetMainMenu(cui.mainMenu.Menu)

	cui.mainWindow.SetContent(cui.mainBox)
	cui.mainWindow.ShowAndRun()
}

func (ui *ui) InitManager() {
	switch ui.settings.String("manager") {
	case choco.ManagerName:
		ui.manager = choco.Connect()
	case scoop.ManagerName:
		ui.manager = scoop.Connect()
	default:
		ui.settings.SetString("manager", choco.ManagerName)
		ui.manager = choco.Connect()
	}
}

func (ui *ui) InitPkgSlice() {
	var (
		packagers []pkg.Packager
		err       error
	)
	switch ui.settings.String("list-mode") {
	case "local":
		packagers, err = ui.manager.LocalPkgs()
	case "repo":
		if ui.search.Entry.Text == "" {
			packagers = []pkg.Packager{}
			return
		}
		packagers, err = ui.manager.SearchInRepo(ui.search.Entry.Text)
	default:
		ui.settings.SetString("list-mode", "local")
		packagers, err = ui.manager.LocalPkgs()
	}
	if err != nil {
		popups.FatalError(err)
	}
	ui.packages = []packager{}
	for _, p := range packagers {
		ui.packages = append(ui.packages, packager{
			Packager: p,
		})
	}
}

func (ui *ui) InitList() {
	ui.list = widget.NewList(
		func() int { return len(ui.packages) },
		func() fyne.CanvasObject {
			return boxes.NewBorder(nil, nil, nil, &widget.Check{}, &widget.Label{})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := co.(*fyne.Container)
			label := c.Objects[0].(*widget.Label)
			label.SetText(ui.packages[lii].Name())

			check := c.Objects[1].(*widget.Check)
			check.OnChanged = func(b bool) {
				ui.packages[lii].Checked = b
			}
			check.SetChecked(ui.packages[lii].Checked)
		},
	)
	ui.list.OnSelected = func(id widget.ListItemID) {
		ui.sideBar.Load(ui.packages[id])
		ui.list.UnselectAll()
	}
}

func (ui *ui) InitBoxes() {
	ui.mainBox = boxes.NewBorder(ui.search.Box, nil, nil, ui.sideBar.Box, ui.list)
}
