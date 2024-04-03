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
		fyne.NewMenu("Manager",
			fyne.NewMenuItem("Reload", func() {
				cui.InitPkgSlice()
			}),
			fyne.NewMenuItem("Change", func() {
				var selected string
				d := dialog.NewForm(
					"Select manager",
					"Select",
					"Cancel",
					[]*widget.FormItem{
						widget.NewFormItem("Manager:", widget.NewSelect(Managers, func(s string) {
							selected = s
						})),
					},
					func(b bool) {
						if !b {
							return
						}
						if selected == "" {
							popups.Error("No option selected.")
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
				LoadingDialog.Show()
				pkgs := checkedPkgs()
				err := cui.manager.InstallByName(pkgs...)
				if err != nil {
					popups.Error(err)
				}
				LoadingDialog.Hide()
			}),
			fyne.NewMenuItem("Uninstall selected", func() {
				LoadingDialog.Show()
				pkgs := checkedPkgs()
				err := cui.manager.UninstallByName(pkgs...)
				if err != nil {
					popups.Error(err)
				}
				LoadingDialog.Hide()
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
