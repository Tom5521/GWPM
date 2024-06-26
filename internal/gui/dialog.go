package gui

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func InfiniteLoadingDialog(functions ...func()) {
	bar := widget.NewProgressBarInfinite()
	bar.Start()
	d := dialog.NewCustomWithoutButtons(
		po.Get("Loading..."),
		bar,
		cui.mainWindow,
	)
	d.Show()
	for _, f := range functions {
		f()
	}
	d.Hide()
}

func FuncLoadingDialog(funcs ...func()) {
	bar := widget.NewProgressBar()
	bar.Value = 0

	d := dialog.NewCustomWithoutButtons(
		po.Get("Loading..."),
		bar,
		cui.mainWindow,
	)
	d.Show()
	run := func(f func()) {
		f()
		bar.SetValue(bar.Value + (1.0 / float64(len(funcs))))
	}
	for _, f := range funcs {
		run(f)
	}
	d.Hide()
}

var (
	LoadingDialog      *dialog.CustomDialog
	InstallingDialog   *dialog.CustomDialog
	UninstallingDialog *dialog.CustomDialog
)

func makeRawProgressDialog(text string) *dialog.CustomDialog {
	bar := widget.NewProgressBarInfinite()
	bar.Start()
	return dialog.NewCustomWithoutButtons(
		po.Get(text),
		bar,
		cui.mainWindow,
	)
}

func InitDialogs() {
	LoadingDialog = makeRawProgressDialog("Loading...")
	InstallingDialog = makeRawProgressDialog("Installing...")
	UninstallingDialog = makeRawProgressDialog("Uninstalling...")
}
