package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/magefile/mage/sh"
)

func copyfile(src, dest string) error {
	fmt.Printf("Copying %s file to %s\n", src, dest)
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func movefile(src, dest string) error {
	fmt.Printf("Moving %s file to %s\n", src, dest)
	source, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	err = os.WriteFile(dest, source, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}

func makeZip() error {
	var zipDir = "windows-tmp"
	if _, err := os.Stat(zipDir); os.IsNotExist(err) {
		fmt.Println("Making temporal dir...")
		err = os.Mkdir(zipDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat("builds/opengl32.dll"); os.IsNotExist(err) {
		err = DownloadOpenGL()
		if err != nil {
			return err
		}
	}
	err := copyfile("builds/opengl32.dll", zipDir+"/opengl32.dll")
	if err != nil {
		return err
	}
	if _, err = os.Stat("builds/gwpm.exe"); os.IsNotExist(err) {
		err = build.App()
		if err != nil {
			return err
		}
	}
	err = copyfile("builds/gwpm.exe", zipDir+"/gwpm.exe")
	if err != nil {
		return err
	}
	err = sh.Rm("builds/GWPMx64.zip")
	if err != nil {
		return err
	}
	err = os.Chdir(zipDir)
	if err != nil {
		return err
	}

	fmt.Println("Zipping content...")
	err = sh.RunV("zip", "-r", "../builds/GWPMx64.zip", ".")
	if err != nil {
		return err
	}
	err = os.Chdir("..")
	if err != nil {
		return err
	}
	fmt.Println("Cleaning...")
	err = os.RemoveAll(zipDir)
	if err != nil {
		return err
	}
	return nil
}

type Metadata struct {
	BuildNumber int `json:"build number"`
}

func (m *Metadata) Get() error {
	f, err := os.ReadFile("magefiles/meta.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(f, &m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Metadata) Write() error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = os.WriteFile("magefiles/meta.json", data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func buildNumberUp(m *Metadata) error {
	err := m.Get()
	if err != nil {
		return err
	}
	m.BuildNumber++
	err = m.Write()
	if err != nil {
		return err
	}
	return nil
}
