package main

import (
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/magefile/mage/sh"
)

func DownloadOpenGL() error {
	file := "builds/opengl32.7z"
	url := "https://downloads.fdossena.com/geth.php?r=mesa64-latest"
	err := sh.Rm(file)
	if err != nil {
		return err
	}
	err = sh.Run("wget", "-O", file, url)
	if err != nil {
		return err
	}
	err = os.Chdir("builds")
	if err != nil {
		return err
	}
	err = sh.Rm("opengl32.dll")
	if err != nil {
		return err
	}
	err = sh.Rm("README.txt")
	if err != nil {
		return err
	}
	err = sh.Run("7z", "e", "opengl32.7z")
	if err != nil {
		return err
	}

	return os.Chdir("..")
}

func Release() error {
	err := build.App()
	if err != nil {
		return err
	}
	err = DownloadOpenGL()
	if err != nil {
		return err
	}
	err = makeZip()
	if err != nil {
		return err
	}
	return nil
}

func Clean() {
	remove := func(f string) {
		err := sh.Rm(f)
		if err != nil {
			msg.Error(err)
		}
	}
	dirs := []string{
		"choco.test.exe",
		"main-test.exe",
		"scoop.test.exe",
		"builds",
		"opengl32.dll",
	}
	for _, d := range dirs {
		remove(d)
	}
}
