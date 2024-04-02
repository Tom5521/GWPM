package gui

import (
	"fyne.io/fyne/v2"
	boxes "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Tom5521/GWPM/pkg"
	"github.com/Tom5521/GWPM/pkg/gui/popups"
)

type Search struct {
	Box *fyne.Container

	ModeSelect *widget.Select

	Entry  *widget.Entry
	Button *widget.Button
}

func (s *Search) Init() {
	s.Entry = widget.NewEntry()
	s.Entry.OnChanged = func(str string) {
		cui.settings.SetString("search-entry", str)
	}
	s.Entry.SetText(cui.settings.String("search-entry"))

	s.Button = widget.NewButton("Search", func() {
		LoadingDialog.Show()
		var (
			err   error
			cpkgs []pkg.Packager
		)

		if cui.settings.String("list-mode") == "local" {
			if s.Entry.Text == "" {
				cpkgs, err = cui.manager.LocalPkgs()
			} else {
				cpkgs, err = cui.manager.SearchInLocal(s.Entry.Text)
			}
		} else {
			if s.Entry.Text == "" {
				cpkgs = []pkg.Packager{}
			} else {
				cpkgs, err = cui.manager.SearchInRepo(s.Entry.Text)
			}
		}

		if err != nil {
			popups.Error(err)
		}
		cui.packages = []packager{}
		for _, p := range cpkgs {
			cui.packages = append(cui.packages, packager{Packager: p})
		}

		cui.list.Refresh()
		LoadingDialog.Hide()
	})

	s.ModeSelect = widget.NewSelect([]string{"Local", "Repository"}, func(str string) {
		switch str {
		case "Local":
			cui.settings.SetString("list-mode", "local")
		case "Repository":
			cui.settings.SetString("list-mode", "repo")
		}
		FuncLoadingDialog(cui.InitPkgSlice)
	})
	s.Box = boxes.NewBorder(nil, nil, s.ModeSelect, s.Button, s.Entry)
}

func (s *Search) InitSelect() {
	s.ModeSelect.SetSelectedIndex(func() int {
		switch cui.settings.String("list-mode") {
		case "local":
			return 0
		case "repo":
			return 1
		default:
			return 0
		}
	}())
}
