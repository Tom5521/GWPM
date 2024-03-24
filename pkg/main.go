package pkg

type Packager interface {
	Install() error
	Uninstall() error
	Version() string
	Name() string
	Installed() bool

	Manager() Managerer
}

type Managerer interface {
	Exists() bool
	Name() string
	RequireAdmin() bool
	Install(...Packager) error
	Uninstall(...Packager) error
	Version() string
	InstalledPkgs() ([]Packager, error)
}
