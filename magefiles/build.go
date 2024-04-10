package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

var build Build

var env = func() map[string]string {
	var env map[string]string
	if runtime.GOOS != "windows" {
		env = map[string]string{
			"GOOS":        "windows",
			"CC":          "x86_64-w64-mingw32-gcc",
			"CGO_ENABLED": "1",
		}
	}
	return env
}()

func (Build) App() error {
	err := sh.RunWithV(env, "go", "build", "-v", "-o=builds/gwpm.exe", ".")
	return err
}

func (Build) WithLdflags(input string) error {
	ldflags := strings.Fields(input)
	var args []string
	args = append(args, "build", "-v", "-o=builds/gwpm.exe")

	if len(ldflags) > 0 {
		arg := `-ldflags=%s`
		var flags string
		for _, flag := range ldflags {
			flags += fmt.Sprintf(`-X 'github.com/Tom5521/GWPM/%s' `, flag)
		}
		args = append(args, fmt.Sprintf(arg, flags))
	}
	args = append(args, ".")
	return sh.RunWithV(env, "go", args...)
}

func (Build) Tests() error {
	err := sh.RunWithV(env, "go", "test", "-c", "pkg/choco/choco_test.go")
	if err != nil {
		return err
	}
	err = sh.RunWithV(env, "go", "test", "-c", "pkg/scoop/scoop_test.go")
	if err != nil {
		return err
	}
	return nil
}
