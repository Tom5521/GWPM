package choco_test

import (
	"fmt"
	"testing"

	"github.com/Tom5521/GWPM/pkg/choco"
)

var m = choco.Connect()

func TestInstalledPkgs(t *testing.T) {
	pkgs, err := m.InstalledPkgs()
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	for _, p := range pkgs {
		fmt.Println(p.Name())
	}
}

func TestInstallPkgs(t *testing.T) {
	err := m.Install("gsudo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	TestInstalledPkgs(t)
}

func TestUninstallPkgs(t *testing.T) {
	err := m.Uninstall("gsudo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestFinalPkgs(t *testing.T) {
	TestInstalledPkgs(t)
}

func TestSearch(t *testing.T) {
	pkgs, err := m.Search("vscode")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(pkgs) == 0 {
		fmt.Println("pkg len is 0!")
		t.Fail()
	}
	for _, p := range pkgs {
		fmt.Printf("%s ", p.Name())
	}
	fmt.Println()
}
