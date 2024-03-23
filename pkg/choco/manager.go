package choco

import (
	"github.com/Tom5521/GWPM/pkg"
)

type Manager struct {
	name         string
	exists       bool
	requireAdmin bool
	version      string
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	for _, p := range pkgs {
		err := p.Install()
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	for _, p := range pkgs {
		err := p.Uninstall()
		if err != nil {
			return err
		}
	}
	return nil
}
func (m *Manager) Version() string {
	return m.name
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
	// TODO: Made functional this
	return nil
}
