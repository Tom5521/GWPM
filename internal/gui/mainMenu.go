package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

type MainMenu struct {
	Menu *fyne.MainMenu
}

func (m *MainMenu) Init() {
	m.Menu = fyne.NewMainMenu(
		fyne.NewMenu(po.Get("File"),
			fyne.NewMenuItem(po.Get("Reload"), func() {
				InfiniteLoadingDialog(cui.InitPkgSlice)
			}),
		),
		fyne.NewMenu(po.Get("Packages"),
			fyne.NewMenuItem(po.Get("Select all"), func() {
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if p.Checked {
						continue
					}
					p.Checked = true
					cui.list.RefreshItem(i)
				}
			}),
			fyne.NewMenuItem(po.Get("Unselect all"), func() {
				for i := range cui.list.Length() {
					p := cui.packages[i]
					if !p.Checked {
						continue
					}
					p.Checked = false
					cui.list.RefreshItem(i)
				}
			}),
			fyne.NewMenuItem(po.Get("Install selected packages"), InstallSelected),
			fyne.NewMenuItem(po.Get("Uninstall selected packages"), UninstallSelected),
		),
		fyne.NewMenu(po.Get("Options"),
			fyne.NewMenuItem(po.Get("Settings"), func() {
				ShowSettingsWindow()
			}),
			fyne.NewMenuItem(po.Get("Change package manager"), func() {
				var selected string
				d := dialog.NewForm(
					po.Get("Select manager"),
					po.Get("Select"),
					po.Get("Cancel"),
					[]*widget.FormItem{
						widget.NewFormItem(po.Get("Manager:"), widget.NewSelect(Managers, func(s string) {
							selected = s
						})),
					},
					func(b bool) {
						if !b {
							return
						}
						if selected == "" {
							popups.Error(po.Get("No option selected."))
							return
						}
						if selected == cui.manager.Name() {
							return
						}
						cui.settings.SetString(ManagerID, selected)
						FuncLoadingDialog(
							cui.InitManager,
							cui.InitPkgSlice,
						)
					},
					cui.mainWindow,
				)
				d.Show()
			}),
		),
	)
}
