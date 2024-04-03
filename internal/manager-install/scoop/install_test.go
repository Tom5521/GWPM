package scoop_test

import (
	"os/exec"
	"testing"

	"github.com/Tom5521/GWPM/internal/manager-install/scoop"
)

func TestInstall(t *testing.T) {
	err := scoop.Install()
	if err != nil {
		t.Fail()
	}
	_, err = exec.LookPath("choco")
	if err != nil {
		t.Fail()
	}
}
