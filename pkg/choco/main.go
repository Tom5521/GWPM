package choco

import (
	"os/exec"

	"github.com/Tom5521/GWPM/pkg/term"
)

// NOTE:Choco is fucking slow.
func Connect() *Manager {
	return &Manager{
		name:         ManagerName,
		requireAdmin: true,
		isInstalled: func() bool {
			_, err := exec.LookPath("choco")
			return err == nil
		}(),
		version: func() string {
			cmd := term.NewCommand("choco", "--version")
			cmd.Hide = true
			out, err := cmd.Output()
			if err != nil {
				return ""
			}
			return out
		}(),
	}
}
