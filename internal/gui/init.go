package gui

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/locales"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/choco"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
	"github.com/Tom5521/GWPM/pkg/scoop"
	"github.com/leonelquinteros/gotext"
	"github.com/ncruces/zenity"

	boxes "fyne.io/fyne/v2/container"
)

var Managers = []string{
	choco.ManagerName,
	scoop.ManagerName,
}

const (
	ManagerID     = "manager"
	ListModeID    = "list-mode"
	SearchEntryID = "search-entry"
	LangID        = "lang"
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

var (
	cui *ui
	po  *gotext.Po
)

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
	InitLocales()
	cui.mainWindow = app.NewWindow(po.Get("Graphic Windows Package Manager"))
	cui.mainWindow.SetMaster()
	cui.mainWindow.Resize(fyne.NewSize(830, 390))

	InitDialogs()
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
func InitLocales() {
	if cui.settings.String(LangID) == "" {
		cui.settings.SetString(LangID, "en")
	}
	po = locales.GetPo(cui.settings.String(LangID))
}

func (ui *ui) InitManager() {
	manager := func(name string) pkg.Managerer {
		var manager pkg.Managerer
		switch name {
		case choco.ManagerName:
			manager = choco.Connect()
		case scoop.ManagerName:
			manager = scoop.Connect()
		default:
			ui.settings.SetString(ManagerID, choco.ManagerName)
			manager = choco.Connect()
		}
		return manager
	}
	ui.manager = manager(ui.settings.String(ManagerID))
	if !ui.manager.IsInstalled() {
		err := zenity.Question(
			po.Get("The current package manager is not installed,Do you want to install a package manager?"),
			zenity.Title(po.Get("Install a package manager?")),
			zenity.OKLabel(po.Get("Install")),
			zenity.CancelLabel(po.Get("Exit")),
		)
		if err != nil {
			os.Exit(1)
			return
		}
		selected, err := zenity.List(
			po.Get("Select a package manager to install"),
			Managers,
			zenity.OKLabel(po.Get("Install")),
		)
		if err != nil {
			os.Exit(0)
		}
		newManager := manager(selected)
		err = newManager.InstallManager()
		if err != nil {
			popups.FatalError(err)
		}
		ui.settings.SetString(ManagerID, newManager.Name())
		ui.manager = newManager
	}
}

func (ui *ui) InitPkgSlice() {
	var (
		packagers []pkg.Packager
		err       error
	)
	switch ui.settings.String(ListModeID) {
	case "local":
		packagers, err = ui.manager.LocalPkgs()
	case "repo":
		if ui.search.Entry.Text == "" {
			packagers = []pkg.Packager{}
			return
		}
		packagers, err = ui.manager.SearchInRepo(ui.search.Entry.Text)
	default:
		ui.settings.SetString(ListModeID, "local")
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
	if ui.list != nil {
		ui.list.Refresh()
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
	installBtn := widget.NewButton(po.Get("Install"), InstallSelected)
	uninstallBtn := widget.NewButton(po.Get("Uninstall"), UninstallSelected)
	buttonsBox := boxes.NewAdaptiveGrid(2, installBtn, uninstallBtn)

	ui.mainBox = boxes.NewBorder(ui.search.Box, buttonsBox, nil, ui.sideBar.Box, ui.list)
}
