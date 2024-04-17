package gui

import "time"

func InitSelectionChecker() {
	cui.bottomButtonsBox.Hide()
	var isShowed bool // Initialize isShowed to avoid unnecessary checks

	for {
		time.Sleep(10 * time.Millisecond)
		hasCheckedPkgs := len(checkedPkgs()) > 0 // Use a boolean variable for readability

		if hasCheckedPkgs && !isShowed {
			cui.bottomButtonsBox.Show()
			isShowed = true
		} else if !hasCheckedPkgs && isShowed {
			cui.bottomButtonsBox.Hide()
			isShowed = false
		}
	}
}
