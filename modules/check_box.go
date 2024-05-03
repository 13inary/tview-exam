package modules

import (
	"tview-exam/core"
)

func BuildCheckbox(app *core.App) *core.Checkbox {
	checkbox := core.NewCheckbox("Hit Enter to check box: ")
	return checkbox
}
