package gui

import "github.com/Tom5521/GWPM/pkg/gui/popups"

func UninstallSelected() {
	UninstallingDialog.Show()
	pkgs := checkedPkgNames()
	err := cui.manager.UninstallByName(pkgs...)
	if err != nil {
		popups.Error(err)
	}
	cui.InitPkgSlice()
	UninstallingDialog.Hide()
}

func InstallSelected() {
	InstallingDialog.Show()
	pkgs := checkedPkgNames()
	err := cui.manager.InstallByName(pkgs...)
	if err != nil {
		popups.Error(err)
	}
	cui.InitPkgSlice()
	InstallingDialog.Hide()
}

func checkedPkgNames() []string {
	var pkgs []string
	for _, p := range cui.packages {
		if p.Checked {
			pkgs = append(pkgs, p.Name())
		}
	}
	return pkgs
}

func checkedPkgs() []packager {
	var pkgs []packager
	for _, p := range cui.packages {
		if p.Checked {
			pkgs = append(pkgs, p)
		}
	}
	return pkgs
}
