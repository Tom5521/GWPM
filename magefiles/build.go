package main

import (
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

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

func (Build) Release() error {
	err := sh.RunWithV(env, "go", "build", "-v", "-o=builds/gwpm.exe", ".")
	return err
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
