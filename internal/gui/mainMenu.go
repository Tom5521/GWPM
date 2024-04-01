package gui

import (
	"fyne.io/fyne/v2"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

type MainMenu struct {
	Menu *fyne.MainMenu
}

func (m *MainMenu) Init() {
	m.Menu = fyne.NewMainMenu(
		/*
			fyne.NewMenu("Manager",
				fyne.NewMenuItem("Reload", func() {
					cui.InitPkgSlice()
				}),
				fyne.NewMenuItem("Change", func() {

				}),
			),*/
		fyne.NewMenu("Packages",
			fyne.NewMenuItem("Select All", func() {
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if p.Checked {
						continue
					}
					p.Checked = true
					cui.list.RefreshItem(i)
				}
			}),
			fyne.NewMenuItem("Unselect all", func() {
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if !p.Checked {
						continue
					}
					p.Checked = false
					cui.list.RefreshItem(i)
				}
			}),
			fyne.NewMenuItem("Install selected", func() {
				cui.search.toggleLoading()
				pkgs := checkedPkgs()
				err := cui.manager.InstallByName(pkgs...)
				if err != nil {
					popups.Error(err)
				}
				cui.search.toggleLoading()
			}),
			fyne.NewMenuItem("Uninstall selected", func() {
				cui.search.toggleLoading()
				pkgs := checkedPkgs()
				err := cui.manager.UninstallByName(pkgs...)
				if err != nil {
					popups.Error(err)
				}
				cui.search.toggleLoading()
			}),
		),
	)
}

func checkedPkgs() []string {
	var pkgs []string
	for _, p := range cui.packages {
		if p.Checked {
			pkgs = append(pkgs, p.Name())
		}
	}
	return pkgs
}
