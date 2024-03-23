package choco

import (
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Manager struct {
	name         string
	exists       bool
	requireAdmin bool
	version      string

	HideActions bool
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	var pkglist []string
	for _, p := range pkgs {
		pkglist = append(pkglist, p.Name())
	}
	cmd := term.NewCommand("choco", "uninstall")
	cmd.Args = append(cmd.Args, pkglist...)
	cmd.Hide = m.HideActions
	return cmd.Run()
}

func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	var pkglist []string
	for _, p := range pkgs {
		pkglist = append(pkglist, p.Name())
	}
	cmd := term.NewCommand("choco", "uninstall")
	cmd.Args = append(cmd.Args, pkglist...)
	cmd.Hide = m.HideActions
	return cmd.Run()
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

func (m *Manager) InstalledPkgs() []pkg.Packager {
	// TODO: Made functional this.
	return nil
}
