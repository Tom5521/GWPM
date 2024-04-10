package gui

import (
	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/internal/gui/credits"
	"github.com/Tom5521/GWPM/locales"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
	"github.com/ncruces/zenity"
)

func ShowSettingsWindow() {
	w := cui.app.NewWindow("Settings")
	w.Resize(fyne.NewSize(390, 420))

	pkgManagerLabel := widget.NewLabel(po.Get("Package manager:"))
	pkgManagerSelect := widget.NewSelect(Managers, func(s string) {})
	pkgManagerSelect.SetSelected(cui.manager.Name())
	pkgManagerSelect.OnChanged = func(s string) {
		if s == "" {
			popups.Error(po.Get("No option selected."))
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

	creditsBtn := widget.NewButton(po.Get("Credits"), func() {
		credits.CreditsWindow(cui.app, fyne.NewSize(770, 430)).Show()
	})

	langLabel := widget.NewLabel(po.Get("Language:"))
	langSelect := widget.NewSelect(locales.Languages, func(s string) {})
	langSelect.SetSelectedIndex(func() int {
		switch cui.settings.String(LangID) {
		case "en":
			return 1
		case "es":
			return 0
		default:
			return 0
		}
	}())
	langSelect.OnChanged = func(s string) {
		var lang string
		switch s {
		case locales.Languages[0]:
			lang = "es"
		case locales.Languages[1]:
			lang = "en"
		default:
			lang = "en"
		}
		cui.settings.SetString(LangID, lang)
		po.Parse(locales.GetParser(lang))
		zenity.Info(po.Get("Restart the application to see it in your language."))
	}

	content := boxes.NewVBox(
		boxes.NewHBox(pkgManagerLabel, pkgManagerSelect),
		boxes.NewHBox(langLabel, langSelect),
		creditsBtn,
	)

	w.SetContent(content)
	w.Show()
}
