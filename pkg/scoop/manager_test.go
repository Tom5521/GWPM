package scoop_test

import (
	"fmt"
	"testing"

	"github.com/Tom5521/GWPM/pkg/scoop"
)

var m = scoop.Connect()

func TestGetPkgs(t *testing.T) {
	pkgs, err := m.LocalPkgs()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(pkgs) == 0 {
		fmt.Println("pkgs len is 0!")
		t.Fail()
	}
	for _, p := range pkgs {
		fmt.Println(p.Name())
	}
}

func TestSearchPkgs(t *testing.T) {
	pkgs, err := m.Search("go")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if len(pkgs) == 0 {
		fmt.Println("pkgs len is 0!")
		t.Fail()
	}
	for _, p := range pkgs {
		fmt.Println("Package name:", p.Name())
	}
}
