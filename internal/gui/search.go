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
	Charge struct {
		Box     *fyne.Container
		Button  *widget.Button
		LoadBar *widget.ProgressBarInfinite
	}
}

func (s *Search) toggleLoading() {
	if s.Charge.Button.Hidden {
		s.Charge.Button.Show()
		s.Charge.LoadBar.Hide()
		return
	}
	s.Charge.LoadBar.Show()
	s.Charge.Button.Hide()
}

func (s *Search) Init() {
	s.Entry = widget.NewEntry()
	s.Entry.OnChanged = func(str string) {
		cui.settings.SetString("search-entry", str)
	}
	s.Entry.SetText(cui.settings.String("search-entry"))

	s.Charge.LoadBar = widget.NewProgressBarInfinite()
	s.Charge.LoadBar.Start()
	s.Charge.LoadBar.Hide()

	s.Charge.Button = widget.NewButton("Search", func() {
		s.toggleLoading()

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

		s.toggleLoading()
	})

	s.Charge.Box = boxes.NewStack(s.Charge.Button, s.Charge.LoadBar)

	s.ModeSelect = widget.NewSelect([]string{"Local", "Repository"}, func(str string) {
		switch str {
		case "Local":
			cui.settings.SetString("list-mode", "local")
		case "Repository":
			cui.settings.SetString("list-mode", "repo")
		}
		s.toggleLoading()
		cui.InitPkgSlice()
		s.toggleLoading()
	})
	s.Box = boxes.NewBorder(nil, nil, s.ModeSelect, s.Charge.Box, s.Entry)
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
