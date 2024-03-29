package gui

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	boxes "fyne.io/fyne/v2/container"
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

	newFormItem := func(title any, text ...any) *widget.FormItem {
		return widget.NewFormItem(fmt.Sprint(title), widget.NewLabel(fmt.Sprint(text...)))
	}

	var sideBarItems struct {
		currentPkg pkg.Packager

		Name      *widget.FormItem
		Version   *widget.FormItem
		Installed *widget.FormItem
		Manager   *widget.FormItem
		Local     *widget.FormItem
		Repo      *widget.FormItem

		loadingBar *widget.ProgressBar
	}
	sideBarItems.loadingBar = widget.NewProgressBar()
	sideBarItems.loadingBar.Hide()
	sideBarItems.Name = newFormItem("Name:")
	sideBarItems.Version = newFormItem("Version:")
	sideBarItems.Installed = newFormItem("Installed:")
	sideBarItems.Manager = newFormItem("Manager:")
	sideBarItems.Local = newFormItem("Local:")
	sideBarItems.Repo = newFormItem("Repo:")

	loadSidebar := func(p pkg.Packager) {
		if reflect.DeepEqual(p, sideBarItems.currentPkg) {
			return
		}
		sideBarItems.currentPkg = p

		setText := func(fi *widget.FormItem, txt ...any) {
			fi.Widget.(interface{ SetText(string) }).SetText(fmt.Sprint(txt...))
			if sideBarItems.loadingBar.Hidden {
				sideBarItems.loadingBar.Show()
			}
			sideBarItems.loadingBar.SetValue(sideBarItems.loadingBar.Value + (1.0 / 6.0))
		}
		func() {
			loadTxt := "loading..."
			clean := func(fi *widget.FormItem) {
				fi.Widget.(interface{ SetText(string) }).SetText(loadTxt)
			}
			clean(sideBarItems.Name)
			clean(sideBarItems.Version)
			clean(sideBarItems.Manager)
			clean(sideBarItems.Local)
			clean(sideBarItems.Repo)
			clean(sideBarItems.Installed)
		}()

		go func() {
			setText(sideBarItems.Name, p.Name())
			setText(sideBarItems.Version, p.Version())
			setText(sideBarItems.Manager, p.Manager().Name())
			setText(sideBarItems.Local, p.Local())
			setText(sideBarItems.Repo, p.Repo())
			setText(sideBarItems.Installed, p.Installed())

			sideBarItems.loadingBar.Hide()
			sideBarItems.loadingBar.SetValue(0)
		}()
	}
	sideBar := widget.NewForm(
		sideBarItems.Name,
		sideBarItems.Version,
		sideBarItems.Installed,
		sideBarItems.Manager,
		sideBarItems.Local,
		sideBarItems.Repo,
	)

	pkgList := widget.NewList(
		func() int { return len(pkgs) },
		func() fyne.CanvasObject {
			return boxes.NewBorder(nil, nil, nil, &widget.Check{}, &widget.Label{})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			c := co.(*fyne.Container)
			label := c.Objects[0].(*widget.Label)
			label.SetText(pkgs[lii].Name())

			check := c.Objects[1].(*widget.Check)
			check.OnChanged = func(b bool) {
				pkgs[lii].Checked = b
			}
		},
	)
	pkgList.OnSelected = func(id widget.ListItemID) {
		loadSidebar(pkgs[id].Packager)
		pkgList.UnselectAll()
	}

	topBar := boxes.NewHBox()

	sideBarBox := boxes.NewBorder(nil, nil, widget.NewSeparator(), nil,
		boxes.NewVBox(
			sideBar,
			sideBarItems.loadingBar,
		),
	)

	// Main Content
	var mcontent = boxes.NewBorder(topBar, nil, nil, sideBarBox, pkgList)

	mw.SetContent(mcontent)
	mw.ShowAndRun()
}
