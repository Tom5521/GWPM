package scoop

import (
	"os/exec"
	"strings"

	"github.com/Tom5521/GWPM/pkg/term"
)

func Connect() *Manager {
	return &Manager{
		name:         ManagerName,
		requireAdmin: false,
		isInstalled: func() bool {
			_, err := exec.LookPath("scoop")
			return err == nil
		}(),
		version: func() string {
			cmd := term.NewCommand("scoop", "--version")
			cmd.Hide = true
			out, err := cmd.Output()
			if err != nil {
				return ""
			}
			line := strings.Split(out, "\n")[1]
			v := strings.Split(line, "-")

			return v[0]
		}(),
	}
}
