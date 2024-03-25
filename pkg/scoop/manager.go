package scoop

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
)

type Manager struct {
	name         string
	isInstalled  bool
	requireAdmin bool
	version      string

	HideActions bool
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) RequireAdmin() bool {
	return m.requireAdmin
}

func (m *Manager) InstallByName(pkgs ...string) error {
	cmd := term.NewCommand("scoop", "install")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	return cmd.Run()
}

func (m *Manager) Install(pkgs ...pkg.Packager) error {
	var strPkgs []string
	for _, p := range pkgs {
		strPkgs = append(strPkgs, p.Name())
	}
	return m.InstallByName(strPkgs...)
}

func (m *Manager) UninstallByName(pkgs ...string) error {
	cmd := term.NewCommand("scoop", "uninstall")
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Hide = m.HideActions
	return cmd.Run()
}

func (m *Manager) Uninstall(pkgs ...pkg.Packager) error {
	var strPkgs []string
	for _, p := range pkgs {
		strPkgs = append(strPkgs, p.Name())
	}
	return m.UninstallByName(strPkgs...)
}

func (m *Manager) Version() string {
	return m.version
}

/*
Export JSON Template
{
    "buckets": [
        {
            "Name": "main",
            "Source": "~\\scoop\\buckets\\main",
            "Updated": "\/Date(1711372226285)\/",
            "Manifests": 1310
        }
    ],
    "apps": [
        {
            "Info": "",
            "Source": "main",
            "Name": "psutils",
            "Version": "0.2023.06.28",
            "Updated": "\/Date(1711372676457)\/"
        }
    ]
}
*/

func (m *Manager) LocalPkgs() ([]pkg.Packager, error) {
	var pkgs []pkg.Packager

	var packages struct {
		Buckets []struct {
			Name      string `json:"Name"`
			Source    string `json:"Source"`
			Updated   string `json:"Updated"`
			Manifests int    `json:"Manifests"`
		} `json:"buckets"`
		Apps []struct {
			Info    string `json:"Info"`
			Source  string `json:"Source"`
			Name    string `json:"Name"`
			Version string `json:"Version"`
			Updated string `json:"Updated"`
		} `json:"apps"`
	}

	cmd := term.NewCommand("scoop", "export")
	cmd.Hide = true
	out, err := cmd.Output()
	if err != nil {
		return pkgs, err
	}
	err = json.Unmarshal([]byte(out), &packages)
	if err != nil {
		return pkgs, err
	}

	for _, p := range packages.Apps {
		pkgs = append(pkgs, &Package{
			name:    p.Name,
			version: p.Version,
			bucket:  p.Source,
			manager: m,
			local:   true,
		})
	}

	return pkgs, nil
}

func (m *Manager) RepoPkgByName(p string) (pkg.Packager, error) {
	var rpkg *Package

	rpkgs, err := m.Search(p)
	if err != nil {
		return rpkg, err
	}

	if len(rpkgs) == 0 {
		return rpkg, pkg.ErrPkgNotFound
	}

	for _, rp := range rpkgs {
		if rpkg.Name() == rp.Name() {
			return rp, nil
		}
	}

	return rpkg, pkg.ErrPkgNotFound
}

func (m *Manager) LocalPkgByName(name string) (pkg.Packager, error) {
	var lpkg *Package

	lpkgs, err := m.LocalPkgs()
	if err != nil {
		return lpkg, err
	}

	for _, p := range lpkgs {
		if p.Name() == name {
			return p, nil
		}
	}

	return lpkg, pkg.ErrPkgNotFound
}

func (m *Manager) IsInstalled() bool {
	return m.isInstalled
}

func (m *Manager) Search(p string) ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	out, err := term.NewCommand("scoop", "search", p).Output()
	if err != nil {
		return pkgs, err
	}

	if strings.Contains(out, "GitHub API rate limit reached.") {
		fmt.Println("Please try again later or configure your API token using 'scoop config gh_token <your token>'.")
		return pkgs, errors.New("github API rate limit reached")
	}

	lines := strings.Split(out, "\n")[4:] // [4:] to skip header
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		if parts[0] == "----" {
			continue
		}
		pkgs = append(pkgs, &Package{
			name:    parts[0],
			bucket:  parts[1],
			version: "unknown",
			repo:    true,
			manager: m,
		})
	}

	return pkgs, nil
}

func (m *Manager) IsInRepo(p pkg.Packager) bool {
	rpkgs, err := m.Search(p.Name())
	if err != nil {
		return false
	}
	for _, rpkg := range rpkgs {
		if rpkg.Name() == p.Name() {
			return true
		}
	}
	return false
}

func (m *Manager) IsInLocal(p pkg.Packager) bool {
	lpkgs, err := m.LocalPkgs()
	if err != nil {
		return false
	}
	for _, lpkg := range lpkgs {
		if lpkg.Name() == p.Name() {
			return true
		}
	}
	return false
}
