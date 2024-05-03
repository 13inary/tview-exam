package modules

import "tview-exam/core"

func BuildModal(app *core.App) *core.Modal {
	modal := core.NewModal()
	modal.SetText("Do you want to quit the application?")
	modal.AddButtons([]string{"Quit", "Cancel"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Quit" {
			app.GetApp().Stop()
		}
	})
	return modal
}
