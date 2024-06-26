package pkg

import (
	"errors"
)

var (
	ErrManagerNotInstalled = errors.New("the package manager isn't installed")
	ErrPkgNotExists        = errors.New("the package not exists")
	ErrPkgNotInstalled     = errors.New("the package isn't installed")
	ErrNotAdministrator    = errors.New("not running as adminstrator")
	ErrPkgNotFound         = errors.New("package not found")
	ErrManagerIsInstalled  = errors.New("the package manager is not installed")
)

// TODO:Comment and document this.
type Packager interface {
	Install() error
	Uninstall() error
	Reinstall() error
	Upgrade() error
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
	Name() string
	RequireAdmin() bool
	InstallByName(...string) error
	Install(...Packager) error
	UninstallByName(...string) error
	Uninstall(...Packager) error
	UpgradeByName(...string) error
	Upgrade(...Packager) error
	ReinstallByName(...string) error
	Reinstall(...Packager) error
	Version() string
	LocalPkgs() ([]Packager, error)
	RepoPkgByName(string) (Packager, error)
	LocalPkgByName(string) (Packager, error)
	IsInstalled() bool
	SearchInRepo(string) ([]Packager, error)
	SearchInLocal(string) ([]Packager, error)
	IsInRepo(Packager) bool
	IsInLocal(Packager) bool
	InstallManager() error
}
