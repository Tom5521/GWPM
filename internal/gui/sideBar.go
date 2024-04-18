package gui

import (
	"fmt"
	"reflect"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

type SideBar struct {
	Box  *fyne.Container
	Form *widget.Form

	CurrentPkg pkg.Packager

	Slice []*widget.FormItem

	Name      *widget.FormItem
	Version   *widget.FormItem
	Installed *widget.FormItem
	Manager   *widget.FormItem
	Local     *widget.FormItem
	Repo      *widget.FormItem

	Close *widget.Button

	installBtn   *widget.Button
	uninstallBtn *widget.Button
	upgradeBtn   *widget.Button

	LoadBar *widget.ProgressBar
}

func (s *SideBar) Init() {
	newFormItem := func(title any, text ...any) *widget.FormItem {
		fi := widget.NewFormItem(po.Get(fmt.Sprint(title)), widget.NewLabel(fmt.Sprint(text...)))
		s.Slice = append(s.Slice, fi)
		return fi
	}
	s.Name = newFormItem("Name:")
	s.Version = newFormItem("Version:")
	s.Installed = newFormItem("Installed:")
	s.Manager = newFormItem("Manager:")
	s.Local = newFormItem("Local:")
	s.Repo = newFormItem("Repo:")

	s.Close = widget.NewButton(po.Get("Close"), func() {
		s.Clean()
		s.Box.Hide()
	})

	s.LoadBar = widget.NewProgressBar()
	s.LoadBar.Hide()

	s.Form = widget.NewForm(s.Slice...)

	s.Box = boxes.NewBorder(nil, nil, widget.NewSeparator(), nil,
		boxes.NewVBox(
			s.Form,
			s.LoadBar,
			s.MakeButtons(),
			widget.NewSeparator(),
			s.Close,
		),
	)
	s.Box.Hide()
}

func (s *SideBar) Load(p pkg.Packager) {
	if reflect.DeepEqual(s.CurrentPkg, p) {
		return
	}
	s.CurrentPkg = p
	if s.Box.Hidden {
		s.Box.Show()
	}

	setText := func(fi *widget.FormItem, txt ...any) {
		fi.Widget.(interface{ SetText(string) }).SetText(fmt.Sprint(txt...))
		if s.LoadBar.Hidden {
			s.LoadBar.Show()
		}
		s.LoadBar.SetValue(s.LoadBar.Value + (1.0 / 6.0))
	}
	s.Clean()

	go func() {
		installed := p.Installed()
		s.LoadButtons(packager{Packager: p}, installed)

		setText(s.Name, p.Name())
		setText(s.Version, p.Version())
		setText(s.Manager, p.Manager().Name())
		setText(s.Local, p.Local())
		setText(s.Repo, p.Repo())
		setText(s.Installed, installed)

		s.LoadBar.Hide()
		s.LoadBar.SetValue(0)
	}()
}

func (s *SideBar) Clean() {
	clean := func(fi *widget.FormItem) {
		fi.Widget.(interface{ SetText(string) }).SetText(po.Get("loading..."))
	}
	for _, i := range s.Slice {
		clean(i)
	}
}

func (s *SideBar) MakeButtons() *fyne.Container {
	s.installBtn = widget.NewButton(po.Get("Install"), nil)
	s.uninstallBtn = widget.NewButton(po.Get("Uninstall"), nil)
	s.upgradeBtn = widget.NewButton(po.Get("Upgrade"), nil)

	return boxes.NewVBox(
		boxes.NewAdaptiveGrid(2,
			s.installBtn,
			s.upgradeBtn,
		),
		s.uninstallBtn,
	)
}

func (s *SideBar) LoadButtons(p packager, installed bool) {
	s.upgradeBtn.Enable()
	s.installBtn.Enable()
	s.uninstallBtn.Enable()

	if !installed {
		s.upgradeBtn.Disable()
		s.uninstallBtn.Disable()
	} else {
		s.upgradeBtn.OnTapped = func() {
			LoadingDialog.Show()
			err := p.Upgrade()
			if err != nil {
				popups.Error(err)
			}
			LoadingDialog.Hide()
			s.Load(p)
		}
		s.uninstallBtn.OnTapped = func() {
			UninstallingDialog.Show()
			err := p.Uninstall()
			if err != nil {
				popups.Error(err)
			}
			UninstallingDialog.Hide()
			s.Load(p)
		}
	}
	s.installBtn.OnTapped = func() {
		InstallingDialog.Show()
		err := p.Install()
		if err != nil {
			popups.Error(err)
		}
		s.LoadButtons(p, !installed)
		InstallingDialog.Hide()
	}
}
