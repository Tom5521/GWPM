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
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if !p.Checked {
						continue
					}
					err := p.Install()
					if err != nil {
						popups.Error(err)
					}
				}
			}),
			fyne.NewMenuItem("Uninstall selected", func() {
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if !p.Checked {
						continue
					}
					err := p.Uninstall()
					if err != nil {
						popups.Error(err)
					}
				}
			}),
		),
	)
}
