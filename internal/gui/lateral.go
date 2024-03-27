package gui

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
)

type Lateral struct {
	*widget.Form

	Name      *widget.FormItem
	Version   *widget.FormItem
	Installed *widget.FormItem
	Manager   *widget.FormItem
	Local     *widget.FormItem
	Repo      *widget.FormItem
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
	l.Form = widget.NewForm(
		l.Name,
		l.Version,
		l.Installed,
		l.Manager,
		l.Local,
		l.Repo,
	)
}

func (l *Lateral) Load(p pkg.Packager) {
	setText := func(fi *widget.FormItem, txt ...any) {
		fi.Widget.(*widget.Label).SetText(fmt.Sprint(txt...))
	}
	setText(l.Name, p.Name())
	setText(l.Version, p.Version())
	setText(l.Installed, p.Installed())
	setText(l.Manager, p.Manager().Name())
	setText(l.Local, p.Local())
	setText(l.Repo, p.Repo())
}
