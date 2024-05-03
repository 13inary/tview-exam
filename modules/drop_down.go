package modules

import (
	"tview-exam/core"
)

func BuildDropDown(app *core.App) *core.DropDown {
	dropDown := core.NewDropDown("Select an option (hit Enter): ")
	dropDown.SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
	return dropDown
}
