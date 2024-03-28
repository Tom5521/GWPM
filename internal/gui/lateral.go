package gui

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

type Lateral struct {
	*widget.Form
	pkg pkg.Packager

	Name      *widget.FormItem
	Version   *widget.FormItem
	Installed *widget.FormItem
	Manager   *widget.FormItem
	Local     *widget.FormItem
	Repo      *widget.FormItem

	Install struct {
		*widget.FormItem
		Widget *widget.Button
	}
	Uninstall struct {
		*widget.FormItem
		Widget *widget.Button
	}
}

func (l *Lateral) makeForm(text string, content ...any) *widget.FormItem {
	return widget.NewFormItem(text, widget.NewLabel(fmt.Sprint(content...)))
}

func (l *Lateral) Init() {
	l.Name = l.makeForm("Name:")
	l.Version = l.makeForm("Version:")
	l.Installed = l.makeForm("Is Installed:")
	l.Manager = l.makeForm("Manager:")
	l.Local = l.makeForm("In Local:")
	l.Repo = l.makeForm("In Repo:")

	l.Install.Widget = widget.NewButton("Install", func() {})
	l.Uninstall.Widget = widget.NewButton("Uninstall", func() {})
	l.Install.Widget.Disable()
	l.Uninstall.Widget.Disable()
	l.Install.FormItem = widget.NewFormItem("", l.Install.Widget)
	l.Uninstall.FormItem = widget.NewFormItem("", l.Uninstall.Widget)

	l.Form = widget.NewForm(
		l.Name,
		l.Version,
		l.Installed,
		l.Manager,
		l.Local,
		l.Repo,
		l.Install.FormItem,
		l.Uninstall.FormItem,
	)
}

func (l *Lateral) Load(p pkg.Packager) {
	if reflect.DeepEqual(p, l.pkg) {
		return
	}
	l.pkg = p
	setText := func(fi *widget.FormItem, txt ...any) {
		fi.Widget.(interface{ SetText(string) }).SetText(fmt.Sprint(txt...))
	}
	setText(l.Name, p.Name())
	setText(l.Version, p.Version())
	setText(l.Installed, p.Installed())
	setText(l.Manager, p.Manager().Name())
	setText(l.Local, p.Local())
	setText(l.Repo, p.Repo())

	refresh := func() {
		l.Load(p)
	}

	if p.Installed() {
		l.Install.Widget.SetText("Reinstall")
		l.Install.Widget.OnTapped = func() {
			err := p.Reinstall()
			if err != nil {
				popups.Error(err)
			}
			refresh()
		}
	} else {
		l.Uninstall.Widget.Disable()
		l.Install.Widget.SetText("Install")
		l.Install.Widget.OnTapped = func() {
			err := p.Install()
			if err != nil {
				popups.Error(err)
			}
			refresh()
		}
	}
}
