package choco

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

func (m *Manager) InstalledPkgs() ([]pkg.Packager, error) {
	out, err := term.NewCommand("choco", "list").Output()
	if err != nil {
		return []pkg.Packager{}, err
	}
	regex := regexp.MustCompile(`([^\s]+)\s+([\d.]+)`)
	matches := regex.FindAllStringSubmatch(out, -1)
	var pkgs []pkg.Packager
	for _, match := range matches {
		pkgs = append(pkgs, &Package{
			name:    match[1],
			version: match[2],
		})
	}
	return pkgs, nil
}
