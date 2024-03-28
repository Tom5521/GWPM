package popups

import (
	"fmt"

	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/ncruces/zenity"
)

var runningFatal bool

func baseError(onAccept func(), e ...any) {
	app := fyne.CurrentApp()
	if app == nil {
		err := zenity.Error(fmt.Sprint(e...))
		if err != nil {
			msg.Warning(err)
			msg.Error(e...)
		}
		if onAccept != nil {
			onAccept()
		}
	}

	w := app.NewWindow("Error")

	errTitle := boxes.NewCenter(widget.NewRichTextFromMarkdown("# Error"))

	errText := &widget.Label{
		Text:      fmt.Sprint(e...),
		TextStyle: fyne.TextStyle{Bold: true},
		Alignment: fyne.TextAlignCenter,
	}

	acceptButton := widget.NewButton("Accept", func() {
		w.Close()
		if onAccept != nil {
			onAccept()
		}
	})

	textBox := boxes.NewBorder(errTitle, nil, nil, nil, errText)
	content := boxes.NewBorder(nil, acceptButton, nil, nil, textBox)

	w.SetContent(content)
	w.Show()
}

func Error(e ...any) {
	baseError(func() {
		msg.Error(e...)
	}, e...)
}

func FatalError(e ...any) {
	if runningFatal {
		return
	}
	runningFatal = true
	baseError(func() {
		msg.FatalError(e...)
	}, e...)
}
