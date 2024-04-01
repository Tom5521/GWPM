package gui

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func InfiniteLoadingDialog(f func()) {
	bar := widget.NewProgressBarInfinite()
	bar.Start()
	d := dialog.NewCustomWithoutButtons(
		"Loading...",
		bar,
		cui.mainWindow,
	)
	d.Show()
	f()
	d.Hide()
}

func FuncLoadingDialog(funcs ...func()) {
	bar := widget.NewProgressBar()
	bar.Value = 0

	d := dialog.NewCustomWithoutButtons(
		"Loading...",
		bar,
		cui.mainWindow,
	)
	d.Show()
	v := float64(len(funcs))
	run := func(f func()) {
		f()
		bar.SetValue(bar.Value + (1.0 / v))
	}
	for _, f := range funcs {
		run(f)
	}
	d.Hide()
}

var LoadingDialog *dialog.CustomDialog

func InitLoadingDialog() {
	bar := widget.NewProgressBarInfinite()
	bar.Start()
	LoadingDialog = dialog.NewCustomWithoutButtons(
		"Loading...",
		bar,
		cui.mainWindow,
	)
}
