package gui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
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
}

func (l *Lateral) Show() {
	l.makeItems()
	l.Form.Show()
}

func (l *Lateral) makeItems() {
	makeFormItem := func(text string, c ...any) *widget.FormItem {
		return widget.NewFormItem(text, widget.NewLabel(fmt.Sprint(c...)))
	}
	l.Name = makeFormItem("Name", l.pkg.Name())
	l.Version = makeFormItem("Name", l.pkg.Version())
	l.Installed = makeFormItem("Installed", l.pkg.Installed())
	l.Manager = makeFormItem("Manager", l.pkg.Manager().Name())
	l.Local = makeFormItem("Local", l.pkg.Local())
	l.Repo = makeFormItem("Repo", l.pkg.Repo())
}

func InitLateral(packager pkg.Packager) *Lateral {
	l := &Lateral{
		pkg: packager,
	}
	l.makeItems()
	l.Form = widget.NewForm(
		l.Name,
		l.Version,
		l.Installed,
		l.Manager,
		l.Local,
		l.Repo,
	)
	l.Hide()
	return l
}
