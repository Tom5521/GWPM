package scoop

import (
	"os"
	"os/exec"
)

func Install() error {
	const (
		command1 = "Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser"
		command2 = "Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression"
	)
	cmd := exec.Command("powershell", "-c", command1)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd.Args = cmd.Args[0:]
	cmd.Args = append(cmd.Args, command2)
	return cmd.Run()
}
