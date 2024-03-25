package scoop

import (
	"github.com/Tom5521/GWPM/pkg"
)

// TODO:Finish this.
type Package struct {
	name    string
	version string

	bucket string

	local, repo bool

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

func (p *Package) Local() bool {
	return p.local
}

func (p *Package) Repo() bool {
	return p.repo
}

func (p *Package) Bucket() string {
	return p.bucket
}

/*
Methods

type Packager interface {
	Install() error
	Uninstall() error
	Version() string
	Name() string
	Installed() bool
	Manager() Managerer
	Local() bool
	Repo() bool
}

*/
