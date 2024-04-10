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
		cui.settings.SetString(SearchEntryID, str)
	}
	s.Entry.SetText(cui.settings.String(SearchEntryID))

	s.Button = widget.NewButton(po.Get("Search"), func() {
		LoadingDialog.Show()
		var (
			err   error
			cpkgs []pkg.Packager
		)

		if cui.settings.String(ListModeID) == "local" {
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

	options := []string{
		po.Get("Local"),
		po.Get("Repository"),
	}
	s.ModeSelect = widget.NewSelect(options, func(str string) {
		switch str {
		case options[0]:
			cui.settings.SetString(ListModeID, "local")
		case options[1]:
			cui.settings.SetString(ListModeID, "repo")
		}
		FuncLoadingDialog(cui.InitPkgSlice)
	})
	s.Box = boxes.NewBorder(nil, nil, s.ModeSelect, s.Button, s.Entry)
}

func (s *Search) InitSelect() {
	s.ModeSelect.SetSelectedIndex(func() int {
		switch cui.settings.String(ListModeID) {
		case "local":
			return 0
		case "repo":
			return 1
		default:
			return 0
		}
	}())
}
