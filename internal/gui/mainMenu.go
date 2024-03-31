package gui

import (
	"fyne.io/fyne/v2"
)

type MainMenu struct {
	Menu *fyne.MainMenu
}

func (m *MainMenu) Init() {
	m.Menu = fyne.NewMainMenu(
		fyne.NewMenu("Manager",
			fyne.NewMenuItem("Reload", func() {
				cui.InitPkgSlice()
			}),
			fyne.NewMenuItem("Change", func() {
				// TODO: Finish this.
			}),
		),
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
			fyne.NewMenuItem("Install selected", func() {}),
			fyne.NewMenuItem("Uninstall selected", func() {}),
		),
	)
}
