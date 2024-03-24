package choco

import (
	"fmt"
	"regexp"

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
	out, err := term.NewCommand("choco", "list").Output()
	if err != nil {
		return pkgs, err
	}
	regex := regexp.MustCompile(`([^\s]+)\s+([\d.]+)`)
	matches := regex.FindAllStringSubmatch(out, -1)

	for _, match := range matches {
		if match[1] == "chocolatey" {
			continue
		}
		pkgs = append(pkgs, &Package{
			name:    match[1],
			version: match[2],
		})
	}
	return pkgs, nil
}
