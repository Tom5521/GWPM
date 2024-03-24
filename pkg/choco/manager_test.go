package choco_test

import (
	"fmt"
	"testing"

	"github.com/Tom5521/GWPM/pkg/choco"
)

func TestInstalledPkgs(t *testing.T) {
	m := choco.Connect()
	pkgs, err := m.InstalledPkgs()
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(pkgs)
}
