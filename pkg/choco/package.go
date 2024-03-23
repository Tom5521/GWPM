package choco

import (
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Package struct {
	name string

	HideCmdOnAction bool

	manager *Manager
}

func (p *Package) Install() error {
	cmd := term.NewCommand("choco", "install", p.name)
	cmd.Hide = p.HideCmdOnAction
	return cmd.Run()
}

func (p *Package) Uninstall() error {
	cmd := term.NewCommand("choco", "uninstall", p.name)
	cmd.Hide = p.HideCmdOnAction
	return cmd.Run()
}

func (p *Package) Version() string {
	// TODO: Fix this
	return ""
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
	for _, ip := range p.manager.InstalledPkgs() {
		if p.Name() == ip.Name() {
			return true
		}
	}
	return false
}
