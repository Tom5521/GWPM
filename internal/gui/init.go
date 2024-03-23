package gui

import (
	"os"

	"github.com/Tom5521/gtk4tools/pkg/widgets"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type ui struct {
	app        *gtk.Application
	mainWindow *gtk.ApplicationWindow

	pkgList *widgets.List
}

func (ui *ui) Initialize(app *gtk.Application) {
	ui.mainWindow = gtk.NewApplicationWindow(app)
	ui.pkgList = widgets.NewList(
		[]string{},
		widgets.SelectionNone,
		func(listitem *gtk.ListItem) {},
		func(listitem *gtk.ListItem, obj string) {},
	)
}

func Initgtk() {
	app := gtk.NewApplication("com.test.window", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		activate(app)
	})
	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	ui := &ui{}
	ui.Initialize(app)

	ui.mainWindow.Show()
}
