package scoop

import (
	"github.com/Tom5521/GWPM/pkg"
)

type Package struct {
	name    string
	version string

	local, repo bool

	manager *Manager
}

func (p *Package) Install() error {
	return p.manager.Install(p)
}

func (p *Package) Uninstall() error {
	err := p.manager.Uninstall(p)
	if err != nil {
		return err
	}
	p.local = false
	return nil
}

func (p *Package) Reinstall() error {
	err := p.Uninstall()
	if err != nil {
		return err
	}
	err = p.Install()
	if err != nil {
		return err
	}
	return nil
}

func (p *Package) Version() string {
	return p.version
}

func (p *Package) Name() string {
	return p.name
}

func (p *Package) Manager() pkg.Managerer {
	return p.manager
}

func (p *Package) Installed() bool {
	lpkgs, err := p.manager.LocalPkgs()
	if err != nil {
		return false
	}
	for _, lp := range lpkgs {
		if lp.Name() == p.Name() {
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

func (p *Package) Upgrade() error {
	return p.manager.Upgrade(p)
}
