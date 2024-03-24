package choco

import (
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Package struct {
	name string

	version string
	manager *Manager
}

func (p *Package) Install() error {
	cmd := term.NewCommand("choco", "install", "-y", p.name)
	cmd.Hide = p.manager.HideActions
	return cmd.Run()
}

func (p *Package) Uninstall() error {
	cmd := term.NewCommand("choco", "uninstall", "-y", p.name)
	cmd.Hide = p.manager.HideActions
	return cmd.Run()
}

func (p *Package) Version() string {
	return p.version
}

func (p *Package) Name() string {
	return p.name
}

func (p *Package) ExistsManager() bool {
	return p.manager.Exists()
}

func (p *Package) Manager() pkg.Managerer {
	return p.manager
}

func (p *Package) Installed() bool {
	ipkgs, _ := p.manager.InstalledPkgs()
	for _, ip := range ipkgs {
		if p.Name() == ip.Name() {
			return true
		}
	}
	return false
}
