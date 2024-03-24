package choco

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/perm"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Manager struct {
	name         string
	exists       bool
	requireAdmin bool
	version      string

	HideActions bool
}

func (m *Manager) Install(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "install", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	fmt.Println(cmd.Make())

	return cmd.Run()
}

func (m *Manager) InstallPkgs(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.Install(packages...)
}

func (m *Manager) Uninstall(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "uninstall", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	fmt.Println(cmd.Make())

	return cmd.Run()
}

func (m *Manager) UninstallPkgs(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.Uninstall(packages...)
}

func (m *Manager) Version() string {
	return m.version
}
func (m *Manager) Exists() bool {
	return m.exists
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) RequireAdmin() bool {
	return m.requireAdmin
}

func (m *Manager) InstalledPkgs() ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.exists {
		return pkgs, pkg.ErrManagerNotInstalled
	}
	out, err := term.NewCommand("choco", "list", "-r").Output()
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
			local:   true,
		})
	}

	return pkgs, nil
}
func (m *Manager) RepoPkgByName(name string) (pkg.Packager, error) {

	return nil, nil
}

func (m *Manager) LocalPkgByName(name string) (pkg.Packager, error) {
	return nil, nil
}

func (m *Manager) IsInstalled(p pkg.Packager) bool {
	ipkgs, err := m.InstalledPkgs()
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

func (m *Manager) Search(pkgName string) ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.exists {
		return pkgs, pkg.ErrManagerNotInstalled
	}
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

func (m *Manager) IsInLocal(pkg pkg.Packager) bool {
	localPkgs, err := m.Search(pkg.Name())
	if err != nil {
		return false
	}
	for _, p := range localPkgs {
		if p.Name() == pkg.Name() {
			return true
		}
	}
	return false
}
