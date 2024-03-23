package choco

import (
	"os/exec"

	"github.com/Tom5521/GWPM/pkg/term"
)

func Connect() *Manager {
	return &Manager{
		name:         "Choco",
		requireAdmin: true,
		exists: func() bool {
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
			return string(out)
		}(),
	}
}
