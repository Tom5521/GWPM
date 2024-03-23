package scoop

import (
	"regexp"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Manager struct {
	name         string
	exists       bool
	requireAdmin bool
	version      string
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	return nil
}
func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	return nil
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
	out, err := term.NewCommand("scoop", "list").Output()
	if err != nil {
		return []pkg.Packager{}
	}
	re := regexp.MustCompile(`([^\s]+)\s+([\d.]+)\s+[^\s]+\s+[\d-]+\s+[\d:]+`)
	matches := re.FindAllStringSubmatch(out, -1)
	var pkgs []pkg.Packager
	for _, match := range matches {
		pkgs = append(pkgs, &Package{
			name:    match[1],
			version: match[2],
		})
	}
	return pkgs
}
