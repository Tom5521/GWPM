package scoop

import "github.com/Tom5521/GWPM/pkg"

type Package struct {
	name    string
	version string

	manager *Manager
}

func (p *Package) Install() error {
	return nil
}

func (p *Package) Uninstall() error {
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
	return true
}
