package gui

import (
	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/internal/gui/credits"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

func ShowSettingsWindow() {
	w := cui.app.NewWindow("Settings")
	w.Resize(fyne.NewSize(390, 420))

	pkgManagerLabel := widget.NewLabel("Package manager:")
	pkgManagerSelect := widget.NewSelect(Managers, func(s string) {})
	pkgManagerSelect.SetSelected(cui.manager.Name())
	pkgManagerSelect.OnChanged = func(s string) {
		if s == "" {
			popups.Error("No option selected.")
			pkgManagerSelect.SetSelected(cui.manager.Name())
			return
		}
		if s == cui.manager.Name() {
			return
		}
		cui.settings.SetString(ManagerID, s)
		FuncLoadingDialog(
			cui.InitManager,
			cui.InitPkgSlice,
		)
	}

	creditsBtn := widget.NewButton("Credits", func() {
		credits.CreditsWindow(cui.app, fyne.NewSize(770, 430)).Show()
	})

	// TODO: Add a language setting...

	content := boxes.NewVBox(
		boxes.NewHBox(pkgManagerLabel, pkgManagerSelect),
		creditsBtn,
	)

	w.SetContent(content)
	w.Show()
}
