package popups

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

func baseError(onAccept func(), e ...any) {
	app := fyne.CurrentApp()
	if app == nil {
		err := zenity.Error(fmt.Sprint(e...))
		if err != nil {
			messages.Warning(err)
			messages.Error(e...)
		}
		if onAccept != nil {
			onAccept()
		}
	}

	w := app.NewWindow("Error")

	errTitle := widget.NewLabel("Error")
	errTitle.Alignment = fyne.TextAlignCenter

	errText := widget.NewRichTextFromMarkdown(fmt.Sprintf("**%s**", fmt.Sprint(e...)))
	acceptButton := widget.NewButton("Accept", func() {
		w.Close()
		if onAccept != nil {
			onAccept()
		}
	})

	textBox := container.NewBorder(errTitle, nil, nil, nil, errText)
	content := container.NewBorder(nil, acceptButton, nil, nil, textBox)

	w.SetContent(content)
	w.Show()
}

func Error(e ...any) {
	baseError(nil, e...)
}

func FatalError(e ...any) {
	baseError(func() {
		log.Fatal()
	}, e...)
}
