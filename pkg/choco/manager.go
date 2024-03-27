package choco

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/perm"
	"github.com/Tom5521/GWPM/pkg/term"
)

const ManagerName = "Chocolatey"

// TODO:Document this.
type Manager struct {
	name         string
	isInstalled  bool
	requireAdmin bool
	version      string

	HideActions bool
}

func (m *Manager) InstallByName(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "install", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	fmt.Println(cmd.Make())

	return cmd.Run()
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.InstallByName(packages...)
}

func (m *Manager) UninstallByName(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "uninstall", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	fmt.Println(cmd.Make())

	return cmd.Run()
}

func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.UninstallByName(packages...)
}

func (m *Manager) Version() string {
	return m.version
}
func (m *Manager) Exists() bool {
	return m.isInstalled
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) RequireAdmin() bool {
	return m.requireAdmin
}

func (m *Manager) LocalPkgs() ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.isInstalled {
		return pkgs, pkg.ErrManagerNotInstalled
	}
	out, err := term.NewCommand("choco", "list", "-r").Output()
	if err != nil {
		return pkgs, err
	}
	// strings.Split > regex
	lines := strings.Split(out, "\n")
	for _, line := range lines {
		p := strings.SplitN(line, "|", 2)
		if len(p) < 2 {
			continue
		}
		pkgs = append(pkgs, &Package{
			name:    p[0],
			version: p[1],
			manager: m,
			local:   true,
		})
	}

	return pkgs, nil
}

func (m *Manager) RepoPkgByName(name string) (pkg.Packager, error) {
	var lpkg *Package
	s, err := m.Search(name)
	if err != nil {
		return lpkg, err
	}
	if len(s) == 0 {
		return lpkg, pkg.ErrPkgNotFound
	}
	lpkg, ok := s[0].(*Package)
	if !ok {
		return lpkg, pkg.ErrPkgNotFound
	}
	return lpkg, nil
}

func (m *Manager) LocalPkgByName(name string) (pkg.Packager, error) {
	var lpkg *Package

	if !m.IsInLocal(&Package{name: name}) {
		return lpkg, pkg.ErrPkgNotExists
	}

	s, err := m.LocalPkgs()
	if err != nil {
		return lpkg, err
	}

	for _, p := range s {
		if p.Name() == name {
			var ok bool
			lpkg, ok = p.(*Package)
			if !ok {
				return &Package{}, pkg.ErrPkgNotFound
			}
		}
	}

	return lpkg, nil
}

func (m *Manager) IsInstalled() bool {
	return m.isInstalled
}

func (m *Manager) Search(pkgName string) ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.isInstalled {
		return pkgs, pkg.ErrManagerNotInstalled
	}
	// strings.Split > regex
	cmd := term.NewCommand("choco", "search", "-r", pkgName)
	out, err := cmd.Output()
	if err != nil {
		return pkgs, err
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		p := strings.SplitN(line, "|", 2)
		if len(p) < 2 {
			continue
		}
		pkgs = append(pkgs, &Package{
			name:    p[0],
			version: p[1],
			manager: m,
			repo:    true,
		})
	}

	return pkgs, nil
}

func (m *Manager) IsInRepo(pkg pkg.Packager) bool {
	repoPkgs, err := m.Search(pkg.Name())
	if err != nil {
		return false
	}
	for _, p := range repoPkgs {
		if p.Name() == pkg.Name() {
			return true
		}
	}
	return false
}

func (m *Manager) IsInLocal(p pkg.Packager) bool {
	ipkgs, err := m.LocalPkgs()
	if err != nil {
		return false
	}
	for _, pkg := range ipkgs {
		if pkg.Name() == p.Name() {
			return true
		}
	}
	return false
}
