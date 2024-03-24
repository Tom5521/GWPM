package pkg

import (
	"errors"
)

var (
	ErrManagerNotInstalled = errors.New("the package manager isn't installed")
	ErrPkgNotExists        = errors.New("the package not exists")
	ErrPkgNotInstalled     = errors.New("the package isn't installed")
	ErrNotAdministrator    = errors.New("not running as adminstrator")
)

// TODO:Comment and document this.
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

// TODO:Comment and document this.
// TODO: Unbloat this...?
type Managerer interface {
	Exists() bool
	Name() string
	RequireAdmin() bool
	Install(...string) error
	InstallPkgs(...Packager) error
	Uninstall(...string) error
	UninstallPkgs(...Packager) error
	Version() string
	InstalledPkgs() ([]Packager, error)
	RepoPkgByName(string) (Packager, error)
	LocalPkgByName(string) (Packager, error)
	IsInstalled(Packager) bool
	Search(string) ([]Packager, error)
	IsInRepo(Packager) bool
	IsInLocal(Packager) bool
}
