package main

import (
	"fmt"
	"runtime"
	"strconv"
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
			"CC":          "zig cc -target x86_64-windows-gnu",
			"CGO_ENABLED": "1",
		}
	}
	return env
}()

func (Build) App() error {
	tags, err := sh.OutCmd("git", "tag")()
	if err != nil {
		return err
	}
	lines := strings.Split(tags, "\n")
	version := lines[len(lines)-1]

	m := new(Metadata)
	err = buildNumberUp(m)
	if err != nil {
		return err
	}

	flags := fmt.Sprintf("internal/meta.DevBuildStr=false internal/meta.ReleaseStr=true internal/meta.Version=%s internal/meta.BuildNumber=%v", version, m.BuildNumber)
	return build.WithLdflags(flags)
}

func (Build) Dev() error {
	m := new(Metadata)
	err := buildNumberUp(m)
	if err != nil {
		return err
	}
	flags := "internal/meta.DevBuildStr=true internal/meta.ReleaseStr=false internal/meta.Version=devBuild internal/meta.BuildNumber=" + strconv.Itoa(m.BuildNumber)
	return build.WithLdflags(flags)
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
