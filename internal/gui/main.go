package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/choco"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
	"github.com/Tom5521/GWPM/pkg/scoop"
)

/*
To summarize, it's fucking hellish, or rather, I don't know fyne's design patterns,
so I'll just put it all in one function.

Is it inefficient?
Yes
Is it hard to read?
Not as hard as the improvised option

DID YOU REALLY HAVE TO USE ABBREVIATIONS?
Yes, that's what comments are for, right?

In conclusion, I'll keep this simple but efficient, nothing more, nothing less.
*/
func InitGUI() {
	app := app.NewWithID("com.github.tom5521.gwpm")
	settings := app.Preferences()

	// Main Window
	var mw = app.NewWindow("Graphic Windows Package Manager") // Main Window
	mw.SetMaster()
	mw.Resize(fyne.NewSize(830, 390))

	// Current Manager.
	var m pkg.Managerer
	switch settings.String("manager") {
	case choco.ManagerName:
		m = choco.Connect()
	case scoop.ManagerName:
		m = scoop.Connect()
	default:
		settings.SetString("manager", choco.ManagerName)
		m = choco.Connect()
	}

	type packager struct {
		pkg.Packager
		Checked bool
	}

	var pkgs []packager
	updatePkgs := func() {
		packagers, err := m.LocalPkgs()
		if err != nil {
			popups.FatalError(err)
		}
		for _, p := range packagers {
			pkgs = append(pkgs, packager{
				Packager: p,
			})
		}
	}
	updatePkgs()

	pkgList := widget.NewList(
		func() int { return len(pkgs) },
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, &widget.Check{}, &widget.Label{})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := co.(*fyne.Container)
			check := c.Objects[0].(*widget.Check)
			check.OnChanged = func(b bool) {
				pkgs[lii].Checked = b
			}
			label := c.Objects[1].(*widget.Label)
			label.SetText(pkgs[lii].Name())
		},
	)

	sideBar := widget.NewForm()
	topBar := widget.NewToolbar()

	// Main Content
	var mcontent = container.NewBorder(topBar, nil, nil, sideBar, pkgList)

	mw.SetContent(mcontent)
	mw.ShowAndRun()
}
