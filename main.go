package main

import (
	"runtime"

	"github.com/Tom5521/GWPM/internal/gui"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func main() {
	if runtime.GOOS != "windows" {
		msg.Warningf("Why are you running this on %s?", runtime.GOOS)
	}
	gui.InitGUI()
}
