package choco_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Tom5521/GWPM/pkg/choco"
)

var m = choco.Connect()

func TestConnectSpeed(t *testing.T) {
	now := time.Now()
	choco.Connect()
	fmt.Println(time.Since(now))
}

func TestInstalledPkgs(t *testing.T) {
	pkgs, err := m.LocalPkgs()
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
	err := m.InstallByName("gsudo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	TestInstalledPkgs(t)
}

func TestLocalSearch(t *testing.T) {
	pkgs, err := m.SearchInLocal("sudo") // g -> gsudo
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	if len(pkgs) == 0 {
		fmt.Println("pkgs len is 0!")
		t.Fail()
	}
	for _, p := range pkgs {
		fmt.Printf("%s", p.Name())
	}
	fmt.Println()
}

func TestUninstallPkgs(t *testing.T) {
	err := m.UninstallByName("gsudo")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestFinalPkgs(t *testing.T) {
	TestInstalledPkgs(t)
}

func TestSearch(t *testing.T) {
	pkgs, err := m.SearchInRepo("vscode")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(pkgs) == 0 {
		fmt.Println("pkgs len is 0!")
		t.Fail()
	}
	for _, p := range pkgs {
		fmt.Printf("%s ", p.Name())
	}
	fmt.Println()
}
