package modules

import "tview-exam/core"

func BuildButton(app *core.App) *core.Button {
	button := core.NewButton("Hit Enter to close", true, 0, 0, 22, 3)
	button.SetSelectedFunc(func() {
		app.GetApp().Stop()
	})
	return button
}
