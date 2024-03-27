package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/choco"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
	"github.com/Tom5521/GWPM/pkg/scoop"
)

func (ui *ui) MakeManager() (pkg.Managerer, error) {
	var m pkg.Managerer
	switch settings.String("manager") {
	case choco.ManagerName:
		m = choco.Connect()
	case scoop.ManagerName:
		m = scoop.Connect()
	default:
		settings.SetString("manager", choco.ManagerName) // Set the default manager
		m = choco.Connect()
	}
	if !m.IsInstalled() {
		return m, pkg.ErrManagerNotInstalled
	}

	return m, nil
}

func (ui *ui) MakeList(pkgs []pkg.Packager) *widget.List {
	return widget.NewList(
		func() int { return len(pkgs) },
		func() fyne.CanvasObject { return &widget.Label{} },
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(pkgs[lii].Name())
		},
	)
}

func (ui *ui) MakePkgList(m pkg.Managerer) []pkg.Packager {
	pkgs, err := m.LocalPkgs()
	if err != nil {
		popups.FatalError(err)
	}
	return pkgs
}
