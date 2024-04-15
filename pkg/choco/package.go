package choco

import (
	"github.com/Tom5521/GWPM/pkg"
)

// TODO:Document this.
type Package struct {
	name string

	local, repo bool

	version string
	manager *Manager
}

func (p *Package) Install() error {
	return p.manager.InstallByName(p.Name())
}

func (p *Package) Uninstall() error {
	return p.manager.UninstallByName(p.Name())
}

func (p *Package) Reinstall() error {
	return p.manager.Reinstall(p)
}

func (p *Package) Upgrade() error {
	return p.manager.Upgrade(p)
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
	ipkgs, _ := p.manager.LocalPkgs()
	for _, ip := range ipkgs {
		if p.Name() == ip.Name() {
			return true
		}
	}
	return false
}

func (p *Package) Local() bool {
	return p.local
}
func (p *Package) Repo() bool {
	return p.repo
}
