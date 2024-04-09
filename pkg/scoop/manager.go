package scoop

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/term"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/sahilm/fuzzy"
)

const ManagerName = "Scoop"

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

func (m *Manager) LocalPkgs() ([]pkg.Packager, error) {
	var pkgs []pkg.Packager

	usr, err := user.Current()
	if err != nil {
		return pkgs, err
	}
	root := usr.HomeDir + "\\scoop\\apps\\"
	dirs, err := os.ReadDir(root)
	if err != nil {
		return pkgs, err
	}
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		if d.Name() == "scoop" {
			continue
		}
		file, err := os.ReadFile(root + d.Name() + "\\current\\manifest.json")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				msg.Warningf("Skipping %s...", d.Name())
				continue
			}
			return pkgs, err
		}
		var data map[string]any
		err = json.Unmarshal(file, &data)
		if err != nil {
			return pkgs, err
		}
		var version string
		for i, f := range data {
			if i == "version" {
				version = fmt.Sprint(f)
				break
			}
		}
		p := &Package{
			name:    d.Name(),
			version: version,
			manager: m,
			local:   true,
		}
		pkgs = append(pkgs, p)
	}
	return pkgs, nil
}

func (m *Manager) RepoPkgByName(p string) (pkg.Packager, error) {
	var rpkg *Package

	rpkgs, err := m.SearchInRepo(p)
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

func (m *Manager) SearchInRepo(p string) ([]pkg.Packager, error) {
	var pkgs []pkg.Packager
	out, err := term.NewCommand("scoop", "search", p).Output()
	if err != nil {
		return pkgs, err
	}

	if strings.Contains(out, "GitHub API rate limit reached.") {
		msg.Warning("Please try again later or configure your API token using 'scoop config gh_token <your token>'.")
		return pkgs, errors.New("github API rate limit reached")
	}
	if strings.Contains(out, "WARN  No matches found.") {
		msg.Warning("No matches found.")
		return pkgs, nil
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
			version: "unknown",
			repo:    true,
			manager: m,
		})
	}

	return pkgs, nil
}

func (m *Manager) IsInRepo(p pkg.Packager) bool {
	rpkgs, err := m.SearchInRepo(p.Name())
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

func (m *Manager) InstallManager() error {
	if m.isInstalled {
		return pkg.ErrManagerIsInstalled
	}
	const (
		command1 = "Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser"
		command2 = "Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression"
	)
	cmd := term.NewCommand("powershell", "-c", command1)
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = term.NewCommand("powershell", "-c", command2)
	return cmd.Run()
}
