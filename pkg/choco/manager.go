package choco

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/perm"
	"github.com/Tom5521/GWPM/pkg/term"

	"github.com/sahilm/fuzzy"
)

const ManagerName = "Chocolatey"

// TODO:Document this.
type Manager struct {
	name         string
	isInstalled  bool
	requireAdmin bool
	version      string

	HideActions bool
}

func (m *Manager) InstallByName(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "install", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions

	return cmd.Run()
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.InstallByName(packages...)
}

func (m *Manager) UninstallByName(pkgs ...string) error {
	if m.requireAdmin && !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	cmd := term.NewCommand("choco", "uninstall", "-y")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	fmt.Println(cmd.Make())

	return cmd.Run()
}

func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	var packages []string
	for _, p := range pkgs {
		packages = append(packages, p.Name())
	}
	return m.UninstallByName(packages...)
}

func (m *Manager) Version() string {
	return m.version
}
func (m *Manager) Exists() bool {
	return m.isInstalled
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) RequireAdmin() bool {
	return m.requireAdmin
}

func (m *Manager) makePkgList(list string, onIteration func(name, version string) *Package) []pkg.Packager {
	var pkgs []pkg.Packager
	// strings.Split > regex
	lines := strings.Split(list, "\n")
	for _, line := range lines {
		p := strings.SplitN(line, "|", 2)
		if len(p) < 2 {
			continue
		}
		pkgs = append(pkgs, onIteration(p[0], p[1]))
	}
	return pkgs
}

func (m *Manager) LocalPkgs() ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.isInstalled {
		return pkgs, pkg.ErrManagerNotInstalled
	}
	out, err := term.NewCommand("choco", "list", "-r").Output()
	if err != nil {
		return pkgs, err
	}

	pkgs = m.makePkgList(out, func(name, version string) *Package {
		return &Package{
			name:    name,
			version: version,
			manager: m,
			local:   true,
		}
	})

	return pkgs, nil
}

func (m *Manager) RepoPkgByName(name string) (pkg.Packager, error) {
	var lpkg *Package
	s, err := m.SearchInRepo(name)
	if err != nil {
		return lpkg, err
	}
	if len(s) == 0 {
		return lpkg, pkg.ErrPkgNotFound
	}
	lpkg, ok := s[0].(*Package)
	if !ok {
		return lpkg, pkg.ErrPkgNotFound
	}
	return lpkg, nil
}

func (m *Manager) LocalPkgByName(name string) (pkg.Packager, error) {
	var lpkg *Package

	if !m.IsInLocal(&Package{name: name}) {
		return lpkg, pkg.ErrPkgNotExists
	}

	s, err := m.LocalPkgs()
	if err != nil {
		return lpkg, err
	}

	for _, p := range s {
		if p.Name() == name {
			var ok bool
			lpkg, ok = p.(*Package)
			if !ok {
				return &Package{}, pkg.ErrPkgNotFound
			}
		}
	}

	return lpkg, nil
}

func (m *Manager) IsInstalled() bool {
	return m.isInstalled
}

func (m *Manager) SearchInLocal(p string) ([]pkg.Packager, error) {
	var spkgs []string
	ipkgs, err := m.LocalPkgs()
	if err != nil {
		return nil, err
	}

	for _, ip := range ipkgs {
		spkgs = append(spkgs, ip.Name())
	}
	var results []pkg.Packager

	findResults := fuzzy.Find(p, spkgs)
	for _, r := range findResults {
		results = append(results, ipkgs[r.Index])
	}
	return results, nil
}

func (m *Manager) SearchInRepo(pkgName string) ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	if !m.isInstalled {
		return pkgs, pkg.ErrManagerNotInstalled
	}
	// strings.Split > regex
	cmd := term.NewCommand("choco", "search", "-r", pkgName)
	out, err := cmd.Output()
	if err != nil {
		return pkgs, err
	}

	pkgs = m.makePkgList(out, func(name, version string) *Package {
		return &Package{
			name:    name,
			version: version,
			repo:    true,
			manager: m,
		}
	})

	var strPkgs []string
	for _, p := range pkgs {
		strPkgs = append(strPkgs, p.Name())
	}
	matches := fuzzy.Find(pkgName, strPkgs)
	var results []pkg.Packager
	for _, match := range matches {
		results = append(results, pkgs[match.Index])
	}

	return results, nil
}

func (m *Manager) IsInRepo(pkg pkg.Packager) bool {
	repoPkgs, err := m.SearchInRepo(pkg.Name())
	if err != nil {
		return false
	}
	for _, p := range repoPkgs {
		if p.Name() == pkg.Name() {
			return true
		}
	}
	return false
}

func (m *Manager) IsInLocal(p pkg.Packager) bool {
	ipkgs, err := m.LocalPkgs()
	if err != nil {
		return false
	}
	for _, pkg := range ipkgs {
		if pkg.Name() == p.Name() {
			return true
		}
	}
	return false
}

func (m *Manager) InstallManager() error {
	if m.isInstalled {
		return pkg.ErrManagerIsInstalled
	}
	if !perm.IsAdmin {
		return pkg.ErrNotAdministrator
	}
	const command = "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))"
	cmd := term.NewCommand("powershell", "-c", command)
	return cmd.Run()
}

func (m *Manager) UpgradeByName(names ...string) error {
	cmd := term.NewCommand("choco", "upgrade", "-y")
	cmd.Args = append(cmd.Args, names...)
	cmd.Hide = m.HideActions

	return cmd.Run()
}

func (m *Manager) Upgrade(pkgs ...pkg.Packager) error {
	var pkgNames []string
	for _, p := range pkgs {
		pkgNames = append(pkgNames, p.Name())
	}
	return m.UpgradeByName(pkgNames...)
}

func (m *Manager) ReinstallByName(names ...string) error {
	cmd := term.NewCommand("choco", "install", "-y", "--force")
	cmd.Args = append(cmd.Args, names...)
	cmd.Hide = m.HideActions

	return cmd.Run()
}

func (m *Manager) Reinstall(pkgs ...pkg.Packager) error {
	var pkgNames []string
	for _, p := range pkgs {
		pkgNames = append(pkgNames, p.Name())
	}

	return m.ReinstallByName(pkgNames...)
}
